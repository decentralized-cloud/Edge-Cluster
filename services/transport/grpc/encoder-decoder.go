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
)

// decodeCreateEdgeClusterRequest decodes CreateEdgeCluster request message from GRPC object to business object
// context: Mandatory The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrong
func decodeCreateEdgeClusterRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.CreateEdgeClusterRequest)

	edgeCluster, err := mapEdgeClusterFromGrpc(castedRequest.EdgeCluster)
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
		edgeCluster, err := mapEdgeClusterToGrpc(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.CreateEdgeClusterResponse{
			Error:         edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeClusterID: castedResponse.EdgeClusterID,
			EdgeCluster:   edgeCluster,
			Cursor:        castedResponse.Cursor,
			ProvisionDetail: &edgeClusterGRPCContract.EdgeClusterProvisionDetail{
				Ingress:           mapIngressToGrpc(castedResponse.ProvisionDetails.Ingress),
				Ports:             mapPortsToGrpc(castedResponse.ProvisionDetails.Ports),
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
		edgeCluster, err := mapEdgeClusterToGrpc(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
			Error:       edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: edgeCluster,
			ProvisionDetail: &edgeClusterGRPCContract.EdgeClusterProvisionDetail{
				Ingress:           mapIngressToGrpc(castedResponse.ProvisionDetails.Ingress),
				Ports:             mapPortsToGrpc(castedResponse.ProvisionDetails.Ports),
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

	edgeCluster, err := mapEdgeClusterFromGrpc(castedRequest.EdgeCluster)
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
		edgeCluster, err := mapEdgeClusterToGrpc(castedResponse.EdgeCluster)
		if err != nil {
			return nil, err
		}

		return &edgeClusterGRPCContract.UpdateEdgeClusterResponse{
			Error:       edgeClusterGRPCContract.Error_NO_ERROR,
			EdgeCluster: edgeCluster,
			Cursor:      castedResponse.Cursor,
			ProvisionDetail: &edgeClusterGRPCContract.EdgeClusterProvisionDetail{
				Ingress:           mapIngressToGrpc(castedResponse.ProvisionDetails.Ingress),
				Ports:             mapPortsToGrpc(castedResponse.ProvisionDetails.Ports),
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

// decodeSearchRequest decodes Search request message from GRPC object to business object
// context: Optional The reference to the context
// request: Mandatory. The reference to the GRPC request
// Returns either the decoded request or error if something goes wrongw
func decodeSearchRequest(
	ctx context.Context,
	request interface{}) (interface{}, error) {
	castedRequest := request.(*edgeClusterGRPCContract.SearchRequest)
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

	return &business.SearchRequest{
		Pagination:     pagination,
		EdgeClusterIDs: castedRequest.EdgeClusterIDs,
		ProjectIDs:     castedRequest.ProjectIDs,
		SortingOptions: sortingOptions,
	}, nil
}

// encodeSearchResponse encodes Search response from business object to GRPC object
// context: Optional The reference to the context
// request: Mandatory. The reference to the business response
// Returns either the decoded response or error if something goes wrong
func encodeSearchResponse(
	ctx context.Context,
	response interface{}) (interface{}, error) {
	castedResponse := response.(*business.SearchResponse)
	if castedResponse.Err == nil {
		return &edgeClusterGRPCContract.SearchResponse{
			Error:           edgeClusterGRPCContract.Error_NO_ERROR,
			HasPreviousPage: castedResponse.HasPreviousPage,
			HasNextPage:     castedResponse.HasNextPage,
			TotalCount:      castedResponse.TotalCount,
			EdgeClusters: funk.Map(castedResponse.EdgeClusters, func(edgeCluster models.EdgeClusterWithCursor) *edgeClusterGRPCContract.EdgeClusterWithCursor {
				mappedEdgeCluster, _ := mapEdgeClusterToGrpc(edgeCluster.EdgeCluster)

				return &edgeClusterGRPCContract.EdgeClusterWithCursor{
					EdgeClusterID: edgeCluster.EdgeClusterID,
					EdgeCluster:   mappedEdgeCluster,
					Cursor:        edgeCluster.Cursor,
					ProvisionDetail: &edgeClusterGRPCContract.EdgeClusterProvisionDetail{
						Ingress:           mapIngressToGrpc(edgeCluster.ProvisionDetails.Ingress),
						Ports:             mapPortsToGrpc(edgeCluster.ProvisionDetails.Ports),
						KubeConfigContent: edgeCluster.ProvisionDetails.KubeconfigContent,
					},
				}
			}).([]*edgeClusterGRPCContract.EdgeClusterWithCursor),
		}, nil
	}

	return &edgeClusterGRPCContract.SearchResponse{
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
			Nodes: mapNodeStatusToGrpc(castedResponse.Nodes),
		}, nil
	}

	return &edgeClusterGRPCContract.ReadEdgeClusterResponse{
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

	panic("Error type undefined.")
}

func mapEdgeClusterFromGrpc(grpcEdgeCluster *edgeClusterGRPCContract.EdgeCluster) (edgeCluster models.EdgeCluster, err error) {
	var clusterType models.ClusterType

	if grpcEdgeCluster.ClusterType == edgeClusterGRPCContract.ClusterType_K3S {
		clusterType = models.K3S
	} else {
		err = fmt.Errorf("Cluster type is not supported: %v", grpcEdgeCluster.ClusterType)

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

func mapEdgeClusterToGrpc(edgeCluster models.EdgeCluster) (grpcEdgeCluster *edgeClusterGRPCContract.EdgeCluster, err error) {
	var clusterType edgeClusterGRPCContract.ClusterType

	if edgeCluster.ClusterType == models.K3S {
		clusterType = edgeClusterGRPCContract.ClusterType_K3S
	} else {
		err = fmt.Errorf("Cluster type is not supported: %v", edgeCluster.ClusterType)

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

func mapIngressToGrpc(values []v1.LoadBalancerIngress) (mappedValues []*edgeClusterGRPCContract.Ingress) {
	mappedValues = []*edgeClusterGRPCContract.Ingress{}

	for _, item := range values {
		mappedValues = append(mappedValues, &edgeClusterGRPCContract.Ingress{
			Ip:       item.IP,
			Hostname: item.Hostname,
		})
	}

	return
}

func mapPortsToGrpc(values []v1.ServicePort) (mappedValues []*edgeClusterGRPCContract.Port) {
	mappedValues = []*edgeClusterGRPCContract.Port{}

	for _, item := range values {
		mappedValues = append(mappedValues, &edgeClusterGRPCContract.Port{
			Port:    item.Port,
			Protcol: edgeClusterGRPCContract.Protocol(edgeClusterGRPCContract.Protocol_value[string(item.Protocol)]),
		})
	}

	return
}

func mapNodeStatusToGrpc(values []models.EdgeClusterNodeStatus) (mappedValues []*edgeClusterGRPCContract.EdgeClusterNodeStatus) {
	mappedValues = []*edgeClusterGRPCContract.EdgeClusterNodeStatus{}

	for _, item := range values {
		conditions := []*edgeClusterGRPCContract.EdgeClusterNodeCondition{}

		for _, condition := range item.Conditions {
			conditions = append(conditions, &edgeClusterGRPCContract.EdgeClusterNodeCondition{
				Type:               edgeClusterGRPCContract.EdgeClusterNodeConditionType(edgeClusterGRPCContract.EdgeClusterNodeConditionType_value[string(condition.Type)]),
				Status:             edgeClusterGRPCContract.EdgeClusterNodeConditionStatus(edgeClusterGRPCContract.EdgeClusterNodeConditionStatus_value[string(condition.Status)]),
				LastHeartbeatTime:  &timestamppb.Timestamp{Seconds: condition.LastHeartbeatTime.Time.Unix()},
				LastTransitionTime: &timestamppb.Timestamp{Seconds: condition.LastTransitionTime.Time.Unix()},
				Reason:             condition.Reason,
				Message:            condition.Message,
			})
		}

		addresses := []*edgeClusterGRPCContract.EdgeClusterNodeAddress{}

		for _, address := range item.Addresses {
			addresses = append(addresses, &edgeClusterGRPCContract.EdgeClusterNodeAddress{
				NodeAddressType: edgeClusterGRPCContract.EdgeClusterNodeAddressType(edgeClusterGRPCContract.EdgeClusterNodeAddressType_value[string(address.Type)]),
				Address:         address.Address,
			})
		}

		nodeInfo := edgeClusterGRPCContract.EdgeClusterNodeSystemInfo{
			MachineID:               item.NodeInfo.MachineID,
			SystemUUID:              item.NodeInfo.SystemUUID,
			BootID:                  item.NodeInfo.BootID,
			KernelVersion:           item.NodeInfo.KernelVersion,
			OsImage:                 item.NodeInfo.OSImage,
			ContainerRuntimeVersion: item.NodeInfo.ContainerRuntimeVersion,
			KubeletVersion:          item.NodeInfo.KubeletVersion,
			KubeProxyVersion:        item.NodeInfo.KubeProxyVersion,
			OperatingSystem:         item.NodeInfo.OperatingSystem,
			Architecture:            item.NodeInfo.Architecture,
		}

		mappedValues = append(mappedValues, &edgeClusterGRPCContract.EdgeClusterNodeStatus{
			Conditions: conditions,
			Addresses:  addresses,
			NodeInfo:   &nodeInfo,
		})
	}

	return
}
