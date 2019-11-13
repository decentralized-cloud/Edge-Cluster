// Package k3s provides functionality to provision a K3S edge cluster type and manage them
package k3s

import (
	"context"
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
	containerName           = "k3sserver"
	containerImage          = "rancher/k3s:v0.8.1"
	deploymentContainerPort = 6443
)

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

	if request.NameSpace == "" {
		request.NameSpace = apiv1.NamespaceDefault
	}

	//if not exisit create a namespace for the deployment
	if err = createProvisionNameSpace(request.NameSpace, service); err != nil {
		return
	}

	// create pod
	if err = createDeployment(service, request); err != nil {
		return
	}

	if err = createService(service, request); err != nil {
		return
	}

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

	service.logger.Info("Updating Provision With Retry")

	err = updateEdgeClient(service, request)

	if err != nil {
		service.logger.Error(
			"failed to update the pod",
			zap.Error(err))
	}

	response = &types.UpdateProvisionResponse{}

	return
}

func (service *k3sProvisioner) DeleteProvision(
	ctx context.Context,
	request *types.DeleteProvisionRequest) (response *types.DeleteProvisionResponse, err error) {
	service.logger.Info("deleting Provisio")

	err = deleteEdgeClient(service, request)

	if err != nil {
		service.logger.Error(
			"failed to delete pod",
			zap.Error(err))
	}

	response = &types.DeleteProvisionResponse{}

	return
}

func deleteEdgeClient(service *k3sProvisioner, request *types.DeleteProvisionRequest) error {
	deleteClient := service.clientset.AppsV1().Deployments(request.NameSpace)
	deletePolicy := metav1.DeletePropagationForeground

	err := deleteClient.Delete(request.Name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	return err
}

func updateEdgeClient(service *k3sProvisioner, request *types.UpdateProvisionRequest) error {
	updateClient := service.clientset.AppsV1().Deployments(request.NameSpace)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := updateClient.Get(request.Name, metav1.GetOptions{})

		if getErr != nil {
			service.logger.Error(
				"failed to update the pod",
				zap.Error(getErr))
		}

		//Do what need to be updated
		//add necessary fileds to update
		for _, container := range result.Spec.Template.Spec.Containers {
			if container.Name == request.Name {
				for _, env := range container.Env {
					if env.Name == "K3S_CLUSTER_SECRET" {
						env.Value = request.K3SClusterSecret
					}
				}
			}
		}
		//update image container
		//result.Spec.Template.Spec.Containers[0].Image = edge.ContainerImage

		_, updateErr := updateClient.Update(result)

		if updateErr != nil {
			service.logger.Error(
				"failed to update the pod",
				zap.Error(updateErr))
		}

		return updateErr
	})

	return retryErr
}

func createDeployment(service *k3sProvisioner, request *types.NewProvisionRequest) (err error) {
	deploymentClient := service.clientset.AppsV1().Deployments(request.NameSpace)
	deploymentConfig := makeDeploymentConfig(request)

	result, err := deploymentClient.Create(deploymentConfig)
	if err != nil {
		service.logger.Error(
			"failed to create pod",
			zap.Error(err),
			zap.Any("Config", deploymentConfig))

		return
	}

	service.logger.Info(
		"created a pod",
		zap.String("PodName", result.GetObjectMeta().GetName()))

	return
}

func createService(service *k3sProvisioner, request *types.NewProvisionRequest) (err error) {
	serviceDeployment := service.clientset.CoreV1().Services(request.NameSpace)
	serviceConfig := makeServiceConfig(request)

	result, err := serviceDeployment.Create(serviceConfig)
	if err != nil {
		service.logger.Error(
			"failed to create service",
			zap.Error(err),
			zap.Any("Config", serviceConfig))

		return
	}

	service.logger.Info(
		"created a service",
		zap.String("ServiceName", result.GetObjectMeta().GetName()))

	return
}

func createProvisionNameSpace(namespace string, service *k3sProvisioner) (err error) {
	service.logger.Info("checking the namespace ", zap.String("ServiceName", namespace))

	ns, err := service.clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil && strings.Contains(err.Error(), "not found") {
		service.logger.Info(err.Error())
		//create the name space
		newNameSpace := makeNameSpaceConfig(namespace)

		service.logger.Info("creating namespace", zap.String("namespace", namespace))

		if _, err = service.clientset.CoreV1().Namespaces().Create(newNameSpace); err != nil {
			service.logger.Error("failed to create namespace", zap.Error(err))

			return
		}

	} else if err != nil {
		service.logger.Error("failed to validate the requested namespace", zap.Error(err))

		return
	}

	if ns != nil && ns.GetName() == namespace {
		service.logger.Info("the namespace exists")

		return
	}

	return
}

func makeDeploymentConfig(request *types.NewProvisionRequest) (deployment *appsv1.Deployment) {
	deployment = &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name + "-deployment",
			Namespace: request.NameSpace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentReplica,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": request.Name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": request.Name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  containerName,
							Image: containerImage,
							Args: []string{
								"server",
								"--disable-agent",
								"--advertise-address=" + request.ClusterPublicIPAddress,
							},
							Env: []apiv1.EnvVar{
								{Name: "K3S_CLUSTER_SECRET", Value: "Random12345"},
								{Name: "K3S_KUBECONFIG_OUTPUT", Value: "/output/kubeconfig.yaml"},
								{Name: "K3S_KUBECONFIG_MODE", Value: "666"},
							},
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: deploymentContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	return
}

func makeServiceConfig(request *types.NewProvisionRequest) (service *apiv1.Service) {
	servicePorts := []v1.ServicePort{
		{
			Protocol:   apiv1.ProtocolTCP,
			Port:       6443,
			TargetPort: intstr.FromInt(6443),
		},
	}

	serviceSelector := map[string]string{
		"app": request.Name,
	}

	service = &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name + "-service",
			Namespace: request.NameSpace,
			Labels: map[string]string{
				"k8s-app": request.Name,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports:     servicePorts,
			Selector:  serviceSelector,
			ClusterIP: "",
		},
	}

	return
}

func makeNameSpaceConfig(namespace string) (ns *v1.Namespace) {
	ns = &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	return
}
