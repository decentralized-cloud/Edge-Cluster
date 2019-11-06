package edgecluster

import (
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

type edgeClusterDeployment struct {
	deployment EdgeClusterDeploymentDetail
}

func NewEdgeClusterDeployment(logger *zap.Logger) (EdgeClusterAdapter, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &edgeClusterDeployment.deployment{
		logger: logger,
	}, nil
}

//EdgeClusterDeploymentDetail microbusiness adapter for deployment
type EdgeClusterDeploymentDetail struct {
	logger         *zap.Logger
	Metaobject     MetaData
	AppName        string
	IPAddress      string
	Replicas       int32
	ContainerName  string
	ContainerImage string
	ConfigName     string
}

//Create deployment
func (edge EdgeClusterDeploymentDetail) Create(clientSet *kubernetes.Clientset) error {
	edge.logger.Info("call Create from deployment")
	deploymentClient := clientSet.AppsV1().Deployments(edge.Metaobject.NameSpace)

	deploymentConfig := edge.populateDeploymentConfigValue()

	edge.logger.Info("creating ...")

	result, err := deploymentClient.Create(deploymentConfig)

	edge.logger.Info("created deployment", zap.String("Deployment Name", result.GetObjectMeta().GetName()))

	return err
}

//UpdateWithRetry deployment
func (edge EdgeClusterDeploymentDetail) UpdateWithRetry(clientSet *kubernetes.Clientset) error {
	edge.logger.Info("call Update from deployment")

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

//Delete deployment
func (edge EdgeClusterDeploymentDetail) Delete(clientSet *kubernetes.Clientset) error {
	edge.logger.Info("call Delete from deployment")

	deleteClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	deletePolicy := metav1.DeletePropagationForeground

	err := deleteClient.Delete(edge.Metaobject.Name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})

	return err
}

//PopulateDeploymentConfigValue create spec object for deployment
func (edge EdgeClusterDeploymentDetail) populateDeploymentConfigValue() *appsv1.Deployment {
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: edge.Metaobject.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: Int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": edge.AppName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": edge.AppName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  edge.ContainerName,
							Image: edge.ContainerImage,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
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
