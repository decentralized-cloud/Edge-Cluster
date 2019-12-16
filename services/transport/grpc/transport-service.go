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
	gokitEndpoint "github.com/go-kit/kit/endpoint"
	gokitgrpc "github.com/go-kit/kit/transport/grpc"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/micro-business/gokit-core/middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type transportService struct {
	logger                    *zap.Logger
	configurationService      configuration.ConfigurationContract
	endpointCreatorService    endpoint.EndpointCreatorContract
	middlewareProviderService middleware.MiddlewareProviderContract
	createEdgeClusterHandler  gokitgrpc.Handler
	readEdgeClusterHandler    gokitgrpc.Handler
	updateEdgeClusterHandler  gokitgrpc.Handler
	deleteEdgeClusterHandler  gokitgrpc.Handler
	searchHandler             gokitgrpc.Handler
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

	return &transportService{
		logger:                    logger,
		configurationService:      configurationService,
		endpointCreatorService:    endpointCreatorService,
		middlewareProviderService: middlewareProviderService,
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

// newServer creates a new GRPC server that can serve edgeCluster GRPC requests and process them
func (service *transportService) setupHandlers() {
	var createEdgeClusterEndpoint gokitEndpoint.Endpoint
	{
		createEdgeClusterEndpoint = service.endpointCreatorService.CreateEdgeClusterEndpoint()
		createEdgeClusterEndpoint = service.middlewareProviderService.CreateLoggingMiddleware("CreateEdgeCluster")(createEdgeClusterEndpoint)
		service.createEdgeClusterHandler = gokitgrpc.NewServer(
			createEdgeClusterEndpoint,
			decodeSearchRequest,
			encodeSearchResponse,
		)
	}

	var readEdgeClusterEndpoint gokitEndpoint.Endpoint
	{
		readEdgeClusterEndpoint = service.endpointCreatorService.ReadEdgeClusterEndpoint()
		readEdgeClusterEndpoint = service.middlewareProviderService.CreateLoggingMiddleware("ReadEdgeCluster")(readEdgeClusterEndpoint)
		service.readEdgeClusterHandler = gokitgrpc.NewServer(
			readEdgeClusterEndpoint,
			decodeSearchRequest,
			encodeSearchResponse,
		)
	}

	var updateEdgeClusterEndpoint gokitEndpoint.Endpoint
	{
		updateEdgeClusterEndpoint = service.endpointCreatorService.UpdateEdgeClusterEndpoint()
		updateEdgeClusterEndpoint = service.middlewareProviderService.CreateLoggingMiddleware("UpdateEdgeCluster")(updateEdgeClusterEndpoint)
		service.updateEdgeClusterHandler = gokitgrpc.NewServer(
			updateEdgeClusterEndpoint,
			decodeSearchRequest,
			encodeSearchResponse,
		)
	}

	var deleteEdgeClusterEndpoint gokitEndpoint.Endpoint
	{
		deleteEdgeClusterEndpoint = service.endpointCreatorService.DeleteEdgeClusterEndpoint()
		deleteEdgeClusterEndpoint = service.middlewareProviderService.CreateLoggingMiddleware("DeleteEdgeCluster")(deleteEdgeClusterEndpoint)
		service.deleteEdgeClusterHandler = gokitgrpc.NewServer(
			deleteEdgeClusterEndpoint,
			decodeSearchRequest,
			encodeSearchResponse,
		)
	}

	var searchEndpoint gokitEndpoint.Endpoint
	{
		searchEndpoint = service.endpointCreatorService.SearchEndpoint()
		searchEndpoint = service.middlewareProviderService.CreateLoggingMiddleware("Search")(searchEndpoint)
		service.searchHandler = gokitgrpc.NewServer(
			searchEndpoint,
			decodeSearchRequest,
			encodeSearchResponse,
		)
	}
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
// Returns the result of reading an exiting edgeCluster
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
// Returns the result of updateing an exiting edgeCluster
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
// Returns the result of deleting an exiting edgeCluster
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
