// Package mongodb implements MongoDB repository services
package mongodb

import (
	"context"
	"fmt"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/configuration"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type edgeCluster struct {
	UserEmail     string             `bson:"userEmail" json:"userEmail"`
	ProjectID     string             `bson:"projectID" json:"projectID"`
	Name          string             `bson:"name" json:"name"`
	ClusterSecret string             `bson:"clusterSecret" json:"clusterSecret"`
	ClusterType   models.ClusterType `bson:"clusterType" json:"clusterType"`
}

type mongodbRepositoryService struct {
	connectionString       string
	databaseName           string
	databaseCollectionName string
}

// NewMongodbRepositoryService creates new instance of the mongodbRepositoryService, setting up all dependencies and returns the instance
// Returns the new service or error if something goes wrong
func NewMongodbRepositoryService(
	configurationService configuration.ConfigurationContract) (repository.RepositoryContract, error) {
	if configurationService == nil {
		return nil, commonErrors.NewArgumentNilError("configurationService", "configurationService is required")
	}

	connectionString, err := configurationService.GetDatabaseConnectionString()
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Failed to get connection string to mongodb", err)
	}

	databaseName, err := configurationService.GetDatabaseName()
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Failed to get the database name", err)
	}

	databaseCollectionName, err := configurationService.GetDatabaseCollectionName()
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Failed to get the database collection name", err)
	}

	return &mongodbRepositoryService{
		connectionString:       connectionString,
		databaseName:           databaseName,
		databaseCollectionName: databaseCollectionName,
	}, nil
}

// CreateEdgeCluster creates a new edge cluster.
// context: Optional The reference to the context
// request: Mandatory. The request to create a new edge cluster
// Returns either the result of creating new edge cluster or error if something goes wrong.
func (service *mongodbRepositoryService) CreateEdgeCluster(
	ctx context.Context,
	request *repository.CreateEdgeClusterRequest) (*repository.CreateEdgeClusterResponse, error) {
	client, collection, err := service.createClientAndCollection(ctx)
	if err != nil {
		return nil, err
	}

	defer disconnect(ctx, client)

	insertResult, err := collection.InsertOne(ctx, mapToInternalEdgeCluster(request.UserEmail, request.EdgeCluster))
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Insert edge cluster failed.", err)
	}

	edgeClusterID := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return &repository.CreateEdgeClusterResponse{
		EdgeClusterID: edgeClusterID,
		EdgeCluster:   request.EdgeCluster,
		Cursor:        edgeClusterID,
	}, nil
}

// ReadEdgeCluster read an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to read an existing edge cluster
// Returns either the result of reading an existing edge cluster or error if something goes wrong.
func (service *mongodbRepositoryService) ReadEdgeCluster(
	ctx context.Context,
	request *repository.ReadEdgeClusterRequest) (*repository.ReadEdgeClusterResponse, error) {
	client, collection, err := service.createClientAndCollection(ctx)
	if err != nil {
		return nil, err
	}

	defer disconnect(ctx, client)

	ObjectID, _ := primitive.ObjectIDFromHex(request.EdgeClusterID)
	filter := bson.D{{Key: "_id", Value: ObjectID}, {Key: "userEmail", Value: request.UserEmail}}
	var edgeCluster edgeCluster

	err = collection.FindOne(ctx, filter).Decode(&edgeCluster)
	if err != nil {
		return nil, err
	}

	return &repository.ReadEdgeClusterResponse{
		EdgeCluster: mapFromInternalEdgeCluster(edgeCluster),
	}, nil
}

// UpdateEdgeCluster update an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to update an existing edge cluster
// Returns either the result of updateing an existing edge cluster or error if something goes wrong.
func (service *mongodbRepositoryService) UpdateEdgeCluster(
	ctx context.Context,
	request *repository.UpdateEdgeClusterRequest) (*repository.UpdateEdgeClusterResponse, error) {
	client, collection, err := service.createClientAndCollection(ctx)
	if err != nil {
		return nil, err
	}

	defer disconnect(ctx, client)

	ObjectID, _ := primitive.ObjectIDFromHex(request.EdgeClusterID)
	filter := bson.D{{Key: "_id", Value: ObjectID}, {Key: "userEmail", Value: request.UserEmail}}

	newEdgeCluster := bson.M{
		"$set": bson.M{
			"name":          request.EdgeCluster.Name,
			"projectID":     request.EdgeCluster.ProjectID,
			"clusterSecret": request.EdgeCluster.ClusterSecret,
		}}
	response, err := collection.UpdateOne(ctx, filter, newEdgeCluster)

	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Update edge cluster failed.", err)
	}

	if response.MatchedCount == 0 {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	return &repository.UpdateEdgeClusterResponse{
		EdgeCluster: request.EdgeCluster,
		Cursor:      request.EdgeClusterID,
	}, nil
}

// DeleteEdgeCluster delete an existing edge cluster
// context: Optional The reference to the context
// request: Mandatory. The request to delete an existing edge cluster
// Returns either the result of deleting an existing edge cluster or error if something goes wrong.
func (service *mongodbRepositoryService) DeleteEdgeCluster(
	ctx context.Context,
	request *repository.DeleteEdgeClusterRequest) (*repository.DeleteEdgeClusterResponse, error) {
	client, collection, err := service.createClientAndCollection(ctx)
	if err != nil {
		return nil, err
	}

	defer disconnect(ctx, client)

	ObjectID, _ := primitive.ObjectIDFromHex(request.EdgeClusterID)
	filter := bson.D{{Key: "_id", Value: ObjectID}, {Key: "userEmail", Value: request.UserEmail}}
	response, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Delete edge cluster failed.", err)
	}

	if response.DeletedCount == 0 {
		return nil, repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
	}

	return &repository.DeleteEdgeClusterResponse{}, nil
}

// Search returns the list of edge clusters that matched the criteria
// ctx: Mandatory The reference to the context
// request: Mandatory. The request contains the search criteria
// Returns the list of edge clusters that matched the criteria
func (service *mongodbRepositoryService) Search(
	ctx context.Context,
	request *repository.SearchRequest) (*repository.SearchResponse, error) {
	response := &repository.SearchResponse{
		HasPreviousPage: false,
		HasNextPage:     false,
	}

	ids := []primitive.ObjectID{}
	for _, edgeClusterID := range request.EdgeClusterIDs {
		objectID, err := primitive.ObjectIDFromHex(edgeClusterID)
		if err != nil {
			return nil, repository.NewUnknownErrorWithError(fmt.Sprintf("Failed to decode the edgeClusterID: %s.", edgeClusterID), err)
		}

		ids = append(ids, objectID)
	}

	filter := bson.M{}
	if len(request.EdgeClusterIDs) > 0 {
		filter["_id"] = bson.M{"$in": ids}
	}

	filter["$and"] = []interface{}{
		bson.M{"userEmail": bson.M{"$eq": request.UserEmail}},
	}

	if len(request.ProjectIDs) > 0 {
		if len(filter) > 0 {
			filter["$and"] = []interface{}{
				bson.M{"projectID": bson.M{"$in": request.ProjectIDs}},
			}
		} else {
			filter["projectID"] = bson.M{"$in": request.ProjectIDs}
		}
	}

	client, collection, err := service.createClientAndCollection(ctx)
	if err != nil {
		return nil, err
	}

	defer disconnect(ctx, client)

	response.TotalCount, err = collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Failed to retrieve the number of edge clusters that match the filter criteria", err)
	}

	if response.TotalCount == 0 {
		// No edge cluster matched the filter criteria
		return response, nil
	}

	if request.Pagination.After != nil {
		after := *request.Pagination.After
		objectID, err := primitive.ObjectIDFromHex(after)
		if err != nil {
			return nil, repository.NewUnknownErrorWithError(fmt.Sprintf("Failed to decode the After: %s.", after), err)
		}

		if len(filter) > 0 {
			filter["$and"] = []interface{}{
				bson.M{"_id": bson.M{"$gt": objectID}},
			}
		} else {
			filter["_id"] = bson.M{"$gt": objectID}
		}
	}

	if request.Pagination.Before != nil {
		before := *request.Pagination.Before
		objectID, err := primitive.ObjectIDFromHex(before)
		if err != nil {
			return nil, repository.NewUnknownErrorWithError(fmt.Sprintf("Failed to decode the Before: %s.", before), err)
		}

		if len(filter) > 0 {
			filter["$and"] = []interface{}{
				bson.M{"_id": bson.M{"$lt": objectID}},
			}
		} else {
			filter["_id"] = bson.M{"$lt": objectID}
		}
	}

	findOptions := options.Find()

	if request.Pagination.First != nil {
		findOptions.SetLimit(int64(*request.Pagination.First))
	}

	if request.Pagination.Last != nil {
		findOptions.SetLimit(int64(*request.Pagination.Last))
	}

	if len(request.SortingOptions) > 0 {
		var sortOptionPairs bson.D

		for _, sortingOption := range request.SortingOptions {
			direction := 1
			if sortingOption.Direction == common.Descending {
				direction = -1
			}

			sortOptionPairs = append(
				sortOptionPairs,
				bson.E{
					Key:   sortingOption.Name,
					Value: direction,
				})
		}

		findOptions.SetSort(sortOptionPairs)
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, repository.NewUnknownErrorWithError("Failed to call the Find function on the collection.", err)
	}

	edgeClusters := []models.EdgeClusterWithCursor{}
	for cursor.Next(ctx) {
		var edgeCluster edgeCluster
		var edgeClusterBson bson.M

		err := cursor.Decode(&edgeCluster)
		if err != nil {
			return nil, repository.NewUnknownErrorWithError("Failed to decode the edge cluster", err)
		}

		err = cursor.Decode(&edgeClusterBson)
		if err != nil {
			return nil, repository.NewUnknownErrorWithError("Could not load the data.", err)
		}

		edgeClusterID := edgeClusterBson["_id"].(primitive.ObjectID).Hex()
		edgeClusterWithCursor := models.EdgeClusterWithCursor{
			EdgeClusterID: edgeClusterID,
			EdgeCluster:   mapFromInternalEdgeCluster(edgeCluster),
			Cursor:        edgeClusterID,
		}

		edgeClusters = append(edgeClusters, edgeClusterWithCursor)
	}

	response.EdgeClusters = edgeClusters
	if (request.Pagination.After != nil && request.Pagination.First != nil && int64(*request.Pagination.First) < response.TotalCount) ||
		(request.Pagination.Before != nil && request.Pagination.Last != nil && int64(*request.Pagination.Last) < response.TotalCount) {
		response.HasNextPage = true
		response.HasPreviousPage = true
	} else if request.Pagination.After == nil && request.Pagination.First != nil && int64(*request.Pagination.First) < response.TotalCount {
		response.HasNextPage = true
		response.HasPreviousPage = false
	} else if request.Pagination.Before == nil && request.Pagination.Last != nil && int64(*request.Pagination.Last) < response.TotalCount {
		response.HasNextPage = false
		response.HasPreviousPage = true
	}

	return response, nil
}

func (service *mongodbRepositoryService) createClientAndCollection(ctx context.Context) (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(service.connectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, repository.NewUnknownErrorWithError("Could not connect to mongodb database.", err)
	}

	return client, client.Database(service.databaseName).Collection(service.databaseCollectionName), nil
}

func disconnect(ctx context.Context, client *mongo.Client) {
	_ = client.Disconnect(ctx)
}

func mapToInternalEdgeCluster(email string, from models.EdgeCluster) edgeCluster {
	return edgeCluster{
		UserEmail:     email,
		ProjectID:     from.ProjectID,
		Name:          from.Name,
		ClusterSecret: from.ClusterSecret,
		ClusterType:   from.ClusterType,
	}
}

func mapFromInternalEdgeCluster(from edgeCluster) models.EdgeCluster {
	return models.EdgeCluster{
		ProjectID:     from.ProjectID,
		Name:          from.Name,
		ClusterSecret: from.ClusterSecret,
		ClusterType:   from.ClusterType,
	}
}
