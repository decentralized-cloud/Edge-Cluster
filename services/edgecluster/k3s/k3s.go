// Package k3s provides functionality to provision a K3S edge cluster type and manage them
package k3s

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/configuration"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/helm"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // register GCP auth provider
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/retry"
)

const (
	containerName      = "k3sserver"
	k3sPort            = 6443
	internalName       = "k3s"
	kubeconfigFilePath = "/etc/rancher/k3s/k3s.yaml"
)

var deploymentReplica int32 = 1
var waitForDeploymentToBeReadyTimeout int64 = 120

type k3sProvisioner struct {
	logger         *zap.Logger
	clientset      *kubernetes.Clientset
	k8sRestConfig  *rest.Config
	k3sDockerImage string
	helmService    helm.HelmHelperContract
}

// NewK3SProvisioner creates new instance of the k3sProvisioner, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// k8sRestConfig: Mandatory. Reference to the Rest config points to the running K8S cluster
// Returns the new service or error if something goes wrong
func NewK3SProvisioner(
	logger *zap.Logger,
	k8sRestConfig *rest.Config,
	configurationService configuration.ConfigurationContract,
	helmService helm.HelmHelperContract) (types.EdgeClusterProvisionerContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if k8sRestConfig == nil {
		return nil, commonErrors.NewArgumentNilError("k8sRestConfig", "k8sRestConfig is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	if helmService == nil {
		return nil, commonErrors.NewArgumentNilError("helmService", "helmService is required")
	}

	k3sDockerImage, err := configurationService.GetK3SDockerImage()
	if err != nil {
		return nil, types.NewUnknownErrorWithError("failed to get the database name", err)
	}

	var clientset *kubernetes.Clientset
	if clientset, err = kubernetes.NewForConfig(k8sRestConfig); err != nil {
		return nil, types.NewUnknownErrorWithError("failed to create client set", err)
	}

	return &k3sProvisioner{
		logger:         logger,
		clientset:      clientset,
		k8sRestConfig:  k8sRestConfig,
		k3sDockerImage: k3sDockerImage,
		helmService:    helmService,
	}, nil
}

// CreateProvision provisions a new edge cluster.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to provision a new edge cluster
// Returns either the result of provisioning new edge cluster or error if something goes wrong.
func (service *k3sProvisioner) CreateProvision(
	ctx context.Context,
	request *types.CreateProvisionRequest) (response *types.CreateProvisionResponse, err error) {
	namespace := getNamespace(request.EdgeClusterID)

	if err = service.createProvisionNameSpace(ctx, namespace); err != nil {
		return
	}

	if err = service.createService(ctx, namespace); err != nil {
		_, _ = service.DeleteProvision(ctx, &types.DeleteProvisionRequest{EdgeClusterID: request.EdgeClusterID})

		return
	}

	if err = service.createDeployment(
		ctx,
		namespace,
		request.ClusterSecret); err != nil {
		_, _ = service.DeleteProvision(ctx, &types.DeleteProvisionRequest{EdgeClusterID: request.EdgeClusterID})

		return
	}

	response = &types.CreateProvisionResponse{}

	isReady, err := service.isK3SPodReady(ctx, namespace)
	if err != nil {
		service.logger.Error("K3S pod status is not ready", zap.Error(err))

		return
	}

	if !isReady {
		service.logger.Error("K3S pod status is not ready")

		return
	}

	if err = service.deployHelmChart(ctx, request.EdgeClusterID); err != nil {
		service.logger.Error("failed to install the helm charts", zap.Error(err))

		return
	}

	return
}

// UpdateProvisionWithRetry updates an existing provision.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to update an existing provision
// Returns either the result of updating an existing provision or error if something goes wrong.
func (service *k3sProvisioner) UpdateProvisionWithRetry(
	ctx context.Context,
	request *types.UpdateProvisionRequest) (response *types.UpdateProvisionResponse, err error) {

	namespace := getNamespace(request.EdgeClusterID)

	err = retry.RetryOnConflict(
		retry.DefaultRetry,
		func() (err error) {
			client := service.clientset.AppsV1().Deployments(namespace)

			deployment, err := client.Get(ctx, internalName, metav1.GetOptions{})
			if err != nil {
				service.logger.Error("failed to update the edge cluster", zap.Error(err))

				return
			}

			deployment.Spec.Template.Spec, err = service.getDeploymentSpec(ctx, namespace, request.ClusterSecret)
			if err != nil {
				return err
			}

			if _, err = client.Update(ctx, deployment, metav1.UpdateOptions{}); err != nil {
				service.logger.Error("failed to update the edge custer", zap.Error(err))

				return
			}

			return
		})

	response = &types.UpdateProvisionResponse{}

	isReady, internalErr := service.isK3SPodReady(ctx, namespace)
	if internalErr != nil {
		service.logger.Error("K3S pod status is not ready", zap.Error(internalErr))

		return
	}

	if !isReady {
		service.logger.Error("K3S pod status is not ready")

		return
	}

	if internalErr = service.deployHelmChart(ctx, request.EdgeClusterID); internalErr != nil {
		service.logger.Error("failed to install the helm charts", zap.Error(internalErr))

		return
	}

	return
}

// DeleteProvision deletes an existing provision.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to delete an existing provision
// Returns either the result of deleting an existing provision or error if something goes wrong.
func (service *k3sProvisioner) DeleteProvision(
	ctx context.Context,
	request *types.DeleteProvisionRequest) (response *types.DeleteProvisionResponse, err error) {
	namespace := getNamespace(request.EdgeClusterID)

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

	response = &types.DeleteProvisionResponse{}

	return
}

// GetProvisionDetails retrieves information on an existing provision.
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to retrieve information on an existing provision
// Returns either the result of retrieving information on an provision or error if something goes wrong.
func (service *k3sProvisioner) GetProvisionDetails(
	ctx context.Context,
	request *types.GetProvisionDetailsRequest) (response *types.GetProvisionDetailsResponse, err error) {
	namespace := getNamespace(request.EdgeClusterID)

	serviceDetails, err := service.getProvvisionedServiceDetails(ctx, namespace)
	if err != nil {
		service.logger.Error("failed to get service details", zap.Error(err))

		return
	}

	kubeconfigContent, err := service.getProvisionDetailsKubeConfigContent(ctx, namespace)
	if err != nil {
		service.logger.Error("failed to get kubeconfig content", zap.Error(err))

		return
	}

	for _, item := range serviceDetails.Status.LoadBalancer.Ingress {
		if item.IP != "" {
			kubeconfigContent = strings.Replace(kubeconfigContent, "127.0.0.1", item.IP, -1)
		} else if item.Hostname != "" {
			kubeconfigContent = strings.Replace(kubeconfigContent, "127.0.0.1", item.Hostname, -1)
		} else {
			kubeconfigContent = strings.Replace(kubeconfigContent, "127.0.0.1", "BLANK", -1)
		}
	}

	for _, item := range serviceDetails.Spec.Ports {
		kubeconfigContent = strings.Replace(kubeconfigContent, fmt.Sprintf("%d", k3sPort), fmt.Sprintf("%d", item.Port), -1)
	}

	response = &types.GetProvisionDetailsResponse{
		ProvisionDetails: models.ProvisionDetails{
			Service:           serviceDetails,
			KubeconfigContent: kubeconfigContent,
		}}

	return
}

// ListNodes lists an existing edge cluster nodes details
// ctx: Mandatory The reference to the context
// request: Mandatory. The request to list an existing edge cluster nodes details
// Returns an existing edge cluster nodes details or error if something goes wrong.
func (service *k3sProvisioner) ListNodes(
	ctx context.Context,
	request *types.ListNodesRequest) (response *types.ListNodesResponse, err error) {
	clientset, err := service.createClientsetForEdgeCluster(ctx, request.EdgeClusterID)
	if err != nil {
		return nil, err
	}

	var nodeList *v1.NodeList
	if nodeList, err = clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{}); err != nil {
		return nil, types.NewUnknownErrorWithError("failed to retreive node list", err)
	}

	response = &types.ListNodesResponse{Nodes: []models.EdgeClusterNode{}}
	for _, node := range nodeList.Items {
		response.Nodes = append(response.Nodes, models.EdgeClusterNode{
			Node: node,
		})
	}

	return
}

// ListPods lists an existing edge cluster pods that matchs the given search criteria
// ctx: Mandatory The reference to the context
// request: Mandatory. The request that contains the search criteria to filter the pods
// Returns the list of running pods that matchs the given search criteria or error if something goes wrong.
func (service *k3sProvisioner) ListPods(
	ctx context.Context,
	request *types.ListPodsRequest) (response *types.ListPodsResponse, err error) {
	clientset, err := service.createClientsetForEdgeCluster(ctx, request.EdgeClusterID)
	if err != nil {
		return nil, err
	}

	var pods *v1.PodList
	if pods, err = clientset.CoreV1().Pods(request.Namespace).List(ctx, metav1.ListOptions{}); err != nil {
		return nil, types.NewUnknownErrorWithError("failed to retreive pod list", err)
	}

	response = &types.ListPodsResponse{Pods: []models.EdgeClusterPod{}}
	for _, pod := range pods.Items {
		response.Pods = append(response.Pods, models.EdgeClusterPod{
			Pod: pod,
		})
	}

	return
}

// ListServices lists an existing edge cluster services that matchs the given search criteria
// ctx: Mandatory The reference to the context
// request: Mandatory. The request that contains the search criteria to filter the services
// Returns the list of services that matches the given search criteria or error if something goes wrong.
func (service *k3sProvisioner) ListServices(
	ctx context.Context,
	request *types.ListServicesRequest) (response *types.ListServicesResponse, err error) {
	clientset, err := service.createClientsetForEdgeCluster(ctx, request.EdgeClusterID)
	if err != nil {
		return nil, err
	}

	var services *v1.ServiceList
	if services, err = clientset.CoreV1().Services(request.Namespace).List(ctx, metav1.ListOptions{}); err != nil {
		return nil, types.NewUnknownErrorWithError("failed to retreive service list", err)
	}

	response = &types.ListServicesResponse{Services: []models.EdgeClusterService{}}
	for _, service := range services.Items {
		response.Services = append(response.Services, models.EdgeClusterService{
			Service: service,
		})
	}

	return
}

func (service *k3sProvisioner) createProvisionNameSpace(ctx context.Context, namespace string) (err error) {
	ns, err := service.clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if err != nil && strings.Contains(err.Error(), "not found") {
		namespaceConfig := &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: namespace,
			},
		}

		if _, err = service.clientset.CoreV1().Namespaces().Create(ctx, namespaceConfig, metav1.CreateOptions{}); err != nil {
			service.logger.Error(
				"failed to create namespace",
				zap.Error(err),
				zap.String("namespace", namespace))

			return
		}

	} else if err != nil {
		service.logger.Error("failed to validate the requested namespace", zap.Error(err), zap.String("namespace", namespace))

		return
	}

	if ns != nil && ns.GetName() == namespace {
		service.logger.Info("namespace already exists", zap.String("namespace", namespace))

		return
	}

	return
}

func (service *k3sProvisioner) createDeployment(
	ctx context.Context,
	namespace string,
	k3SClusterSecret string) (err error) {
	spec, err := service.getDeploymentSpec(ctx, namespace, k3SClusterSecret)
	if err != nil {
		return err
	}

	client := service.clientset.AppsV1().Deployments(namespace)
	deploymentConfig := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      internalName,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentReplica,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					internalName: internalName,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						internalName: internalName,
					},
				},
				Spec: spec,
			},
		},
	}

	if _, err := client.Create(ctx, deploymentConfig, metav1.CreateOptions{}); err != nil {
		service.logger.Error("failed to create edge cluster", zap.Error(err), zap.Any("Config", deploymentConfig))
	}

	return
}

func (service *k3sProvisioner) createService(ctx context.Context, namespace string) (err error) {
	serviceDeployment := service.clientset.CoreV1().Services(namespace)

	servicePorts := []v1.ServicePort{
		{
			Name:       internalName,
			Protocol:   v1.ProtocolTCP,
			Port:       k3sPort,
			TargetPort: intstr.FromInt(k3sPort),
		},
	}

	serviceSelector := map[string]string{
		internalName: internalName,
	}

	serviceConfig := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      internalName,
			Namespace: namespace,
			Labels: map[string]string{
				"k8s-app": internalName,
			},
		},
		Spec: v1.ServiceSpec{
			Ports:    servicePorts,
			Selector: serviceSelector,
			Type:     v1.ServiceTypeLoadBalancer,
		},
	}

	if _, err = serviceDeployment.Create(ctx, serviceConfig, metav1.CreateOptions{}); err != nil {
		service.logger.Error("failed to create service", zap.Error(err), zap.Any("Config", serviceConfig))

		return
	}

	return
}

func getNamespace(edgeClusterID string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(edgeClusterID)))
}

func (service *k3sProvisioner) getDeploymentSpec(ctx context.Context, namespace string, k3SClusterSecret string) (v1.PodSpec, error) {
	advertiseAddress, err := service.getAdvertiseAddress(ctx, namespace)
	if err != nil {
		return v1.PodSpec{}, err
	}

	return v1.PodSpec{
		Containers: []v1.Container{
			{
				Name:  containerName,
				Image: service.k3sDockerImage,
				Args: []string{
					"server",
					fmt.Sprintf("--advertise-address=%s", advertiseAddress),
				},
				Env: []v1.EnvVar{
					{Name: "K3S_CLUSTER_SECRET", Value: k3SClusterSecret},
				},
				Ports: []v1.ContainerPort{
					{
						Name:          internalName,
						ContainerPort: k3sPort,
					},
				},
			},
		},
	}, nil
}

func (service *k3sProvisioner) getProvvisionedServiceDetails(
	ctx context.Context,
	namespace string) (*v1.Service, error) {

	var serviceInfo *v1.Service
	var err error
	if serviceInfo, err = service.clientset.CoreV1().Services(namespace).Get(ctx, internalName, metav1.GetOptions{}); err != nil {
		service.logger.Error("failed to fetch service info", zap.Error(err))

		return nil, err
	}

	return serviceInfo, nil
}

func (service *k3sProvisioner) getProvisionDetailsKubeConfigContent(
	ctx context.Context,
	namespace string) (kubeconfigContent string, err error) {
	pods, err := service.clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return
	}

	if len(pods.Items) <= 0 {
		return "", types.NewUnknownError("Pod is not ready yet")
	}

	execRequest := service.clientset.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Name(pods.Items[0].ObjectMeta.Name).
		Namespace(pods.Items[0].ObjectMeta.Namespace).
		SubResource("exec").
		Param("stdout", "true").
		Param("command", "cat").
		Param("command", kubeconfigFilePath)

	executor, err := remotecommand.NewSPDYExecutor(service.k8sRestConfig, http.MethodPost, execRequest.URL())
	if err != nil {
		err = types.NewUnknownErrorWithError("failed to retrieve KubeConfig content.", err)

		return
	}

	output := &bytes.Buffer{}

	if err = executor.Stream(remotecommand.StreamOptions{Stdout: output}); err != nil {
		err = types.NewUnknownErrorWithError("failed to retrieve KubeConfig content", err)

		return
	}

	kubeconfigContent = output.String()

	return
}

func (service *k3sProvisioner) getAdvertiseAddress(
	ctx context.Context,
	namespace string) (string, error) {
	watch, err := service.clientset.CoreV1().Services(namespace).Watch(ctx, metav1.ListOptions{
		Watch:          true,
		TimeoutSeconds: &waitForDeploymentToBeReadyTimeout,
	})
	if err != nil {
		service.logger.Error("failed to retrieve advertise address", zap.Error(err))

		return "", err
	}

	for event := range watch.ResultChan() {
		if service, ok := event.Object.(*v1.Service); ok {
			for _, item := range service.Status.LoadBalancer.Ingress {
				if item.IP != "" {
					watch.Stop()

					return item.IP, nil
				} else if item.Hostname != "" {
					watch.Stop()

					return item.Hostname, nil
				}
			}
		}
	}

	return "", types.NewUnknownError("failed to retrieve advertise address")
}

func (service *k3sProvisioner) createClientsetForEdgeCluster(
	ctx context.Context,
	edgeClusterID string) (clientset *kubernetes.Clientset, err error) {
	getProvisionDetailsResponse, err := service.GetProvisionDetails(
		ctx,
		&types.GetProvisionDetailsRequest{
			EdgeClusterID: edgeClusterID,
		})
	if err != nil {
		return
	}

	var restConfig *rest.Config
	if restConfig, err = clientcmd.RESTConfigFromKubeConfig(
		[]byte(getProvisionDetailsResponse.ProvisionDetails.KubeconfigContent)); err != nil {
		service.logger.Error("failed to create Rest config from the given kube config", zap.Error(err))

		return
	}

	if clientset, err = kubernetes.NewForConfig(restConfig); err != nil {
		service.logger.Error("failed to create client set", zap.Error(err))

		return
	}

	return
}

func (service *k3sProvisioner) isK3SPodReady(
	ctx context.Context,
	namespace string) (bool, error) {
	watch, err := service.clientset.CoreV1().Pods(namespace).Watch(ctx, metav1.ListOptions{
		Watch:          true,
		TimeoutSeconds: &waitForDeploymentToBeReadyTimeout,
	})
	if err != nil {
		service.logger.Error("failed to retrieve K3S pod status", zap.Error(err))

		return false, err
	}

	for event := range watch.ResultChan() {
		if service, ok := event.Object.(*v1.Pod); ok {
			for _, item := range service.Status.Conditions {
				if item.Type != v1.ContainersReady {
					watch.Stop()

					return true, nil
				}
			}
		}
	}

	return false, types.NewUnknownError("failed to retrieve K3S pod status")
}

func (service *k3sProvisioner) deployHelmChart(ctx context.Context, edgeClusterID string) error {
	provisionDetails, err := service.GetProvisionDetails(ctx, &types.GetProvisionDetailsRequest{EdgeClusterID: edgeClusterID})
	if err != nil {
		return err
	}

	errorsChan := make(chan error)
	waitGroupDoneChan := make(chan bool)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()

		if err := service.helmService.InstallChart(
			provisionDetails.ProvisionDetails.KubeconfigContent,
			"portainer",
			"portainer",
			"portainer",
			"portainer",
			map[string]string{
				"set": "service.type=LoadBalancer",
			}); err != nil {
			errorsChan <- err
		}
	}()

	go func() {
		defer waitGroup.Done()

		if err := service.helmService.InstallChart(
			provisionDetails.ProvisionDetails.KubeconfigContent,
			"edgecluster",
			"edge-core",
			"decentralized-cloud",
			"edge-core",
			map[string]string{
				"set": "pod.edgeClusterType=K3S",
			}); err != nil {
			errorsChan <- err
		}
	}()

	go func() {
		waitGroup.Wait()
		close(waitGroupDoneChan)
	}()

	select {
	case <-waitGroupDoneChan:
		close(errorsChan)

		return nil

	case err := <-errorsChan:
		close(errorsChan)

		return err
	}
}
