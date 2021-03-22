// Package configuration implements configuration service required by the edge-cluster service
package configuration

// ConfigurationContract declares the service that provides configuration required by different Tenat modules
type ConfigurationContract interface {
	// GetGrpcHost returns gRPC host name
	// Returns the gRPC host name or error if something goes wrong
	GetGrpcHost() (string, error)

	// GetGrpcPort returns gRPC port number
	// Returns the gRPC port number or error if something goes wrong
	GetGrpcPort() (int, error)

	// GetHttpHost returns HTTP host name
	// Returns the HTTP host name or error if something goes wrong
	GetHttpHost() (string, error)

	// GetHttpPort returns HTTP port number
	// Returns the HTTP port number or error if something goes wrong
	GetHttpPort() (int, error)

	// GetDatabaseConnectionString returns the database connection string
	// Returns the database connection string or error if something goes wrong
	GetDatabaseConnectionString() (string, error)

	// GetDatabaseName returns the database name
	// Returns the database name or error if something goes wrong
	GetDatabaseName() (string, error)

	// GetDatabaseCollectionName returns the database collection name
	// Returns the database collection name or error if something goes wrong
	GetDatabaseCollectionName() (string, error)

	// GetJwksURL returns the JWKS URL
	// Returns the JWKS URL or error if something goes wrong
	GetJwksURL() (string, error)

	// GetK3SDockerImage returns the K3S docker image to be used when creating edge cluster service of type K3S
	// Returns the K3S docker image to be used when creating edge cluster service of type K3S or error if something goes wrong
	GetK3SDockerImage() (string, error)
}
