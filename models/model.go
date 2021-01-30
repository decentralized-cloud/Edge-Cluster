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

// Ingress represents the status of a load-balancer ingress point
type Ingress struct {
	// IP is set for load-balancer ingress points that are IP based
	// (typically GCE or OpenStack load-balancers)
	// +optional
	IP string

	// Hostname is set for load-balancer ingress points that are DNS based
	// (typically AWS load-balancers)
	// +optional
	Hostname string
}

// Port contains information on service's port.
type Port struct {
	// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP".
	// Default is TCP.
	// +optional
	Protocol v1.Protocol

	// The port that will be exposed by this service.
	Port int32
}

// EdgeClusterWithCursor implements the pair of the edge cluster with a cursor that determines the
// location of the edge cluster in the repository.
type EdgeClusterWithCursor struct {
	EdgeClusterID string
	EdgeCluster   EdgeCluster
	Cursor        string
	Ingress       []Ingress
	Ports         []Port
}
