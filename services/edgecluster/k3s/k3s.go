// Package k3s provides functionality to provision a K3S edge cluster type and manage them
package k3s

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"

	"github.com/decentralized-cloud/edge-cluster/models"
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
	"k8s.io/client-go/tools/remotecommand"
	"k8s.io/client-go/util/retry"
)

const (
	containerName      = "k3sserver"
	containerImage     = "rancher/k3s:v1.20.0-k3s2"
	k3sPort            = 6443
	internalName       = "k3s"
	kubeconfigFilePath = "/etc/rancher/k3s/k3s.yaml"
)

var deploymentReplica int32 = 1
var waitForDeploymentToBeReadyTimeout int64 = 60

type k3sProvisioner struct {
	logger        *zap.Logger
	clientset     *kubernetes.Clientset
	k8sRestConfig *rest.Config
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

	var clientset *kubernetes.Clientset
	var err error

	if clientset, err = kubernetes.NewForConfig(k8sRestConfig); err != nil {
		return nil, types.NewUnknownErrorWithError("Failed to create client set", err)
	}

	return &k3sProvisioner{
		logger:        logger,
		clientset:     clientset,
		k8sRestConfig: k8sRestConfig,
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
				service.logger.Error("Failed to update the edge cluster", zap.Error(err))

				return
			}

			deployment.Spec.Template.Spec, err = service.getDeploymentSpec(ctx, namespace, request.ClusterSecret)
			if err != nil {
				return err
			}

			if _, err = client.Update(ctx, deployment, metav1.UpdateOptions{}); err != nil {
				service.logger.Error("Failed to update the edge custer", zap.Error(err))

				return
			}

			return
		})

	response = &types.UpdateProvisionResponse{}

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
		service.logger.Error("Failed to delete namespace", zap.Error(err))

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

	ingress, ports, err := service.getProvisionDetailsServiceDetails(ctx, namespace)
	if err != nil {
		service.logger.Error("Failed to get service details", zap.Error(err))

		return
	}

	kubeconfigContent, err := service.getProvisionDetailsKubeConfigContent(ctx, namespace)
	if err != nil {
		service.logger.Error("Failed to get kubeconfig content", zap.Error(err))

		return
	}

	for _, item := range ingress {
		if item.IP != "" {
			kubeconfigContent = strings.Replace(kubeconfigContent, "127.0.0.1", item.IP, -1)
		} else if item.Hostname != "" {
			kubeconfigContent = strings.Replace(kubeconfigContent, "127.0.0.1", item.Hostname, -1)
		} else {
			kubeconfigContent = strings.Replace(kubeconfigContent, "127.0.0.1", "BLANK", -1)
		}
	}

	for _, item := range ports {
		kubeconfigContent = strings.Replace(kubeconfigContent, fmt.Sprintf("%d", k3sPort), fmt.Sprintf("%d", item.Port), -1)
	}

	response = &types.GetProvisionDetailsResponse{
		ProvisionDetails: models.ProvisionDetails{
			Ingress:           ingress,
			Ports:             ports,
			KubeconfigContent: kubeconfigContent,
		}}

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
				"Failed to create namespace",
				zap.Error(err),
				zap.String("namespace", namespace))

			return
		}

	} else if err != nil {
		service.logger.Error("Failed to validate the requested namespace", zap.Error(err), zap.String("namespace", namespace))

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
			Template: apiv1.PodTemplateSpec{
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
		service.logger.Error("Failed to create edge cluster", zap.Error(err), zap.Any("Config", deploymentConfig))
	}

	return
}

func (service *k3sProvisioner) createService(ctx context.Context, namespace string) (err error) {
	serviceDeployment := service.clientset.CoreV1().Services(namespace)

	servicePorts := []v1.ServicePort{
		{
			Name:       internalName,
			Protocol:   apiv1.ProtocolTCP,
			Port:       k3sPort,
			TargetPort: intstr.FromInt(k3sPort),
		},
	}

	serviceSelector := map[string]string{
		internalName: internalName,
	}

	serviceConfig := &apiv1.Service{
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

	if _, err = serviceDeployment.Create(ctx, serviceConfig, metav1.CreateOptions{}); err != nil {
		service.logger.Error("Failed to create service", zap.Error(err), zap.Any("Config", serviceConfig))

		return
	}

	return
}

func getNamespace(edgeClusterID string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(edgeClusterID)))
}

func (service *k3sProvisioner) getDeploymentSpec(ctx context.Context, namespace string, k3SClusterSecret string) (apiv1.PodSpec, error) {
	advertiseAddress, err := service.getAdvertiseAddress(ctx, namespace)
	if err != nil {
		return apiv1.PodSpec{}, err
	}

	return apiv1.PodSpec{
		Containers: []apiv1.Container{
			{
				Name:  containerName,
				Image: containerImage,
				Args: []string{
					"server",
					fmt.Sprintf("--advertise-address=%s", advertiseAddress),
				},
				Env: []apiv1.EnvVar{
					{Name: "K3S_CLUSTER_SECRET", Value: k3SClusterSecret},
				},
				Ports: []apiv1.ContainerPort{
					{
						Name:          internalName,
						ContainerPort: k3sPort,
					},
				},
			},
		},
	}, nil
}

func (service *k3sProvisioner) getProvisionDetailsServiceDetails(
	ctx context.Context,
	namespace string) (ingress []models.Ingress, ports []models.Port, err error) {

	var serviceInfo *v1.Service

	if serviceInfo, err = service.clientset.CoreV1().Services(namespace).Get(ctx, internalName, metav1.GetOptions{}); err != nil {
		service.logger.Error("Failed to fetch service info", zap.Error(err))

		return
	}

	for _, item := range serviceInfo.Status.LoadBalancer.Ingress {
		ingress = append(
			ingress,
			models.Ingress{
				IP:       item.IP,
				Hostname: item.Hostname})
	}

	for _, port := range serviceInfo.Spec.Ports {
		ports = append(
			ports,
			models.Port{
				Protocol: port.Protocol,
				Port:     port.Port})
	}

	return
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
		err = types.NewUnknownErrorWithError("Failed to retrieve KubeConfig content.", err)

		return
	}

	output := &bytes.Buffer{}

	if err = executor.Stream(remotecommand.StreamOptions{Stdout: output}); err != nil {
		err = types.NewUnknownErrorWithError("Failed to retrieve KubeConfig content", err)

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
		service.logger.Error("Failed to retrieve advertise address", zap.Error(err))

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

	return "", types.NewUnknownError("Failed to retrieve advertise address")
}
