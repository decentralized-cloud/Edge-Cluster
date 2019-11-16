package memory_test

import (
	"context"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	"github.com/decentralized-cloud/edge-cluster/services/repository/memory"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	"github.com/thoas/go-funk"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInMemoryRepositoryService(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	RegisterFailHandler(Fail)
	RunSpecs(t, "In-Memory Repository Service Tests")
}

var _ = Describe("In-Memory Repository Service Tests", func() {
	var (
		sut           repository.RepositoryContract
		ctx           context.Context
		createRequest repository.CreateEdgeClusterRequest
	)

	BeforeEach(func() {
		sut, _ = memory.NewRepositoryService()
		ctx = context.Background()
		createRequest = repository.CreateEdgeClusterRequest{
			EdgeCluster: models.EdgeCluster{
				TenantID:         cuid.New(),
				Name:             cuid.New(),
				K3SClusterSecret: cuid.New(),
			}}
	})

	Context("user tries to instantiate RepositoryService", func() {
		When("all dependecies are resolved and NewRepositoryService is called", func() {
			It("should instantiate the new RepositoryService", func() {
				service, err := memory.NewRepositoryService()
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
				Ω(response.Cursor).Should(Equal(response.EdgeClusterID))
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

		When("user reads the edge cluster", func() {
			It("should return the edge cluster information", func() {
				response, err := sut.ReadEdgeCluster(ctx, &repository.ReadEdgeClusterRequest{EdgeClusterID: edgeClusterID})
				Ω(err).Should(BeNil())
				assertEdgeCluster(response.EdgeCluster, createRequest.EdgeCluster)
			})
		})

		When("user updates the existing edge cluster", func() {
			It("should update the edge cluster information", func() {
				updateRequest := repository.UpdateEdgeClusterRequest{
					EdgeClusterID: edgeClusterID,
					EdgeCluster: models.EdgeCluster{
						TenantID:         cuid.New(),
						Name:             cuid.New(),
						K3SClusterSecret: cuid.New(),
					}}

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

				Ω(repository.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

				var notFoundErr repository.EdgeClusterNotFoundError
				_ = errors.As(err, &notFoundErr)

				Ω(notFoundErr.EdgeClusterID).Should(Equal(edgeClusterID))
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

				Ω(repository.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

				var notFoundErr repository.EdgeClusterNotFoundError
				_ = errors.As(err, &notFoundErr)

				Ω(notFoundErr.EdgeClusterID).Should(Equal(edgeClusterID))
			})
		})

		When("user tries to update the edge cluster", func() {
			It("should return NotFoundError", func() {
				updateRequest := repository.UpdateEdgeClusterRequest{
					EdgeClusterID: edgeClusterID,
					EdgeCluster: models.EdgeCluster{
						TenantID:         cuid.New(),
						Name:             cuid.New(),
						K3SClusterSecret: cuid.New(),
					}}
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

	Context("edge clusters exist", func() {
		var (
			edgeClusterIDs []string
			tenantIDs      []string
			edgeClusters   []models.EdgeCluster
		)

		BeforeEach(func() {
			edgeClusters = []models.EdgeCluster{}
			tenantIDs = []string{}
			for idx := 0; idx < rand.Intn(20)+10; idx++ {
				tenantID := cuid.New()
				tenantIDs = append(tenantIDs, tenantID)

				edgeClusters = append(
					edgeClusters,
					models.EdgeCluster{
						TenantID:         tenantID,
						Name:             cuid.New(),
						K3SClusterSecret: cuid.New(),
					})
			}

			edgeClusterIDs = funk.Map(edgeClusters, func(edgeCluster models.EdgeCluster) string {
				response, _ := sut.CreateEdgeCluster(ctx, &repository.CreateEdgeClusterRequest{
					EdgeCluster: models.EdgeCluster{
						TenantID:         edgeCluster.TenantID,
						Name:             edgeCluster.Name,
						K3SClusterSecret: edgeCluster.K3SClusterSecret,
					},
				})

				return response.EdgeClusterID
			}).([]string)
		})

		When("user search for edge clusters with/without sorting options", func() {
			It("should sort the result ascending when no sorting direction is provided", func() {
				response, _ := sut.Search(ctx, &repository.SearchRequest{})
				convertedEdgeClusters := funk.Map(response.EdgeClusters, func(edgeClusterWithCursor models.EdgeClusterWithCursor) models.EdgeCluster {
					return edgeClusterWithCursor.EdgeCluster
				}).([]models.EdgeCluster)

				for idx := range convertedEdgeClusters[:len(convertedEdgeClusters)-1] {
					Ω(convertedEdgeClusters[idx].Name < convertedEdgeClusters[idx+1].Name).Should(BeTrue())
				}
			})

			It("should sort the result ascending when sorting direction is ascending", func() {
				response, _ := sut.Search(ctx, &repository.SearchRequest{
					SortingOptions: []common.SortingOptionPair{
						common.SortingOptionPair{
							Name:      "name",
							Direction: common.Ascending,
						}}})
				convertedEdgeClusters := funk.Map(response.EdgeClusters, func(edgeClusterWithCursor models.EdgeClusterWithCursor) models.EdgeCluster {
					return edgeClusterWithCursor.EdgeCluster
				}).([]models.EdgeCluster)

				for idx := range convertedEdgeClusters[:len(convertedEdgeClusters)-1] {
					Ω(convertedEdgeClusters[idx].Name < convertedEdgeClusters[idx+1].Name).Should(BeTrue())
				}
			})

			It("should sort the result descending when sorting direction is descending", func() {
				response, _ := sut.Search(ctx, &repository.SearchRequest{
					SortingOptions: []common.SortingOptionPair{
						common.SortingOptionPair{
							Name:      "name",
							Direction: common.Descending,
						}}})
				convertedEdgeClusters := funk.Map(response.EdgeClusters, func(edgeClusterWithCursor models.EdgeClusterWithCursor) models.EdgeCluster {
					return edgeClusterWithCursor.EdgeCluster
				}).([]models.EdgeCluster)

				for idx := range convertedEdgeClusters[:len(convertedEdgeClusters)-1] {
					Ω(convertedEdgeClusters[idx].Name > convertedEdgeClusters[idx+1].Name).Should(BeTrue())
				}
			})
		})

		When("user search for edge clusters without any edge cluster ID or tenant ID provided", func() {
			It("should return all edge clusters", func() {
				response, err := sut.Search(ctx, &repository.SearchRequest{})
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).Should(HaveLen(len(edgeClusterIDs)))
				Ω(response.TotalCount).Should(Equal(int64(len(edgeClusterIDs))))

				filteredEdgeClusters := funk.Filter(response.EdgeClusters, func(edgeClusterWithCursor models.EdgeClusterWithCursor) bool {
					return !funk.Contains(edgeClusterIDs, edgeClusterWithCursor.EdgeClusterID)
				}).([]models.EdgeClusterWithCursor)

				Ω(filteredEdgeClusters).Should(HaveLen(0))
			})
		})

		When("user search for edge clusters with edge cluster IDs provided", func() {
			var (
				numberOfEdgeClusterIDs  int
				shuffeledEdgeClusterIDs []string
			)

			BeforeEach(func() {
				shuffeledEdgeClusterIDs = funk.ShuffleString(edgeClusterIDs)
				numberOfEdgeClusterIDs = rand.Intn(10)
			})

			It("should return filtered edge cluster list", func() {
				response, err := sut.Search(ctx, &repository.SearchRequest{
					EdgeClusterIDs: shuffeledEdgeClusterIDs[:numberOfEdgeClusterIDs],
				})
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).Should(HaveLen(numberOfEdgeClusterIDs))
				Ω(response.TotalCount).Should(Equal(int64(numberOfEdgeClusterIDs)))

				filteredEdgeClusters := funk.Filter(response.EdgeClusters, func(edgeClusterWithCursor models.EdgeClusterWithCursor) bool {
					return !funk.Contains(edgeClusterIDs, edgeClusterWithCursor.EdgeClusterID)
				}).([]models.EdgeClusterWithCursor)

				Ω(filteredEdgeClusters).Should(HaveLen(0))
			})
		})

		When("user search for edge clusters with tenant IDs provided", func() {
			var (
				numberOfTenantIDs  int
				shuffeledTenantIDs []string
			)

			BeforeEach(func() {
				shuffeledTenantIDs = funk.ShuffleString(tenantIDs)
				numberOfTenantIDs = rand.Intn(10)
			})

			It("should return filtered edge cluster list", func() {
				response, err := sut.Search(ctx, &repository.SearchRequest{
					TenantIDs: shuffeledTenantIDs[:numberOfTenantIDs],
				})
				Ω(err).Should(BeNil())
				Ω(response.EdgeClusters).Should(HaveLen(numberOfTenantIDs))
				Ω(response.TotalCount).Should(Equal(int64(numberOfTenantIDs)))

				filteredEdgeClusters := funk.Filter(response.EdgeClusters, func(edgeClusterWithCursor models.EdgeClusterWithCursor) bool {
					return !funk.Contains(tenantIDs, edgeClusterWithCursor.EdgeCluster.TenantID)
				}).([]models.EdgeClusterWithCursor)

				Ω(filteredEdgeClusters).Should(HaveLen(0))
			})
		})
	})
})

func assertEdgeCluster(edgeCluster, expectedEdgeCluster models.EdgeCluster) {
	Ω(edgeCluster).ShouldNot(BeNil())
	Ω(edgeCluster.TenantID).Should(Equal(expectedEdgeCluster.TenantID))
	Ω(edgeCluster.Name).Should(Equal(expectedEdgeCluster.Name))
	Ω(edgeCluster.K3SClusterSecret).Should(Equal(expectedEdgeCluster.K3SClusterSecret))
}
