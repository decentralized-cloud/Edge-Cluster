// Package k3s provides functionality to provision a K3S edge cluster type and manage them
package k3s

import (
	"context"
	"os"
	"path/filepath"
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
	"k8s.io/client-go/tools/clientcmd"
)

const (
	ContainerName           = "k3sserver"
	ContainerImage          = "rancher/k3s:v0.8.1"
	DeploymentReplica       = 1
	DeploymentContainerPort = 80
)

type k3sProvisioner struct {
	logger    *zap.Logger
	clientset *kubernetes.Clientset
}

// NewK3SProvisioner creates new instance of the k3sProvisioner, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewK3SProvisioner(logger *zap.Logger) (types.EdgeClusterProvisionerContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	restConfig, err := getRestConfig(logger)

	if err != nil {
		return nil, types.NewUnknownErrorWithError("Failed to retrieve rest config", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
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
	request *types.NewProvisionRequest) (*types.NewProvisionResponse, error) {

	if request.NameSpace == "" {
		request.NameSpace = apiv1.NamespaceDefault
	}

	//if not exisit create a namespace for the deployment
	err := createProvisionNameSpace(request.NameSpace, service)

	if err != nil {
		return &types.NewProvisionResponse{}, err
	}

	// create pod
	err = createDeployment(service, request)

	if err != nil {
		return &types.NewProvisionResponse{}, err
	}

	err = createService(service, request)

	if err != nil {
		return &types.NewProvisionResponse{}, err
	}

	return &types.NewProvisionResponse{}, nil
}

func createDeployment(service *k3sProvisioner, request *types.NewProvisionRequest) error {
	deploymentClient := service.clientset.AppsV1().Deployments(request.NameSpace)

	deploymentConfig := makeDeploymentConfig(request)

	result, err := deploymentClient.Create(deploymentConfig)

	if err != nil {
		service.logger.Fatal("failed to create pod",
			zap.Any("Error", err),
			zap.Any("Config", deploymentConfig))

		return err
	}

	service.logger.Info("created a pod",
		zap.String("PodName", result.GetObjectMeta().GetName()))

	return nil
}

func createService(service *k3sProvisioner, request *types.NewProvisionRequest) error {
	serviceDeployment := service.clientset.CoreV1().Services(request.NameSpace)

	serviceConfig := makeServiceConfig(request)

	result, err := serviceDeployment.Create(serviceConfig)

	if err != nil {
		service.logger.Fatal("failed to create service",
			zap.Any("Error", err),
			zap.Any("Config", serviceConfig))

		return err
	}

	service.logger.Info("created a service",
		zap.String("ServiceName", result.GetObjectMeta().GetName()))

	return nil
}

func createProvisionNameSpace(namespace string, service *k3sProvisioner) error {
	service.logger.Info("checking the namespace ",
		zap.String("ServiceName", namespace))

	ns, err := service.clientset.CoreV1().Namespaces().Get(namespace, metav1.GetOptions{})

	if err != nil && strings.Contains(err.Error(), "not found") {
		service.logger.Info(err.Error())
		//create the name space
		newNameSpace := makeNameSpaceConfig(namespace)

		service.logger.Info("creating namespace" + namespace)

		_, err = service.clientset.CoreV1().Namespaces().Create(newNameSpace)

		if err != nil {
			service.logger.Fatal("failed to create namespace",
				zap.Any("Error", err))

			return err
		}

	} else if err != nil {
		service.logger.Fatal("failed to validate the requested namespace",
			zap.Any("Error", err))
		return err
	}

	if ns != nil && ns.GetName() == namespace {
		service.logger.Info("the namespace exists")
		return nil
	}

	return nil
}

func getRestConfig(logger *zap.Logger) (*rest.Config, error) {
	kubeconfig := os.Getenv("KUBECONFIG")

	logger.Info("path ",
		zap.String("KUBECONFIG", kubeconfig))

	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	homePath, err := os.UserHomeDir()

	logger.Info("homePath ",
		zap.String("path", homePath))
	if err != nil {
		return nil, err
	}

	kubeconfigFilePath := filepath.Join(homePath, ".kube", "config")

	logger.Info("kubePath ",
		zap.String("kube path", kubeconfigFilePath))

	_, err = os.Stat(kubeconfigFilePath)
	if !os.IsNotExist(err) {
		return clientcmd.BuildConfigFromFlags("", kubeconfigFilePath)
	}

	return rest.InClusterConfig()
}

func makeDeploymentConfig(request *types.NewProvisionRequest) *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      request.Name + "-deployment",
			Namespace: request.NameSpace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(DeploymentReplica),
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
							Name:  ContainerName,
							Image: ContainerImage,
							Args: []string{
								"server",
								"--disable-agent",
								"--advertise-address=" + request.ContainerIpAddress,
							},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: DeploymentContainerPort,
								},
							},
						},
					},
				},
			},
		},
	}

	return deployment
}

func makeServiceConfig(request *types.NewProvisionRequest) *apiv1.Service {

	servicePorts := []v1.ServicePort{
		{
			Protocol:   apiv1.ProtocolTCP,
			Port:       request.ServicePort,
			TargetPort: intstr.FromInt(request.TargetPort),
		},
	}

	serviceSelector := map[string]string{
		"app": request.Name,
	}

	service := &apiv1.Service{
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

	return service
}

func makeNameSpaceConfig(namespace string) *v1.Namespace {
	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace,
		},
	}

	return ns
}

func int32Ptr(i int32) *int32 {
	return &i
}
