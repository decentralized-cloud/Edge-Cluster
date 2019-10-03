// Package grpctransport implements functions to expose EdgeCluster service endpoint using GRPC protocol.
package grpctransport

import (
	"context"
	"fmt"
	"net"

	edgeClusterGRPCContract "github.com/decentralized-cloud/edge-cluster-contract/grpc"
	configurationServiceContract "github.com/decentralized-cloud/edge-cluster/services/configuration/contract"
	endpointContract "github.com/decentralized-cloud/edge-cluster/services/endpoint/contract"
	transportContract "github.com/decentralized-cloud/edge-cluster/transport/contract"
	gokitgrpc "github.com/go-kit/kit/transport/grpc"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type transportService struct {
	logger                   *zap.Logger
	endpointCreatorService   endpointContract.EndpointCreatorContract
	configurationService     configurationServiceContract.ConfigurationServiceContract
	createEdgeClusterHandler gokitgrpc.Handler
	readEdgeClusterHandler   gokitgrpc.Handler
	updateEdgeClusterHandler gokitgrpc.Handler
	deleteEdgeClusterHandler gokitgrpc.Handler
}

// NewTransportService creates new instance of the GRPCService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the service that provides required configurations
// endpointCreatorService: Mandatory. Reference to the service that creates go-kit compatible endpoints
// Returns the new service or error if something goes wrong
func NewTransportService(
	logger *zap.Logger,
	configurationService configurationServiceContract.ConfigurationServiceContract,
	endpointCreatorService endpointContract.EndpointCreatorContract) (transportContract.TransportServiceContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	if endpointCreatorService == nil {
		return nil, commonErrors.NewArgumentNilError("endpointCreatorService", "endpointCreatorService is required")
	}

	return &transportService{
		logger:                 logger,
		configurationService:   configurationService,
		endpointCreatorService: endpointCreatorService,
	}, nil
}

// Start starts the GRPC transport service
// Returns error if something goes wrong
func (service *transportService) Start() error {
	service.setupHandlers()

	portNumber, err := service.configurationService.GetGRPCPort()
	if err != nil {
		return err
	}

	host, err := service.configurationService.GetGRPCHost()
	if err != nil {
		return err
	}

	address := fmt.Sprintf("%s:%d", host, portNumber)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	gRPCServer := grpc.NewServer()
	edgeClusterGRPCContract.RegisterEdgeClusterServiceServer(gRPCServer, service)
	service.logger.Info("gRPC server started", zap.String("address", address))

	return gRPCServer.Serve(listener)
}

// Stop stops the GRPC transport service
// Returns error if something goes wrong
func (service *transportService) Stop() error {
	return nil
}

// newServer creates a new GRPC server that can serve edgeCluster GRPC requests and process them
func (service *transportService) setupHandlers() {
	service.createEdgeClusterHandler = gokitgrpc.NewServer(
		service.endpointCreatorService.CreateEdgeClusterEndpoint(),
		decodeCreateEdgeClusterRequest,
		encodeCreateEdgeClusterResponse,
	)

	service.readEdgeClusterHandler = gokitgrpc.NewServer(
		service.endpointCreatorService.ReadEdgeClusterEndpoint(),
		decodeReadEdgeClusterRequest,
		encodeReadEdgeClusterResponse,
	)

	service.updateEdgeClusterHandler = gokitgrpc.NewServer(
		service.endpointCreatorService.UpdateEdgeClusterEndpoint(),
		decodeUpdateEdgeClusterRequest,
		encodeUpdateEdgeClusterResponse,
	)

	service.deleteEdgeClusterHandler = gokitgrpc.NewServer(
		service.endpointCreatorService.DeleteEdgeClusterEndpoint(),
		decodeDeleteEdgeClusterRequest,
		encodeDeleteEdgeClusterResponse,
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
