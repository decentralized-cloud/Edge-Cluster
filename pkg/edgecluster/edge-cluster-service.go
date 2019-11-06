package edgecluster

import (
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/util/retry"
)

type edgeClusterService struct {
	service EdgeClusterServiceDetail
}

func NewEdgeClusterService(logger *zap.Logger) (EdgeClusterAdapter, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &edgeClusterService.service{
		logger: logger,
	}, nil
}

//EdgeClusterServiceDetail micro business adapter for service
type EdgeClusterServiceDetail struct {
	logger         *zap.Logger
	Metaobject     MetaData
	AppName        string
	IPAddress      string
	Ports          []apiv1.ServicePort
	Replicas       int32
	ContainerName  string
	ContainerImage string
	ConfigName     string
	Selector       map[string]string
	LabelName      string
}

//Create service
func (edge EdgeClusterServiceDetail) Create(clientSet *kubernetes.Clientset) error {
	edge.logger.Info("call Create from service")

	serviceDeployment := clientSet.CoreV1().Services(apiv1.NamespaceDefault)

	serviceConfig := edge.populateDeploymentConfigValue()

	edge.logger.Info("creating ....")

	result, err := serviceDeployment.Create(serviceConfig)

	edge.logger.Info("created service ", zap.String("Service Name", result.GetObjectMeta().GetName()))

	return err
}

//Update service
func (edge EdgeClusterServiceDetail) UpdateWithRetry(clientSet *kubernetes.Clientset) error {
	edge.logger.Info("call Update from service")

	updateClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

		result, getErr := updateClient.Get(edge.Metaobject.Name, metav1.GetOptions{})
		if getErr != nil {
			edge.logger.Info("Failed to get the deployment for update..")
		}

		//Do what need to be updated
		//Todo complete the whole necessary fileds
		result.Spec.Replicas = Int32Ptr(edge.Replicas)
		result.Spec.Template.Spec.Containers[0].Image = edge.ContainerImage

		_, updateErr := updateClient.Update(result)

		return updateErr
	})

	edge.logger.Info("Update complete ...")

	return retryErr
}

//Delete service
func (edge EdgeClusterServiceDetail) Delete(clientSet *kubernetes.Clientset) error {
	edge.logger.Info("call Delete from service")

	deleteClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground

	err := deleteClient.Delete(edge.Metaobject.Name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	return err
}

func (edge EdgeClusterServiceDetail) populateDeploymentConfigValue() *apiv1.Service {
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      edge.Metaobject.Name,
			Namespace: apiv1.NamespaceDefault,
			Labels: map[string]string{
				"k8s-app": edge.LabelName,
			},
		},
		Spec: apiv1.ServiceSpec{
			Ports:     edge.Ports,
			Selector:  edge.Selector,
			ClusterIP: edge.IPAddress,
		},
	}

	return service
}
