// Package k3s provides functionality to provision a K3S edge cluster type and manage them
package k3s

import (
	"context"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/retry"
)

const (
	containerName  = "k3sserver"
	containerImage = "rancher/k3s:v1.20.0-k3s2"
	k3sPort        = 6443
	internalName   = "k3s"
)

var waitForFunctionToBeReadyTimeout int64 = 60
var deploymentReplica int32 = 1

type k3sProvisioner struct {
	logger    *zap.Logger
	clientset *kubernetes.Clientset
}

// NewK3SProvisioner creates new instance of the k3sProvisioner, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// k8sRestConfig: Mandatory. Reference to the Rest config points to the running K8S cluster
// Returns the new service or error if something goes wrong
func NewK3SProvisioner(
	logger *zap.Logger,
	k8sRestConfig *rest.Config) (types.EdgeClusterProvisionerContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if k8sRestConfig == nil {
		return nil, commonErrors.NewArgumentNilError("k8sRestConfig", "k8sRestConfig is required")
	}

	clientset, err := kubernetes.NewForConfig(k8sRestConfig)
	if err != nil {
		return nil, types.NewUnknownErrorWithError("Failed to create client set", err)
	}

	return &k3sProvisioner{
		logger:    logger,
		clientset: clientset,
	}, nil
}

// NewProvision provisions a new K3S edge cluster.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to provision a new edge cluster
// Returns either the result of provisioning new K3S edge cluster or error if something goes wrong.
func (service *k3sProvisioner) NewProvision(
	ctx context.Context,
	request *types.NewProvisionRequest) (response *types.NewProvisionResponse, err error) {
	response = nil
	namespace := getNamespace(request.EdgeClusterID)

	service.logger.Info("metadata", zap.String("namespace", namespace))

	//if not exisit create a namespace for the deployment
	if err = service.createProvisionNameSpace(ctx, namespace); err != nil {
		return
	}

	//create a loadbalancer
	externalIP, err := service.createService(ctx, namespace)

	if err != nil {
		return
	}

	service.logger.Info("load balancer:", zap.String("externalIP", externalIP))

	// create edge cluster
	if err = service.createDeployment(
		ctx,
		namespace,
		internalName,
		request.ClusterSecret,
		externalIP); err != nil {
		return
	}

	//TODO: we can watch and wait for status and return it here or through a separate API call

	response = &types.NewProvisionResponse{}

	return
}

// UpdateProvisionWithRetry update and existing K3S edge cluster.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to update an edge cluster
// Returns either the result of updating the K3S edge cluster or error if something goes wrong.
func (service *k3sProvisioner) UpdateProvisionWithRetry(
	ctx context.Context,
	request *types.UpdateProvisionRequest) (response *types.UpdateProvisionResponse, err error) {
	response = nil

	service.logger.Info("metadata", zap.Any("request", request))

	service.logger.Info("Updating Provision With Retry")

	namespace := getNamespace(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("namespace", namespace))

	err = service.updateEdgeClient(ctx, namespace, request.ClusterSecret)

	if err != nil {
		service.logger.Error(
			"failed to update the edge cluster",
			zap.Error(err))
	}

	response = &types.UpdateProvisionResponse{}

	return
}

func (service *k3sProvisioner) DeleteProvision(
	ctx context.Context,
	request *types.DeleteProvisionRequest) (response *types.DeleteProvisionResponse, err error) {
	service.logger.Info("deleting Provision")

	namespace := getNamespace(request.EdgeClusterID)

	service.logger.Info("metadata", zap.String("namespace", namespace))

	err = service.deleteEdgeClient(ctx, namespace)

	if err != nil {
		service.logger.Error(
			"failed to delete edge cluster",
			zap.Error(err))
	}

	err = service.deleteProvisionNameSpace(ctx, namespace)

	if err != nil {
		service.logger.Error(
			"failed to delete namespace",
			zap.Error(err))
	}

	response = &types.DeleteProvisionResponse{}

	return
}

func (service *k3sProvisioner) deleteEdgeClient(ctx context.Context, namespace string) error {
	client := service.clientset.AppsV1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground

	return client.Delete(ctx, internalName, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
}

func (service *k3sProvisioner) updateEdgeClient(
	ctx context.Context,
	namespace string,
	secretKey string) error {
	return retry.RetryOnConflict(
		retry.DefaultRetry,
		func() (err error) {
			client := service.clientset.AppsV1().Deployments(namespace)

			deployment, err := client.Get(ctx, internalName, metav1.GetOptions{})
			if err != nil {
				service.logger.Error(
					"failed to update the edge cluster",
					zap.Error(err))

				return
			}

			//Do what need to be updated
			//add necessary fileds to update
			containerIndex := getContainerIndex(deployment.Spec.Template.Spec.Containers)
			if containerIndex > -1 {
				envIndex := getEnvIndex(deployment.Spec.Template.Spec.Containers[containerIndex].Env)

				deployment.Spec.Template.Spec.Containers[containerIndex].Env[envIndex].Value = secretKey
			}

			//update image container
			//result.Spec.Template.Spec.Containers[0].Image = edge.ContainerImage
			service.logger.Info("config", zap.Any("spec", deployment))

			_, err = client.Update(ctx, deployment, metav1.UpdateOptions{})
			if err != nil {
				service.logger.Error(
					"failed to update the edge custer",
					zap.Error(err))

				return
			}

			return
		})
}

func getContainerIndex(containers []v1.Container) int {
	for index, container := range containers {
		if container.Name == containerName {
			return index
		}
	}

	return -1
}

func getEnvIndex(envList []v1.EnvVar) int {
	for index, env := range envList {
		if env.Name == "K3S_CLUSTER_SECRET" {
			return index
		}
	}

	return -1
}

func (service *k3sProvisioner) createDeployment(
	ctx context.Context,
	namespace string,
	clusterName string,
	k3SClusterSecret string,
	publicIP string) (err error) {
	client := service.clientset.AppsV1().Deployments(namespace)
	deploymentConfig := service.makeDeploymentConfig(namespace, clusterName, k3SClusterSecret, publicIP)

	if result, err := client.Create(ctx, deploymentConfig, metav1.CreateOptions{}); err != nil {
		service.logger.Error(
			"failed to create edge cluster",
			zap.Error(err),
			zap.Any("Config", deploymentConfig))
	} else {
		service.logger.Info(
			"created a edge cluster",
			zap.String("Edge cluster Name", result.GetObjectMeta().GetName()))
	}

	return
}

func (service *k3sProvisioner) createService(ctx context.Context, namespace string) (externalIP string, err error) {
	serviceDeployment := service.clientset.CoreV1().Services(namespace)
	serviceConfig := service.makeServiceConfig(namespace)

	result, err := serviceDeployment.Create(ctx, serviceConfig, metav1.CreateOptions{})

	if err != nil {
		service.logger.Error(
			"failed to create service",
			zap.Error(err),
			zap.Any("Config", serviceConfig))
	} else {
		service.logger.Info(
			"created the service",
			zap.String("ServiceName", result.GetObjectMeta().GetName()))
	}

	//need to wait for a couple of seconds to get the external IPs
	watch, err := service.clientset.CoreV1().Services(namespace).Watch(ctx, metav1.ListOptions{
		Watch:          true,
		TimeoutSeconds: &waitForFunctionToBeReadyTimeout,
	})

	if err != nil {
		service.logger.Error(
			"failed to fetch the service",
			zap.Error(err))
		return "", err
	}

	for event := range watch.ResultChan() {
		if service, ok := event.Object.(*v1.Service); ok {
			for _, item := range service.Status.LoadBalancer.Ingress {
				if item.IP != "" {
					watch.Stop()

					return item.IP, nil
				}
			}
		}
	}

	return "", nil
}

func (service *k3sProvisioner) createProvisionNameSpace(ctx context.Context, namespace string) (err error) {
	service.logger.Info("checking the namespace ", zap.String("ServiceName", namespace))

	ns, err := service.clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if err != nil && strings.Contains(err.Error(), "not found") {
		service.logger.Info(
			"creating namespace",
			zap.String("namespace", namespace))

		namespaceConfig := service.makeNameSpaceConfig(namespace)
		if _, err = service.clientset.CoreV1().Namespaces().Create(ctx, namespaceConfig, metav1.CreateOptions{}); err != nil {
			service.logger.Error(
				"failed to create namespace",
				zap.Error(err),
				zap.String("namespace", namespace))

			return
		}

	} else if err != nil {
		service.logger.Error(
			"failed to validate the requested namespace",
			zap.Error(err),
			zap.String("namespace", namespace))

		return
	}

	if ns != nil && ns.GetName() == namespace {
		service.logger.Info(
			"namespace already exists",
			zap.String("namespace", namespace))

		return
	}

	return
}

func (service *k3sProvisioner) deleteProvisionNameSpace(ctx context.Context, namespace string) (err error) {
	deletePolicy := metav1.DeletePropagationForeground
	if err = service.clientset.CoreV1().Namespaces().Delete(
		ctx,
		namespace,
		metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
		service.logger.Error("failed to delete namespace", zap.Error(err))

		return
	}

	return
}

func (service *k3sProvisioner) makeDeploymentConfig(namespace string,
	clusterName string,
	k3SClusterSecret string,
	advertiseIPAddress string) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentReplica,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": clusterName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": clusterName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  containerName,
							Image: containerImage,
							Args: []string{
								"server",
							},
							Env: []apiv1.EnvVar{
								{Name: "K3S_CLUSTER_SECRET", Value: k3SClusterSecret},
								{Name: "K3S_KUBECONFIG_OUTPUT", Value: "/output/kubeconfig.yaml"},
								{Name: "K3S_KUBECONFIG_MODE", Value: "666"},
							},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "k3s",
									ContainerPort: k3sPort,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (service *k3sProvisioner) makeServiceConfig(namespace string) (apiv1Service *apiv1.Service) {
	servicePorts := []v1.ServicePort{
		{
			Name:       "k3s",
			Protocol:   apiv1.ProtocolTCP,
			Port:       k3sPort,
			TargetPort: intstr.FromInt(k3sPort),
		},
	}

	serviceSelector := map[string]string{
		"app": internalName,
	}

	apiv1Service = &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      internalName,
			Namespace: namespace,
			Labels: map[string]string{
				"k8s-app": internalName,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports:    servicePorts,
			Selector: serviceSelector,
			Type:     apiv1.ServiceTypeLoadBalancer,
		},
	}

	return
}

func (service *k3sProvisioner) makeNameSpaceConfig(namespace string) *v1.Namespace {
	return &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}
}

func getNamespace(edgeClusterID string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(edgeClusterID)))
}
