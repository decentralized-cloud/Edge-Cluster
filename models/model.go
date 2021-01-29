// Package models defines the different object models used in EdgeCluster
package models

import v1 "k8s.io/api/core/v1"

type ClusterType int

const (
	K3S ClusterType = iota
)

// EdgeCluster defines the Edge Cluster object
type EdgeCluster struct {
	TenantID      string      `bson:"tenantID" json:"tenantID"`
	Name          string      `bson:"name" json:"name"`
	ClusterSecret string      `bson:"clusterSecret" json:"clusterSecret"`
	ClusterType   ClusterType `bson:"clusterType" json:"clusterType"`
}

// EdgeClusterWithCursor implements the pair of the edge cluster with a cursor that determines the
// location of the edge cluster in the repository.
type EdgeClusterWithCursor struct {
	EdgeClusterID string
	EdgeCluster   EdgeCluster
	Ingress       []v1.LoadBalancerIngress
	Cursor        string
}
