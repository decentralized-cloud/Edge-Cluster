// Package grpc implements functions to expose edge-cluster service endpoint using GRPC protocol.
package grpc

import (
	"context"
	"fmt"
	"net"

	edgeClusterGRPCContract "github.com/decentralized-cloud/edge-cluster/contract/grpc/go"
	"github.com/decentralized-cloud/edge-cluster/services/configuration"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint"
	"github.com/decentralized-cloud/edge-cluster/services/transport"
	gokitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/micro-business/go-core/gokit/middleware"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type transportService struct {
	logger                      *zap.Logger
	configurationService        configuration.ConfigurationContract
	endpointCreatorService      endpoint.EndpointCreatorContract
	middlewareProviderService   middleware.MiddlewareProviderContract
	jwksURL                     string
	createEdgeClusterHandler    gokitgrpc.Handler
	readEdgeClusterHandler      gokitgrpc.Handler
	updateEdgeClusterHandler    gokitgrpc.Handler
	deleteEdgeClusterHandler    gokitgrpc.Handler
	searchHandler               gokitgrpc.Handler
	listEdgeClusterNodesHandler gokitgrpc.Handler
	listEdgeClusterPodsHandler  gokitgrpc.Handler
}

var Live bool
var Ready bool

func init() {
	Live = false
	Ready = false
}

// NewTransportService creates new instance of the transportService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the service that provides required configurations
// endpointCreatorService: Mandatory. Reference to the service that creates go-kit compatible endpoints
// middlewareProviderService: Mandatory. Reference to the service that provides different go-kit middlewares
// Returns the new service or error if something goes wrong
func NewTransportService(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract,
	endpointCreatorService endpoint.EndpointCreatorContract,
	middlewareProviderService middleware.MiddlewareProviderContract) (transport.TransportContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	if endpointCreatorService == nil {
		return nil, commonErrors.NewArgumentNilError("endpointCreatorService", "endpointCreatorService is required")
	}

	if middlewareProviderService == nil {
		return nil, commonErrors.NewArgumentNilError("middlewareProviderService", "middlewareProviderService is required")
	}

	jwksURL, err := configurationService.GetJwksURL()
	if err != nil {
		return nil, err
	}

	return &transportService{
		logger:                    logger,
		configurationService:      configurationService,
		endpointCreatorService:    endpointCreatorService,
		middlewareProviderService: middlewareProviderService,
		jwksURL:                   jwksURL,
	}, nil
}

// Start starts the GRPC transport service
// Returns error if something goes wrong
func (service *transportService) Start() error {
	service.setupHandlers()

	host, err := service.configurationService.GetGrpcHost()
	if err != nil {
		return err
	}

	port, err := service.configurationService.GetGrpcPort()
	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()
	edgeClusterGRPCContract.RegisterEdgeClusterServiceServer(gRPCServer, service)
	service.logger.Info("gRPC service started", zap.String("address", address))

	Live = true
	Ready = true

	err = gRPCServer.Serve(listener)

	Live = false
	Ready = false

	return err
}

// Stop stops the GRPC transport service
// Returns error if something goes wrong
func (service *transportService) Stop() error {
	return nil
}

func (service *transportService) setupHandlers() {
	endpoint := service.endpointCreatorService.CreateEdgeClusterEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("CreateEdgeCluster")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.createEdgeClusterHandler = gokitgrpc.NewServer(
		endpoint,
		decodeCreateEdgeClusterRequest,
		encodeCreateEdgeClusterResponse,
	)

	endpoint = service.endpointCreatorService.ReadEdgeClusterEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("ReadEdgeCluster")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.readEdgeClusterHandler = gokitgrpc.NewServer(
		endpoint,
		decodeReadEdgeClusterRequest,
		encodeReadEdgeClusterResponse,
	)

	endpoint = service.endpointCreatorService.UpdateEdgeClusterEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("UpdateEdgeCluster")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.updateEdgeClusterHandler = gokitgrpc.NewServer(
		endpoint,
		decodeUpdateEdgeClusterRequest,
		encodeUpdateEdgeClusterResponse,
	)

	endpoint = service.endpointCreatorService.DeleteEdgeClusterEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("DeleteEdgeCluster")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.deleteEdgeClusterHandler = gokitgrpc.NewServer(
		endpoint,
		decodeDeleteEdgeClusterRequest,
		encodeDeleteEdgeClusterResponse,
	)

	endpoint = service.endpointCreatorService.SearchEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("Search")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.searchHandler = gokitgrpc.NewServer(
		endpoint,
		decodeSearchRequest,
		encodeSearchResponse,
	)

	endpoint = service.endpointCreatorService.ListEdgeClusterNodesEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("ListEdgeClusterNodes")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.listEdgeClusterNodesHandler = gokitgrpc.NewServer(
		endpoint,
		decodeListEdgeClusterNodesRequest,
		encodeListEdgeClusterNodesResponse,
	)

	endpoint = service.endpointCreatorService.ListEdgeClusterPodsEndpoint()
	endpoint = service.middlewareProviderService.CreateLoggingMiddleware("ListEdgeClusterPods")(endpoint)
	endpoint = service.createAuthMiddleware()(endpoint)
	service.listEdgeClusterPodsHandler = gokitgrpc.NewServer(
		endpoint,
		decodeListEdgeClusterPodsRequest,
		encodeListEdgeClusterPodsResponse,
	)
}

// CreateEdgeCluster creates a new edgeCluster
// context: Mandatory. The reference to the context
// request: mandatory. The request to create a new edgeCluster
// Returns the result of creating new edgeCluster
func (service *transportService) CreateEdgeCluster(
	ctx context.Context,
	request *edgeClusterGRPCContract.CreateEdgeClusterRequest) (*edgeClusterGRPCContract.CreateEdgeClusterResponse, error) {
	_, response, err := service.createEdgeClusterHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.CreateEdgeClusterResponse), nil
}

// ReadEdgeCluster read an existing edgeCluster
// context: Mandatory. The reference to the context
// request: Mandatory. The request to read an existing edgeCluster
// Returns the result of reading an existing edgeCluster
func (service *transportService) ReadEdgeCluster(
	ctx context.Context,
	request *edgeClusterGRPCContract.ReadEdgeClusterRequest) (*edgeClusterGRPCContract.ReadEdgeClusterResponse, error) {
	_, response, err := service.readEdgeClusterHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.ReadEdgeClusterResponse), nil

}

// UpdateEdgeCluster update an existing edgeCluster
// context: Mandatory. The reference to the context
// request: Mandatory. The request to update an existing edgeCluster
// Returns the result of updateing an existing edgeCluster
func (service *transportService) UpdateEdgeCluster(
	ctx context.Context,
	request *edgeClusterGRPCContract.UpdateEdgeClusterRequest) (*edgeClusterGRPCContract.UpdateEdgeClusterResponse, error) {
	_, response, err := service.updateEdgeClusterHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.UpdateEdgeClusterResponse), nil

}

// DeleteEdgeCluster delete an existing edgeCluster
// context: Mandatory. The reference to the context
// request: Mandatory. The request to delete an existing edgeCluster
// Returns the result of deleting an existing edgeCluster
func (service *transportService) DeleteEdgeCluster(
	ctx context.Context,
	request *edgeClusterGRPCContract.DeleteEdgeClusterRequest) (*edgeClusterGRPCContract.DeleteEdgeClusterResponse, error) {
	_, response, err := service.deleteEdgeClusterHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.DeleteEdgeClusterResponse), nil

}

// Search returns the list  of edge clusters that matched the provided criteria
// context: Mandatory. The reference to the context
// request: Mandatory. The request contains the filter criteria to look for existing edge clusters
// Returns the list of edge clusters that matched the provided criteria
func (service *transportService) Search(
	ctx context.Context,
	request *edgeClusterGRPCContract.SearchRequest) (*edgeClusterGRPCContract.SearchResponse, error) {
	_, response, err := service.searchHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.SearchResponse), nil
}

// Search returns the list  of edge clusters that matched the provided criteria
// context: Mandatory. The reference to the context
// request: Mandatory. The request contains the filter criteria to look for existing edge clusters
// Returns the list of edge clusters that matched the provided criteria
func (service *transportService) ListEdgeClusterNodes(
	ctx context.Context,
	request *edgeClusterGRPCContract.ListEdgeClusterNodesRequest) (*edgeClusterGRPCContract.ListEdgeClusterNodesResponse, error) {
	_, response, err := service.listEdgeClusterNodesHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.ListEdgeClusterNodesResponse), nil
}

// Search returns the list  of edge clusters that matched the provided criteria
// context: Mandatory. The reference to the context
// request: Mandatory. The request contains the filter criteria to look for existing edge clusters
// Returns the list of edge clusters that matched the provided criteria
func (service *transportService) ListEdgeClusterPods(
	ctx context.Context,
	request *edgeClusterGRPCContract.ListEdgeClusterPodsRequest) (*edgeClusterGRPCContract.ListEdgeClusterPodsResponse, error) {
	_, response, err := service.listEdgeClusterPodsHandler.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}

	return response.(*edgeClusterGRPCContract.ListEdgeClusterPodsResponse), nil
}
