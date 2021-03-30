// Package models defines the different object models used in EdgeCluster
package models

import (
	v1 "k8s.io/api/core/v1"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	// ContextKeyParsedToken var
	ContextKeyParsedToken = contextKey("ParsedToken")
)

// ParsedToken contains details that are encoded in the received JWT token
type ParsedToken struct {
	Email string
}

// ClusterType is the edge cluster type
type ClusterType int

const (
	// K3S is an edge cluster using K3S server and agent nodes
	K3S ClusterType = iota
)

// ProvisionDetails represents the provision detail of an edge cluster
type ProvisionDetails struct {
	Ingress           []v1.LoadBalancerIngress
	Ports             []v1.ServicePort
	KubeconfigContent string
}

// EdgeCluster defines the Edge Cluster object
type EdgeCluster struct {
	ProjectID     string      `bson:"projectID" json:"projectID"`
	Name          string      `bson:"name" json:"name"`
	ClusterSecret string      `bson:"clusterSecret" json:"clusterSecret"`
	ClusterType   ClusterType `bson:"clusterType" json:"clusterType"`
}

// EdgeClusterWithCursor implements the pair of the edge cluster with a cursor that determines the
// location of the edge cluster in the repository.
type EdgeClusterWithCursor struct {
	EdgeClusterID    string
	EdgeCluster      EdgeCluster
	Cursor           string
	ProvisionDetails ProvisionDetails
}

// EdgeClusterNodeStatus is information about the current status of a node.
type EdgeClusterNodeStatus struct {
	// Node contains information about an existing edge cluster node
	Node v1.Node
}

// EdgeClusterPod is information about the current status of a pod.
type EdgeClusterPod struct {
	// Pod contains information about a deployed edge cluster node pod
	Pod v1.Pod
}
