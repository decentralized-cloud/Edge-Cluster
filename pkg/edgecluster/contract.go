package edgecluster

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

// KubeCRUDAdapter microbusiness Kubernetes interface for create,update and delete
type EdgeClusterAdapter interface {

	//Create create a new edge cluster.
	//ClientSet: kubernetes client object for create operation
	Create(ctx context.Context, clientSet *kubernetes.Clientset) error

	//UpdateWithRetry update an existing cluster.
	//ClientSet: kubernetes client object for update operation
	UpdateWithRetry(ctx context.Context, clientSet *kubernetes.Clientset) error

	//Replace replace an existing cluster.
	//ClientSet: kubernetes client object for replace operation
	//Replace(ctx context.Context, clientSet *kubernetes.Clientset) error

	//Delete delete an existing cluster.
	//ClientSet: kubernetes client object for delete operation
	Delete(ctx context.Context, clientSet *kubernetes.Clientset) error
}

// KubeMonitor microbusiness Kubernetes interface for monitoring
// clientSet: kubernetes client object for crup operation
type KubeMonitor interface {
	GetPods(clientSet *kubernetes.Clientset) (*v1.PodList, error)
	GetPod(clientSet *kubernetes.Clientset) (*v1.Pod, error)
}
