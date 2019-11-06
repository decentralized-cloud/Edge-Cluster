package edgecluster

import (
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// KubeAdapter a struct for grtting metadata information from kubernetes
type KubeAdapter struct {
	logger    *zap.Logger
	NameSpace string
	PodName   string
}

func NewEdgeClusterKubeAdapter(logger *zap.Logger) (KubeMonitor, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	return &KubeAdapter{
		logger: logger,
	}, nil
}

//GetPods getting available pods in Kubernetes
func (adapter KubeAdapter) GetPods(clientSet *kubernetes.Clientset) (*v1.PodList, error) {
	adapter.logger.Info("get All Pods")

	pods, err := clientSet.CoreV1().Pods(adapter.NameSpace).List(metav1.ListOptions{})

	adapter.logger.Info(" number of pods is ", zap.Int("Number of pods", len(pods.Items)))

	return pods, err
}

//GetPod getting a specific pod by name in Kubernetes
func (adapter KubeAdapter) GetPod(clientSet *kubernetes.Clientset) (*v1.Pod, error) {
	adapter.logger.Info(" get pod ", zap.String("Pod Name", adapter.PodName))

	pod, err := clientSet.CoreV1().Pods(adapter.NameSpace).Get(adapter.PodName, metav1.GetOptions{})

	statusError, isStatus := err.(*errors.StatusError)

	if errors.IsNotFound(err) {
		adapter.logger.Info("Pod Info ", zap.String("Pod Name", adapter.PodName), zap.String("Pod NameSpace", adapter.NameSpace))
	} else if err != nil && isStatus {
		adapter.logger.Info("Error getting pod : ",
			zap.String("Pod Name", adapter.PodName),
			zap.String("Pod NameSpace", adapter.NameSpace),
			zap.String("Error Message", statusError.ErrStatus.Message))
	} else if err != nil {
		return nil, err
	}

	return pod, err
}
