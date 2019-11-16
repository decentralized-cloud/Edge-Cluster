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

	nameSpace, clusterName := getMetaData(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("nameSpace", nameSpace),
		zap.String("name", clusterName))

	if nameSpace == "" {
		nameSpace = apiv1.NamespaceDefault
	}

	//if not exisit create a namespace for the deployment
	if err = createProvisionNameSpace(service, nameSpace); err != nil {
		return
	}

	// create edge cluster
	if err = createDeployment(service, nameSpace, clusterName, request.ClusterPublicIPAddress); err != nil {
		return
	}

	if err = createService(service, nameSpace, clusterName); err != nil {
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

	service.logger.Info("metadata", zap.Any("request", request))

	service.logger.Info("Updating Provision With Retry")

	nameSpace, clusterName := getMetaData(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("nameSpace", nameSpace),
		zap.String("name", clusterName))

	err = updateEdgeClient(service, nameSpace, clusterName, request.K3SClusterSecret)

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
	service.logger.Info("deleting Provisio")

	nameSpace, clusterName := getMetaData(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("nameSpace", nameSpace),
		zap.String("name", clusterName))

	err = deleteEdgeClient(service, nameSpace, clusterName)

	if err != nil {
		service.logger.Error(
			"failed to delete edge cluster",
			zap.Error(err))
	}

	err = deleteProvisionNameSpace(service, nameSpace)

	if err != nil {
		service.logger.Error(
			"failed to delete namespace",
			zap.Error(err))
	}

	response = &types.DeleteProvisionResponse{}

	return
}

func deleteEdgeClient(service *k3sProvisioner,
	namespace string,
	clusterName string) error {

	deleteClient := service.clientset.AppsV1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground

	err := deleteClient.Delete(clusterName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	return err
}

func updateEdgeClient(service *k3sProvisioner,
	namespace string,
	clusterName string,
	secretKey string) error {

	updateClient := service.clientset.AppsV1().Deployments(namespace)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := updateClient.Get(clusterName, metav1.GetOptions{})

		if getErr != nil {
			service.logger.Error(
				"failed to update the edge cluster",
				zap.Error(getErr))
		}

		//Do what need to be updated
		//add necessary fileds to update
		for _, container := range result.Spec.Template.Spec.Containers {

			service.logger.Info("update value", zap.String("name:", container.Name))

			if container.Name == clusterName {
				for _, env := range container.Env {
					if env.Name == "K3S_CLUSTER_SECRET" {
						env.Value = secretKey
					}
				}
			}
		}
		//update image container
		//result.Spec.Template.Spec.Containers[0].Image = edge.ContainerImage

		_, updateErr := updateClient.Update(result)

		if updateErr != nil {
			service.logger.Error(
				"failed to update the edge custer",
				zap.Error(updateErr))
		}

		return updateErr
	})

	return retryErr
}

func createDeployment(service *k3sProvisioner,
	namespace string,
	clusterName string,
	clusterIP string) (err error) {

	deploymentClient := service.clientset.AppsV1().Deployments(namespace)

	deploymentConfig := makeDeploymentConfig(namespace, clusterName, clusterIP)

	result, err := deploymentClient.Create(deploymentConfig)
	if err != nil {
		service.logger.Error(
			"failed to create edge cluster",
			zap.Error(err),
			zap.Any("Config", deploymentConfig))

		return
	}

	service.logger.Info(
		"created a edge cluster",
		zap.String("Edge cluster Name", result.GetObjectMeta().GetName()))

	return
}

func createService(service *k3sProvisioner,
	namespace string,
	clusterName string) (err error) {

	serviceDeployment := service.clientset.CoreV1().Services(namespace)

	serviceConfig := makeServiceConfig(namespace, clusterName)

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

func createProvisionNameSpace(service *k3sProvisioner, namespace string) (err error) {
	service.logger.Info("checking the namespace ", zap.String("ServiceName", namespace))

	ns, err := service.clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil && strings.Contains(err.Error(), "not found") {
		service.logger.Info(err.Error())
		//create the name space

		service.logger.Info("creating namespace", zap.String("namespace", namespace))

		newNameSpace := makeNameSpaceConfig(namespace)

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

func deleteProvisionNameSpace(service *k3sProvisioner, namespace string) (err error) {

	deletePolicy := metav1.DeletePropagationForeground

	if err = service.clientset.CoreV1().Namespaces().Delete(namespace, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		service.logger.Error("failed to delete namespace", zap.Error(err))

		return err
	}

	return nil
}

func makeDeploymentConfig(namespace string,
	clusterName string,
	clusterIP string) (deployment *appsv1.Deployment) {

	deployment = &appsv1.Deployment{
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
								"--disable-agent",
								"--advertise-address=" + clusterIP,
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

func makeServiceConfig(namespace string,
	clusterName string) (service *apiv1.Service) {

	servicePorts := []v1.ServicePort{
		{
			Protocol:   apiv1.ProtocolTCP,
			Port:       6443,
			TargetPort: intstr.FromInt(6443),
		},
	}

	serviceSelector := map[string]string{
		"app": clusterName,
	}

	annotation := map[string]string{
		"metallb.universe.tf/address-pool": "default",
	}

	//todo add anotation to connect metallb
	service = &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      clusterName,
			Namespace: namespace,
			Labels: map[string]string{
				"k8s-app": clusterName,
			},
			Annotations: annotation,
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

func getMetaData(edgeClusterID string) (string, string) {

	var namespace, clustername string = "ns", "edge"

	if edgeClusterID != "" {
		hashCode := fmt.Sprintf("%x", sha256.Sum224([]byte(edgeClusterID)))

		sequence := []rune(hashCode)

		for i, s := range sequence {
			if i < 32 {
				namespace += strings.TrimSpace(string(s))
			} else {
				clustername += strings.TrimSpace(string(s))
			}
		}
	}

	return namespace, clustername
}
