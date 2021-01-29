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
				Ingress: mapIngressToGrpc(castedResponse.Ingress),
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
				Ingress: mapIngressToGrpc(castedResponse.Ingress),
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
				Ingress: mapIngressToGrpc(castedResponse.Ingress),
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
		TenantIDs:      castedRequest.TenantIDs,
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
						Ingress: mapIngressToGrpc(edgeCluster.Ingress),
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

func mapError(err error) edgeClusterGRPCContract.Error {
	if business.IsUnknownError(err) {
		return edgeClusterGRPCContract.Error_UNKNOWN
	}

	if business.IsEdgeClusterAlreadyExistsError(err) {
		return edgeClusterGRPCContract.Error_EDGE_CLUSTER_ALREADY_EXISTS
	}

	if business.IsEdgeClusterNotFoundError(err) {
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
		TenantID:      grpcEdgeCluster.TenantID,
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
		TenantID:      edgeCluster.TenantID,
		Name:          edgeCluster.Name,
		ClusterSecret: edgeCluster.ClusterSecret,
		ClusterType:   clusterType,
	}

	return
}

func mapIngressToGrpc(ingress []v1.LoadBalancerIngress) (mappedIngress []*edgeClusterGRPCContract.Ingress) {
	mappedIngress = []*edgeClusterGRPCContract.Ingress{}

	for _, item := range ingress {
		ipPort := edgeClusterGRPCContract.Ingress{
			Ip:       item.IP,
			Hostname: item.Hostname,
			Ports:    []*edgeClusterGRPCContract.PortStatus{},
		}

		ipPort.Ports = []*edgeClusterGRPCContract.PortStatus{}

		for _, port := range item.Ports {
			ipPort.Ports = append(
				ipPort.Ports,
				&edgeClusterGRPCContract.PortStatus{
					Port:    port.Port,
					Protcol: edgeClusterGRPCContract.Protocol(edgeClusterGRPCContract.Protocol_value[string(port.Protocol)])},
			)
		}

		mappedIngress = append(mappedIngress, &ipPort)
	}

	return
}
