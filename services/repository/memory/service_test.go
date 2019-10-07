package memory_test

import (
	"context"
	"errors"
	"testing"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	"github.com/decentralized-cloud/edge-cluster/services/repository/memory"
	"github.com/lucsky/cuid"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInMemoryRepositoryService(t *testing.T) {
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
			TenantID: cuid.New(),
			EdgeCluster: models.EdgeCluster{
				Name: cuid.New(),
			}}
	})

	Context("user tries to instantiate EdgeClusterRepositoryService", func() {
		When("all dependecies are resolved and NewRepositoryService is called", func() {
			It("should instantiate the new EdgeClusterRepositoryService", func() {
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
				response, err := sut.ReadEdgeCluster(
					ctx,
					&repository.ReadEdgeClusterRequest{
						TenantID:      createRequest.TenantID,
						EdgeClusterID: edgeClusterID,
					})
				Ω(err).Should(BeNil())
				Ω(response.EdgeCluster).ShouldNot(BeNil())
				Ω(response.EdgeCluster.Name).Should(Equal(createRequest.EdgeCluster.Name))
			})
		})

		When("user updates the existing edge cluster", func() {
			It("should update the edge cluster information", func() {
				updateRequest := repository.UpdateEdgeClusterRequest{
					TenantID:      createRequest.TenantID,
					EdgeClusterID: edgeClusterID,
					EdgeCluster: models.EdgeCluster{
						Name: cuid.New(),
					}}

				_, err := sut.UpdateEdgeCluster(ctx, &updateRequest)
				Ω(err).Should(BeNil())

				response, err := sut.ReadEdgeCluster(
					ctx,
					&repository.ReadEdgeClusterRequest{
						TenantID:      createRequest.TenantID,
						EdgeClusterID: edgeClusterID,
					})
				Ω(err).Should(BeNil())
				Ω(response.EdgeCluster).ShouldNot(BeNil())
				Ω(response.EdgeCluster.Name).Should(Equal(updateRequest.EdgeCluster.Name))
			})
		})

		When("user deletes the edge cluster", func() {
			It("should delete the edge cluster", func() {
				_, err := sut.DeleteEdgeCluster(
					ctx,
					&repository.DeleteEdgeClusterRequest{
						TenantID:      createRequest.TenantID,
						EdgeClusterID: edgeClusterID,
					})
				Ω(err).Should(BeNil())

				response, err := sut.ReadEdgeCluster(
					ctx,
					&repository.ReadEdgeClusterRequest{
						TenantID:      createRequest.TenantID,
						EdgeClusterID: edgeClusterID,
					})
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
			_, _ = sut.CreateEdgeCluster(ctx, &createRequest)
			_, _ = sut.DeleteEdgeCluster(
				ctx,
				&repository.DeleteEdgeClusterRequest{
					TenantID:      createRequest.TenantID,
					EdgeClusterID: edgeClusterID,
				})
		})

		When("user reads the edge cluster", func() {
			It("should return NotFoundError", func() {
				response, err := sut.ReadEdgeCluster(
					ctx,
					&repository.ReadEdgeClusterRequest{
						TenantID:      createRequest.TenantID,
						EdgeClusterID: edgeClusterID,
					})
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
					TenantID:      createRequest.TenantID,
					EdgeClusterID: edgeClusterID,
					EdgeCluster: models.EdgeCluster{
						Name: cuid.New(),
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
				response, err := sut.DeleteEdgeCluster(
					ctx,
					&repository.DeleteEdgeClusterRequest{
						TenantID:      createRequest.TenantID,
						EdgeClusterID: edgeClusterID,
					})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())

				Ω(repository.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

				var notFoundErr repository.EdgeClusterNotFoundError
				_ = errors.As(err, &notFoundErr)

				Ω(notFoundErr.EdgeClusterID).Should(Equal(edgeClusterID))
			})
		})
	})

	Context("tenant does not exist", func() {

		var (
			tenantID string
		)

		BeforeEach(func() {
			tenantID = cuid.New()
		})

		When("user tries to read the edge cluster", func() {
			It("should return TenantNotFoundError", func() {
				response, err := sut.ReadEdgeCluster(
					ctx,
					&repository.ReadEdgeClusterRequest{
						TenantID:      tenantID,
						EdgeClusterID: cuid.New(),
					})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())

				Ω(repository.IsTenantNotFoundError(err)).Should(BeTrue())

				var tenantNotFoundErr repository.TenantNotFoundError
				_ = errors.As(err, &tenantNotFoundErr)

				Ω(tenantNotFoundErr.TenantID).Should(Equal(tenantID))
			})
		})

		When("user tries to update an existing edge cluster", func() {
			It("should return TenantNotFoundError", func() {
				response, err := sut.UpdateEdgeCluster(
					ctx,
					&repository.UpdateEdgeClusterRequest{
						TenantID:      tenantID,
						EdgeClusterID: cuid.New(),
					})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())

				Ω(repository.IsTenantNotFoundError(err)).Should(BeTrue())

				var tenantNotFoundErr repository.TenantNotFoundError
				_ = errors.As(err, &tenantNotFoundErr)

				Ω(tenantNotFoundErr.TenantID).Should(Equal(tenantID))
			})
		})

		When("user tries to delete an existing edge cluster", func() {
			It("should return TenantNotFoundError", func() {
				response, err := sut.DeleteEdgeCluster(
					ctx,
					&repository.DeleteEdgeClusterRequest{
						TenantID:      tenantID,
						EdgeClusterID: cuid.New(),
					})
				Ω(err).Should(HaveOccurred())
				Ω(response).Should(BeNil())

				Ω(repository.IsTenantNotFoundError(err)).Should(BeTrue())

				var tenantNotFoundErr repository.TenantNotFoundError
				_ = errors.As(err, &tenantNotFoundErr)

				Ω(tenantNotFoundErr.TenantID).Should(Equal(tenantID))
			})
		})
	})
})
