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
	nameSpace, clusterName := service.getMetaData(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("nameSpace", nameSpace),
		zap.String("name", clusterName))

	if nameSpace == "" {
		nameSpace = apiv1.NamespaceDefault
	}

	//if not exisit create a namespace for the deployment
	if err = service.createProvisionNameSpace(nameSpace); err != nil {
		return
	}

	// create edge cluster
	if err = service.createDeployment(nameSpace, clusterName, request.K3SClusterSecret); err != nil {
		return
	}

	if err = service.createService(nameSpace, clusterName); err != nil {
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

	nameSpace, clusterName := service.getMetaData(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("nameSpace", nameSpace),
		zap.String("name", clusterName))

	err = service.updateEdgeClient(nameSpace, clusterName, request.K3SClusterSecret)

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

	nameSpace, clusterName := service.getMetaData(request.EdgeClusterID)

	service.logger.Info("metadata",
		zap.String("nameSpace", nameSpace),
		zap.String("name", clusterName))

	err = service.deleteEdgeClient(nameSpace, clusterName)

	if err != nil {
		service.logger.Error(
			"failed to delete edge cluster",
			zap.Error(err))
	}

	err = service.deleteProvisionNameSpace(nameSpace)

	if err != nil {
		service.logger.Error(
			"failed to delete namespace",
			zap.Error(err))
	}

	response = &types.DeleteProvisionResponse{}

	return
}

func (service *k3sProvisioner) deleteEdgeClient(
	namespace string,
	clusterName string) error {
	deleteClient := service.clientset.AppsV1().Deployments(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	err := deleteClient.Delete(clusterName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	return err
}

func (service *k3sProvisioner) updateEdgeClient(
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
		containerIndex := getContainerIndex(result.Spec.Template.Spec.Containers)
		if containerIndex > -1 {

			envIndex := getEnvIndex(result.Spec.Template.Spec.Containers[containerIndex].Env)

			result.Spec.Template.Spec.Containers[containerIndex].Env[envIndex].Value = secretKey
		}

		//update image container
		//result.Spec.Template.Spec.Containers[0].Image = edge.ContainerImage
		service.logger.Info("config", zap.Any("spec", result))
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
	namespace string,
	clusterName string,
	k3SClusterSecret string) (err error) {
	deploymentClient := service.clientset.AppsV1().Deployments(namespace)
	deploymentConfig := service.makeDeploymentConfig(namespace, clusterName, k3SClusterSecret)

	if result, err := deploymentClient.Create(deploymentConfig); err != nil {
		service.logger.Error(
			"failed to create edge cluster",
			zap.Error(err),
			zap.Any("Config", deploymentConfig))
	} else {
		service.logger.Info(
			"created a edge cluster",
			zap.String("Edge cluster Name", result.GetObjectMeta().GetName()))
	}

	service.logger.Info(
		"created a edge cluster",
		zap.String("Edge cluster Name", result.GetObjectMeta().GetName()))

	return
}

func (service *k3sProvisioner) createService(
	namespace string,
	clusterName string) (err error) {
	serviceDeployment := service.clientset.CoreV1().Services(namespace)
	serviceConfig := service.makeServiceConfig(namespace, clusterName)
	if result, err := serviceDeployment.Create(serviceConfig); err != nil {
		service.logger.Error(
			"failed to create service",
			zap.Error(err),
			zap.Any("Config", serviceConfig))
	} else {
		service.logger.Info(
			"created a service",
			zap.String("ServiceName", result.GetObjectMeta().GetName()))
	}

	return
}

func (service *k3sProvisioner) createProvisionNameSpace(namespace string) (err error) {
	service.logger.Info("checking the namespace ", zap.String("ServiceName", namespace))

	ns, err := service.clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil && strings.Contains(err.Error(), "not found") {
		service.logger.Info(err.Error())
		//create the name space

		service.logger.Info("creating namespace", zap.String("namespace", namespace))

		newNameSpace := service.makeNameSpaceConfig(namespace)

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

func (service *k3sProvisioner) deleteProvisionNameSpace(namespace string) (err error) {
	deletePolicy := metav1.DeletePropagationForeground
	if err = service.clientset.CoreV1().Namespaces().Delete(
		namespace,
		&metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
		service.logger.Error("failed to delete namespace", zap.Error(err))

		return
	}

	return
}

func (service *k3sProvisioner) makeDeploymentConfig(namespace string,
	clusterName string,
	k3SClusterSecret string) (deployment *appsv1.Deployment) {
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
								"--advertise-address=10.0.0.1",
							},
							Env: []apiv1.EnvVar{
								{Name: "K3S_CLUSTER_SECRET", Value: k3SClusterSecret},
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

func (service *k3sProvisioner) makeServiceConfig(
	namespace string,
	clusterName string) (apiv1Service *apiv1.Service) {
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
	apiv1Service = &apiv1.Service{
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

func (service *k3sProvisioner) makeNameSpaceConfig(namespace string) (ns *v1.Namespace) {
	ns = &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	return
}

func (service *k3sProvisioner) getMetaData(edgeClusterID string) (namespace string, clusterName string) {
	namespace, clusterName = "ns", "edge"
	hashCode := fmt.Sprintf("%x", sha256.Sum224([]byte(edgeClusterID)))
	sequence := []rune(hashCode)

	for i, s := range sequence {
		if i < 32 {
			namespace += strings.TrimSpace(string(s))
		} else {
			clusterName += strings.TrimSpace(string(s))
		}
	}

	return
}
