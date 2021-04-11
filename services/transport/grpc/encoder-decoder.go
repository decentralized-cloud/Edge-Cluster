// Package grpc implements functions to expose edge-cluster service endpoint using GRPC protocol.
package grpc

import (
	"context"
	"fmt"

	edgeClusterGRPCContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/thoas/go-funk"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// decodeCreateEdgeClusterRequest decodes CreateEdgeCluster request message from GRPC object to business object
// context: Mandatory The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeCreateEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.CreateEdgeClusterRequest)

	edgeCluster, err := mapToEdgeCluster(castedRequest.EdgeCluster)
	if err != nil {
		return nil, err
	}

	return &business.CreateEdgeClusterRequest{
		EdgeCluster: edgeCluster,
	}, nil
}

// encodeCreateEdgeClusterResponse encodes CreateEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeCreateEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.CreateEdgeClusterResponse)

	if castedResponse.Err == nil {
		edgeCluster, err := mapFromEdgeCluster(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.CreateEdgeClusterResponse{
			Error:         edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeClusterID: castedResponse.EdgeClusterID,
			EdgeCluster:   edgeCluster,
			Cursor:        castedResponse.Cursor,
			ProvisionDetail: &edgeClusterGRPCContract.ProvisionDetail{
				LoadBalancer:      mapFromLoadBalancerStatus(castedResponse.ProvisionDetails.Service.Status.LoadBalancer),
				KubeConfigContent: castedResponse.ProvisionDetails.KubeconfigContent,
			},
		}, nil
	}

	return &edgeClusterGRPCContract.CreateEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeReadEdgeClusterRequest decodes ReadEdgeCluster request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeReadEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ReadEdgeClusterRequest)

	return &business.ReadEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
	}, nil
}

// encodeReadEdgeClusterResponse encodes ReadEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeReadEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ReadEdgeClusterResponse)

	if castedResponse.Err == nil {
		edgeCluster, err := mapFromEdgeCluster(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
			Error:       edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: edgeCluster,
			ProvisionDetail: &edgeClusterGRPCContract.ProvisionDetail{
				LoadBalancer:      mapFromLoadBalancerStatus(castedResponse.ProvisionDetails.Service.Status.LoadBalancer),
				KubeConfigContent: castedResponse.ProvisionDetails.KubeconfigContent,
			},
		}, nil
	}

	return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeUpdateEdgeClusterRequest decodes UpdateEdgeCluster request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeUpdateEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.UpdateEdgeClusterRequest)

	edgeCluster, err := mapToEdgeCluster(castedRequest.EdgeCluster)
	if err != nil {
		return nil, err
	}

	return &business.UpdateEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		EdgeCluster:   edgeCluster,
	}, nil
}

// encodeUpdateEdgeClusterResponse encodes UpdateEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeUpdateEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.UpdateEdgeClusterResponse)

	if castedResponse.Err == nil {
		edgeCluster, err := mapFromEdgeCluster(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.UpdateEdgeClusterResponse{
			Error:       edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: edgeCluster,
			Cursor:      castedResponse.Cursor,
			ProvisionDetail: &edgeClusterGRPCContract.ProvisionDetail{
				LoadBalancer:      mapFromLoadBalancerStatus(castedResponse.ProvisionDetails.Service.Status.LoadBalancer),
				KubeConfigContent: castedResponse.ProvisionDetails.KubeconfigContent,
			},
		}, nil
	}

	return &edgeClusterGRPCContract.UpdateEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeDeleteEdgeClusterRequest decodes DeleteEdgeCluster request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeDeleteEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.DeleteEdgeClusterRequest)

	return &business.DeleteEdgeClusterRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
	}, nil
}

// encodeDeleteEdgeClusterResponse encodes DeleteEdgeCluster response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeDeleteEdgeClusterResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.DeleteEdgeClusterResponse)
	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.DeleteEdgeClusterResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
		}, nil
	}

	return &edgeClusterGRPCContract.DeleteEdgeClusterResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClustersRequest decodes ListEdgeClusters request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrongw
func decodeListEdgeClustersRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClustersRequest)
	sortingOptions := []common.SortingOptionPair{}

	if len(castedRequest.SortingOptions) > 0 {
		sortingOptions = funk.Map(
			castedRequest.SortingOptions,
			func(sortingOption *edgeClusterGRPCContract.SortingOptionPair) common.SortingOptionPair {
				direction := common.Ascending

				if sortingOption.Direction == edgeClusterGRPCContract.SortingDirection_DESCENDING {
					direction = common.Descending
				}

				return common.SortingOptionPair{
					Name:      sortingOption.Name,
					Direction: direction,
				}
			}).([]common.SortingOptionPair)
	}

	pagination := common.Pagination{}

	if castedRequest.Pagination.HasAfter {
		pagination.After = &castedRequest.Pagination.After
	}

	if castedRequest.Pagination.HasFirst {
		first := int(castedRequest.Pagination.First)
		pagination.First = &first
	}

	if castedRequest.Pagination.HasBefore {
		pagination.Before = &castedRequest.Pagination.Before
	}

	if castedRequest.Pagination.HasLast {
		last := int(castedRequest.Pagination.Last)
		pagination.Last = &last
	}

	return &business.ListEdgeClustersRequest{
		Pagination:     pagination,
		EdgeClusterIDs: castedRequest.EdgeClusterIDs,
		ProjectIDs:     castedRequest.ProjectIDs,
		SortingOptions: sortingOptions,
	}, nil
}

// encodeListEdgeClustersResponse encodes ListEdgeClusters response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClustersResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClustersResponse)
	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClustersResponse{
			Error:           edgeClusterGRPCContract.Error_NO_ERROR,
			HasPreviousPage: castedResponse.HasPreviousPage,
			HasNextPage:     castedResponse.HasNextPage,
			TotalCount:      castedResponse.TotalCount,
			EdgeClusters: funk.Map(castedResponse.EdgeClusters, func(edgeCluster models.EdgeClusterWithCursor) *edgeClusterGRPCContract.EdgeClusterWithCursor {
				mappedEdgeCluster, _ := mapFromEdgeCluster(edgeCluster.EdgeCluster)

				return &edgeClusterGRPCContract.EdgeClusterWithCursor{
					EdgeClusterID: edgeCluster.EdgeClusterID,
					EdgeCluster:   mappedEdgeCluster,
					Cursor:        edgeCluster.Cursor,
					ProvisionDetail: &edgeClusterGRPCContract.ProvisionDetail{
						LoadBalancer:      mapFromLoadBalancerStatus(edgeCluster.ProvisionDetails.Service.Status.LoadBalancer),
						KubeConfigContent: edgeCluster.ProvisionDetails.KubeconfigContent,
					},
				}
			}).([]*edgeClusterGRPCContract.EdgeClusterWithCursor),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClustersResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClusterNodesRequest decodes ListEdgeClusterNodes request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeListEdgeClusterNodesRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClusterNodesRequest)

	return &business.ListEdgeClusterNodesRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
	}, nil
}

// encodeListEdgeClusterNodesResponse encodes ListEdgeClusterNodes response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClusterNodesResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClusterNodesResponse)

	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClusterNodesResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
			Nodes: mapFromNodeStatus(castedResponse.Nodes),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClusterNodesResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClusterPodsRequest decodes ListEdgeClusterPods request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeListEdgeClusterPodsRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClusterPodsRequest)

	return &business.ListEdgeClusterPodsRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		Namespace:     castedRequest.Namespace,
		NodeName:      castedRequest.NodeName,
	}, nil
}

// encodeListEdgeClusterPodsResponse encodes ListEdgeClusterPods response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClusterPodsResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClusterPodsResponse)

	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClusterPodsResponse{
			Error: edgeClusterGRPCContract.Error_NO_ERROR,
			Pods:  mapFromPods(castedResponse.Pods),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClusterPodsResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

// decodeListEdgeClusterServicesRequest decodes ListEdgeClusterServices request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeListEdgeClusterServicesRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.ListEdgeClusterServicesRequest)

	return &business.ListEdgeClusterServicesRequest{
		EdgeClusterID: castedRequest.EdgeClusterID,
		Namespace:     castedRequest.Namespace,
	}, nil
}

// encodeListEdgeClusterServicesResponse encodes ListEdgeClusterServices response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeListEdgeClusterServicesResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.ListEdgeClusterServicesResponse)

	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.ListEdgeClusterServicesResponse{
			Error:    edgeClusterGRPCContract.Error_NO_ERROR,
			Services: mapFromServices(castedResponse.Services),
		}, nil
	}

	return &edgeClusterGRPCContract.ListEdgeClusterServicesResponse{
		Error:        mapError(castedResponse.Err),
		ErrorMessage: castedResponse.Err.Error(),
	}, nil
}

func mapError(err error) edgeClusterGRPCContract.Error {
	if commonErrors.IsUnknownError(err) {
		return edgeClusterGRPCContract.Error_UNKNOWN
	}

	if commonErrors.IsAlreadyExistsError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_ALREADY_EXISTS
	}

	if commonErrors.IsNotFoundError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_NOT_FOUND
	}

	if commonErrors.IsArgumentNilError(err) || commonErrors.IsArgumentError(err) {
		return edgeClusterGRPCContract.Error_BAD_REQUEST
	}

	return edgeClusterGRPCContract.Error_UNKNOWN
}

func mapToEdgeCluster(grpcEdgeCluster *edgeClusterGRPCContract.EdgeCluster) (edgeCluster models.EdgeCluster, err error) {
	var clusterType models.ClusterType

	if grpcEdgeCluster.ClusterType == edgeClusterGRPCContract.ClusterType_K3S {
		clusterType = models.K3S
	} else {
		err = fmt.Errorf("cluster type is not supported: %v", grpcEdgeCluster.ClusterType)

		return
	}

	edgeCluster = models.EdgeCluster{
		ProjectID:     grpcEdgeCluster.ProjectID,
		Name:          grpcEdgeCluster.Name,
		ClusterSecret: grpcEdgeCluster.ClusterSecret,
		ClusterType:   clusterType,
	}

	return
}

func mapFromEdgeCluster(edgeCluster models.EdgeCluster) (grpcEdgeCluster *edgeClusterGRPCContract.EdgeCluster, err error) {
	var clusterType edgeClusterGRPCContract.ClusterType

	if edgeCluster.ClusterType == models.K3S {
		clusterType = edgeClusterGRPCContract.ClusterType_K3S
	} else {
		err = fmt.Errorf("cluster type is not supported: %v", edgeCluster.ClusterType)

		return
	}

	grpcEdgeCluster = &edgeClusterGRPCContract.EdgeCluster{
		ProjectID:     edgeCluster.ProjectID,
		Name:          edgeCluster.Name,
		ClusterSecret: edgeCluster.ClusterSecret,
		ClusterType:   clusterType,
	}

	return
}

func mapFromNodeStatus(nodes []models.EdgeClusterNode) (mappedValues []*edgeClusterGRPCContract.EdgeClusterNode) {
	mappedValues = []*edgeClusterGRPCContract.EdgeClusterNode{}

	for _, node := range nodes {
		conditions := []*edgeClusterGRPCContract.NodeCondition{}

		for _, condition := range node.Node.Status.Conditions {
			conditions = append(conditions, &edgeClusterGRPCContract.NodeCondition{
				Type:               edgeClusterGRPCContract.NodeConditionType(edgeClusterGRPCContract.NodeConditionType_value[string(condition.Type)]),
				Status:             edgeClusterGRPCContract.ConditionStatus(edgeClusterGRPCContract.ConditionStatus_value[string(condition.Status)]),
				LastHeartbeatTime:  &timestamppb.Timestamp{Seconds: condition.LastHeartbeatTime.Time.Unix()},
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			})
		}

		addresses := []*edgeClusterGRPCContract.EdgeClusterNodeAddress{}

		for _, address := range node.Node.Status.Addresses {
			addresses = append(addresses, &edgeClusterGRPCContract.EdgeClusterNodeAddress{
				NodeAddressType: edgeClusterGRPCContract.NodeAddressType(edgeClusterGRPCContract.NodeAddressType_value[string(address.Type)]),
				Address:         address.Address,
			})
		}

		nodeInfo := edgeClusterGRPCContract.NodeSystemInfo{
			MachineID:               node.Node.Status.NodeInfo.MachineID,
			SystemUUID:              node.Node.Status.NodeInfo.SystemUUID,
			BootID:                  node.Node.Status.NodeInfo.BootID,
			KernelVersion:           node.Node.Status.NodeInfo.KernelVersion,
			OsImage:                 node.Node.Status.NodeInfo.OSImage,
			ContainerRuntimeVersion: node.Node.Status.NodeInfo.ContainerRuntimeVersion,
			KubeletVersion:          node.Node.Status.NodeInfo.KubeletVersion,
			KubeProxyVersion:        node.Node.Status.NodeInfo.KubeProxyVersion,
			OperatingSystem:         node.Node.Status.NodeInfo.OperatingSystem,
			Architecture:            node.Node.Status.NodeInfo.Architecture,
		}

		mappedValues = append(mappedValues, &edgeClusterGRPCContract.EdgeClusterNode{
			Metadata: &edgeClusterGRPCContract.ObjectMeta{
				Id:        string(node.Node.UID),
				Name:      node.Node.Name,
				Namespace: node.Node.Namespace,
			},
			Status: &edgeClusterGRPCContract.NodeStatus{
				Conditions: conditions,
				Addresses:  addresses,
				NodeInfo:   &nodeInfo,
			},
		})
	}

	return
}

func mapFromPods(pods []models.EdgeClusterPod) (mappedValues []*edgeClusterGRPCContract.EdgeClusterPod) {
	mappedValues = []*edgeClusterGRPCContract.EdgeClusterPod{}

	for _, pod := range pods {
		conditions := []*edgeClusterGRPCContract.PodCondition{}

		for _, condition := range pod.Pod.Status.Conditions {
			conditions = append(conditions, &edgeClusterGRPCContract.PodCondition{
				Type:               edgeClusterGRPCContract.PodConditionType(edgeClusterGRPCContract.PodConditionType_value[string(condition.Type)]),
				Status:             edgeClusterGRPCContract.ConditionStatus(edgeClusterGRPCContract.ConditionStatus_value[string(condition.Status)]),
				LastProbeTime:      &timestamppb.Timestamp{Seconds: condition.LastProbeTime.Time.Unix()},
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			})
		}

		mappedValues = append(mappedValues, &edgeClusterGRPCContract.EdgeClusterPod{
			Metadata: mapFromObjectMeta(pod.Pod.ObjectMeta),
			Spec: &edgeClusterGRPCContract.PodSpec{
				NodeName: pod.Pod.Spec.NodeName,
			},
			Status: &edgeClusterGRPCContract.PodStatus{
				HostIP:     pod.Pod.Status.HostIP,
				PodIP:      pod.Pod.Status.PodIP,
				Conditions: conditions,
			},
		})
	}

	return
}

func mapFromServices(services []models.EdgeClusterService) (mappedValues []*edgeClusterGRPCContract.EdgeClusterService) {
	mappedValues = []*edgeClusterGRPCContract.EdgeClusterService{}

	for _, service := range services {
		conditions := []*edgeClusterGRPCContract.ServiceCondition{}

		for _, condition := range service.Service.Status.Conditions {
			conditions = append(conditions, &edgeClusterGRPCContract.ServiceCondition{
				Type:               condition.Type,
				Status:             edgeClusterGRPCContract.ConditionStatus(edgeClusterGRPCContract.ConditionStatus_value[string(condition.Status)]),
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			})
		}

		ports := []*edgeClusterGRPCContract.ServicePort{}
		for _, port := range service.Service.Spec.Ports {
			ports = append(ports, &edgeClusterGRPCContract.ServicePort{
				Name:       port.Name,
				Protcol:    edgeClusterGRPCContract.Protocol(edgeClusterGRPCContract.Protocol_value[string(port.Protocol)]),
				Port:       port.Port,
				TargetPort: port.TargetPort.String(),
				NodePort:   port.NodePort,
			})
		}

		mappedValues = append(mappedValues, &edgeClusterGRPCContract.EdgeClusterService{
			Metadata: mapFromObjectMeta(service.Service.ObjectMeta),
			Spec: &edgeClusterGRPCContract.ServiceSpec{
				Ports:        ports,
				Type:         edgeClusterGRPCContract.ServiceType(edgeClusterGRPCContract.ServiceType_value[string(service.Service.Spec.Type)]),
				ClusterIPs:   service.Service.Spec.ClusterIPs,
				ExternalIPs:  service.Service.Spec.ExternalIPs,
				ExternalName: service.Service.Spec.ExternalName,
			},
			Status: &edgeClusterGRPCContract.ServiceStatus{
				LoadBalancer: mapFromLoadBalancerStatus(service.Service.Status.LoadBalancer),
				Conditions:   conditions,
			},
		})
	}

	return
}

func mapFromObjectMeta(objectMeta metav1.ObjectMeta) *edgeClusterGRPCContract.ObjectMeta {
	return &edgeClusterGRPCContract.ObjectMeta{
		Id:        string(objectMeta.UID),
		Name:      objectMeta.Name,
		Namespace: objectMeta.Namespace,
	}
}

func mapFromLoadBalancerStatus(loadBalancerStatus v1.LoadBalancerStatus) (mappedValues *edgeClusterGRPCContract.LoadBalancerStatus) {
	mappedValue := edgeClusterGRPCContract.LoadBalancerStatus{}
	mappedValue.Ingress = []*edgeClusterGRPCContract.LoadBalancerIngress{}

	for _, ingress := range loadBalancerStatus.Ingress {
		mappedValue.Ingress = append(mappedValue.Ingress, &edgeClusterGRPCContract.LoadBalancerIngress{
			Ip:         ingress.IP,
			Hostname:   ingress.Hostname,
			PortStatus: mapFromPortStatus(ingress.Ports),
		})
	}

	return
}

func mapFromPortStatus(ports []v1.PortStatus) (portStatus []*edgeClusterGRPCContract.PortStatus) {
	portStatus = []*edgeClusterGRPCContract.PortStatus{}

	for _, port := range ports {
		error := ""
		if port.Error != nil {
			error = *port.Error
		}

		portStatus = append(portStatus, &edgeClusterGRPCContract.PortStatus{
			Port:     port.Port,
			Protocol: edgeClusterGRPCContract.Protocol(edgeClusterGRPCContract.Protocol_value[string(port.Protocol)]),
			Error:    error,
		})
	}

	return
}
