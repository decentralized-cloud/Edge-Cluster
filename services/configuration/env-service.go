// Package configuration implements configuration service required by the edge-cluster service
package configuration

import (
	"os"
	"strconv"
	"strings"
)

type envConfigurationService struct {
}

// NewEnvConfigurationService creates new instance of the EnvConfigurationService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewEnvConfigurationService() (ConfigurationContract, error) {
	return &envConfigurationService{}, nil
}

// GetGrpcHost returns gRPC host name
// Returns the gRPC host name or error if something goes wrong
func (service *envConfigurationService) GetGrpcHost() (string, error) {
	return os.Getenv("GRPC_HOST"), nil
}

// GetGrpcPort returns gRPC port number
// Returns the gRPC port number or error if something goes wrong
func (service *envConfigurationService) GetGrpcPort() (int, error) {
	valueStr := os.Getenv("GRPC_PORT")
	if strings.Trim(valueStr, " ") == "" {
		return 0, NewUnknownError("GRPC_PORT is required")
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert GRPC_PORT to integer", err)
	}

	return value, nil
}

// GetHttpHost returns HTTP host name
// Returns the HTTP host name or error if something goes wrong
func (service *envConfigurationService) GetHttpHost() (string, error) {
	return os.Getenv("HTTP_HOST"), nil
}

// GetHttpPort returns HTTP port number
// Returns the HTTP port number or error if something goes wrong
func (service *envConfigurationService) GetHttpPort() (int, error) {
	valueStr := os.Getenv("HTTP_PORT")
	if strings.Trim(valueStr, " ") == "" {
		return 0, NewUnknownError("HTTP_PORT is required")
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, NewUnknownErrorWithError("Failed to convert HTTP_PORT to integer", err)
	}

	return value, nil
}

// GetDatabaseConnectionString returns the database connection string
// Returns the database connection string or error if something goes wrong
func (service *envConfigurationService) GetDatabaseConnectionString() (string, error) {
	value := os.Getenv("DATABASE_CONNECTION_STRING")

	if strings.Trim(value, " ") == "" {
		return "", NewUnknownError("DATABASE_CONNECTION_STRING is required")
	}

	return value, nil
}

// GetDatabaseName returns the database name
// Returns the database name or error if something goes wrong
func (service *envConfigurationService) GetDatabaseName() (string, error) {
	value := os.Getenv("EDGE_CLUSTER_DATABASE_NAME")

	if strings.Trim(value, " ") == "" {
		return "", NewUnknownError("EDGE_CLUSTER_DATABASE_NAME is required")
	}

	return value, nil
}

// GetDatabaseCollectionName returns the database collection name
// Returns the database collection name or error if something goes wrong
func (service *envConfigurationService) GetDatabaseCollectionName() (string, error) {
	value := os.Getenv("EDGE_CLUSTER_DATABASE_COLLECTION_NAME")

	if strings.Trim(value, " ") == "" {
		return "", NewUnknownError("EDGE_CLUSTER_DATABASE_COLLECTION_NAME is required")
	}

	return value, nil
}

// GetJwksURL returns the JWKS URL
// Returns the JWKS URL or error if something goes wrong
func (service *envConfigurationService) GetJwksURL() (string, error) {
	value := os.Getenv("JWKS_URL")

	if strings.Trim(value, " ") == "" {
		return "", NewUnknownError("JWKS_URL is required")
	}

	return value, nil
}

// GetK3SDockerImage returns the K3S docker image to be used when creating edge cluster service of type K3S
// Returns the K3S docker image to be used when creating edge cluster service of type K3S or error if something goes wrong
func (service *envConfigurationService) GetK3SDockerImage() (string, error) {
	value := os.Getenv("K3S_DOCKER_IMAGE")

	if strings.Trim(value, " ") == "" {
		return "", NewUnknownError("K3S_DOCKER_IMAGE is required")
	}

	return value, nil
}
