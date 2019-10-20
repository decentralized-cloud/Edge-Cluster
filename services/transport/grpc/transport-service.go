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
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type transportService struct {
	logger                   *zap.Logger
	configurationService     configuration.ConfigurationContract
	endpointCreatorService   endpoint.EndpointCreatorContract
	createEdgeClusterHandler gokitgrpc.Handler
	readEdgeClusterHandler   gokitgrpc.Handler
	updateEdgeClusterHandler gokitgrpc.Handler
	deleteEdgeClusterHandler gokitgrpc.Handler
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
// Returns the new service or error if something goes wrong
func NewTransportService(
	logger *zap.Logger,
	configurationService configuration.ConfigurationContract,
	endpointCreatorService endpoint.EndpointCreatorContract) (transport.TransportContract, error) {
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
