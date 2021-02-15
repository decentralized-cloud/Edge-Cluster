package mongodb_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/decentralized-cloud/edge-cluster/models"
	configurationMock "github.com/decentralized-cloud/edge-cluster/services/configuration/mock"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	"github.com/decentralized-cloud/edge-cluster/services/repository/mongodb"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMongodbRepositoryService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mongodb Repository Service Tests")
}

var _ = Describe("Mongodb Repository Service Tests", func() {
	var (
		mockCtrl      *gomock.Controller
		sut           repository.RepositoryContract
		ctx           context.Context
		createRequest repository.CreateEdgeClusterRequest
	)

	BeforeEach(func() {
		connectionString := os.Getenv("DATABASE_CONNECTION_STRING")
		if strings.Trim(connectionString, " ") == "" {
			connectionString = "mongodb://mongodb:27017"
		}

		mockCtrl = gomock.NewController(GinkgoT())
		mockConfigurationService := configurationMock.NewMockConfigurationContract(mockCtrl)
		mockConfigurationService.
			EXPECT().
			GetDatabaseConnectionString().
			Return(connectionString, nil)

		mockConfigurationService.
			EXPECT().
			GetDatabaseName().
			Return("edge-clusters", nil)

		mockConfigurationService.
			EXPECT().
			GetDatabaseCollectionName().
			Return("edge-clusters", nil)

		sut, _ = mongodb.NewMongodbRepositoryService(mockConfigurationService)
		ctx = context.Background()
		createRequest = repository.CreateEdgeClusterRequest{
			EdgeCluster: models.EdgeCluster{
				ProjectID:     cuid.New(),
				Name:          cuid.New(),
				ClusterSecret: cuid.New(),
				ClusterType:   models.K3S,
			},
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate RepositoryService", func() {
		When("all dependencies are resolved and NewRepositoryService is called", func() {
			It("should instantiate the new RepositoryService", func() {
				mockConfigurationService := configurationMock.NewMockConfigurationContract(mockCtrl)
				mockConfigurationService.
					EXPECT().
					GetDatabaseConnectionString().
					Return(cuid.New(), nil)

				mockConfigurationService.
					EXPECT().
					GetDatabaseName().
					Return(cuid.New(), nil)

				mockConfigurationService.
					EXPECT().
					GetDatabaseCollectionName().
					Return(cuid.New(), nil)

				service, err := mongodb.NewMongodbRepositoryService(mockConfigurationService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Context("user going to create a new edge cluster", func() {
		When("create edge cluster is called", func() {
			It("should create the new edge cluster", func() {
				response, err := sut.CreateEdgeCluster(ctx, &createRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusterID).ShouldNot(BeNil())
				Ω(response.EdgeClusterID).Should(Equal(response.EdgeClusterID))
				Ω(response.Cursor).ShouldNot(BeNil())
				Ω(response.Cursor).Should(Equal(response.Cursor))
				assertEdgeCluster(response.EdgeCluster, createRequest.EdgeCluster)
			})
		})
	})

	Context("edge cluster already exists", func() {
		var (
			edgeClusterID string
		)

		BeforeEach(func() {
			response, _ := sut.CreateEdgeCluster(ctx, &createRequest)
			edgeClusterID = response.EdgeClusterID
		})

		When("user reads a edge cluster by Id", func() {
			It("should return a edge cluster", func() {
				response, err := sut.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(BeNil())
				assertEdgeCluster(response.EdgeCluster, createRequest.EdgeCluster)
			})
		})

		When("user updates the existing edge cluster", func() {
			It("should update the edge cluster's information", func() {
				updateRequest := repository.UpdateEdgeClusterRequest{
					EdgeClusterID: edgeClusterID,
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					},
				}

				updateResponse, err := sut.UpdateEdgeCluster(ctx, &updateRequest)
				Ω(err).Should(BeNil())
				Ω(updateResponse.Cursor).Should(Equal(edgeClusterID))
				assertEdgeCluster(updateResponse.EdgeCluster, updateRequest.EdgeCluster)

				readResponse, err := sut.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(BeNil())
				assertEdgeCluster(readResponse.EdgeCluster, updateRequest.EdgeCluster)
			})
		})

		When("user deletes the edge cluster", func() {
			It("should delete the edge cluster", func() {
				_, err := sut.DeleteEdgeCluster(ctx, &repository.DeleteEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(BeNil())

				response, err := sut.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())
			})
		})
	})

	Context("edge cluster does not exist", func() {
		var (
			edgeClusterID string
		)

		BeforeEach(func() {
			edgeClusterID = cuid.New()
		})

		When("user reads the edge cluster", func() {
			It("should return NotFoundError", func() {
				response, err := sut.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())
			})
		})

		When("user tries to update the edge cluster", func() {
			It("should return NotFoundError", func() {
				updateRequest := repository.UpdateEdgeClusterRequest{
					EdgeClusterID: edgeClusterID,
					EdgeCluster: models.EdgeCluster{
						ProjectID:     cuid.New(),
						Name:          cuid.New(),
						ClusterSecret: cuid.New(),
						ClusterType:   models.K3S,
					},
				}

				response, err := sut.UpdateEdgeCluster(ctx, &updateRequest)
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())

				Ω(repository.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

				var notFoundErr repository.EdgeClusterNotFoundError
				_ = errors.As(err, &notFoundErr)

				Ω(notFoundErr.EdgeClusterID).Should(Equal(edgeClusterID))
			})
		})

		When("user tries to delete the edge cluster", func() {
			It("should return NotFoundError", func() {
				response, err := sut.DeleteEdgeCluster(ctx, &repository.DeleteEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())

				Ω(repository.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

				var notFoundErr repository.EdgeClusterNotFoundError
				_ = errors.As(err, &notFoundErr)

				Ω(notFoundErr.EdgeClusterID).Should(Equal(edgeClusterID))
			})
		})
	})

	Context("edge clusters already exists", func() {
		var (
			edgeClusterIDs []string
			projectIDs     []string
		)

		BeforeEach(func() {
			projectIDs = []string{}
			edgeClusterIDs = []string{}

			for i := 0; i < 10; i++ {
				edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
				edgeClusterSecret := fmt.Sprintf("%s%d", "ClusterSecret", i)
				createRequest.EdgeCluster.Name = edgeClusterName
				createRequest.EdgeCluster.ClusterSecret = edgeClusterSecret
				createRequest.EdgeCluster.ClusterType = models.K3S
				response, _ := sut.CreateEdgeCluster(ctx, &createRequest)
				edgeClusterIDs = append(edgeClusterIDs, response.EdgeClusterID)
				projectIDs = append(projectIDs, response.EdgeCluster.ProjectID)
			}
		})

		When("user searches for edge clusters with selected edge cluster Ids and first 10 projects without providing project Id", func() {
			It("should return first 10 edge clusters", func() {
				first := 10
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     []string{},
					Pagination: common.Pagination{
						After: nil,
						First: &first,
					},
					SortingOptions: []common.SortingOptionPair{},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(10))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 10; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge cluster Ids and first 10 projects", func() {
			It("should return first 10 edge clusters", func() {
				first := 10
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						After: nil,
						First: &first,
					},
					SortingOptions: []common.SortingOptionPair{},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(10))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 10; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge cluster Ids and first 5 projects", func() {
			It("should return first 5 edge clusters", func() {
				first := 5
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						After: nil,
						First: &first,
					},
					SortingOptions: []common.SortingOptionPair{},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(5))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 5; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge clusters Ids with After parameter provided.", func() {
			It("should return first 9 edge clusters after provided edge clusters id", func() {
				first := 9
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						After: &edgeClusterIDs[0],
						First: &first,
					},
					SortingOptions: []common.SortingOptionPair{},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(9))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 1; i < 10; i++ {
					Ω(response.EdgeClusters[i-1].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i-1].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i-1].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge clusters Ids and last 10 edge clusters.", func() {
			It("should return last 10 edge clusters.", func() {
				last := 10
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						Before: nil,
						Last:   &last,
					},
					SortingOptions: []common.SortingOptionPair{},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(10))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 10; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge clusters Ids with Before parameter provided.", func() {
			It("should return first 9 edge clusters before provided edge cluster id", func() {
				last := 9
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						Before: &edgeClusterIDs[9],
						Last:   &last,
					},
					SortingOptions: []common.SortingOptionPair{},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(9))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 9; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge cluster Ids and first 10 edge clusters with ascending order on name property", func() {
			It("should return first 10 edge clusters in adcending order on name field", func() {
				first := 10
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						After: nil,
						First: &first,
					},
					SortingOptions: []common.SortingOptionPair{
						{Name: "name", Direction: common.Ascending},
					},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(10))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 10; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[i]))
				}
			})
		})

		When("user searches for edge clusters with selected edge cluster Ids and first 10 edge clusters with descending order on name property", func() {
			It("should return first 10 edge clusters in descending order on name field", func() {
				first := 10
				searchRequest := repository.SearchRequest{
					EdgeClusterIDs: edgeClusterIDs,
					ProjectIDs:     projectIDs,
					Pagination: common.Pagination{
						After: nil,
						First: &first,
					},
					SortingOptions: []common.SortingOptionPair{
						{Name: "name", Direction: common.Descending},
					},
				}

				response, err := sut.Search(ctx, &searchRequest)
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).ShouldNot(BeNil())
				Ω(len(response.EdgeClusters)).Should(Equal(10))
				Ω(response.TotalCount).Should(Equal(int64(10)))
				for i := 0; i < 10; i++ {
					Ω(response.EdgeClusters[i].EdgeClusterID).Should(Equal(edgeClusterIDs[9-i]))
					edgeClusterName := fmt.Sprintf("%s%d", "Name", 9-i)
					Ω(response.EdgeClusters[i].EdgeCluster.Name).Should(Equal(edgeClusterName))
					Ω(response.EdgeClusters[i].EdgeCluster.ProjectID).Should(Equal(projectIDs[9-i]))
				}
			})
		})

	})

})

func assertEdgeCluster(edgeCluster, expectedEdgeCluster models.EdgeCluster) {
	Ω(edgeCluster).ShouldNot(BeNil())
	Ω(edgeCluster.Name).Should(Equal(expectedEdgeCluster.Name))
	Ω(edgeCluster.ProjectID).Should(Equal(expectedEdgeCluster.ProjectID))
	Ω(edgeCluster.ClusterSecret).Should(Equal(expectedEdgeCluster.ClusterSecret))
	Ω(edgeCluster.ClusterType).Should(Equal(expectedEdgeCluster.ClusterType))
}
