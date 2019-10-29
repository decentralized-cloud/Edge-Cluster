package business_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	"github.com/decentralized-cloud/edge-cluster/services/repository"
	repsoitoryMock "github.com/decentralized-cloud/edge-cluster/services/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBusinessService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Business Service Tests")
}

var _ = Describe("Business Service Tests", func() {
	var (
		mockCtrl              *gomock.Controller
		sut                   business.BusinessContract
		mockRepositoryService *repsoitoryMock.MockRepositoryContract
		ctx                   context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockRepositoryService = repsoitoryMock.NewMockRepositoryContract(mockCtrl)
		sut, _ = business.NewBusinessService(mockRepositoryService)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate EdgeClusterService", func() {
		When("edge cluster repository service is not provided and NewEdgeClusterService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := business.NewBusinessService(nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("repositoryService", "", err)
			})
		})

		When("all dependencies are resolved and NewEdgeClusterService is called", func() {
			It("should instantiate the new EdgeClusterService", func() {
				service, err := business.NewBusinessService(mockRepositoryService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Describe("CreateEdgeCluster", func() {
		var (
			request business.CreateEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.CreateEdgeClusterRequest{
				TenantID: cuid.New(),
				EdgeCluster: models.EdgeCluster{
					Name: cuid.New(),
				}}
		})

		Context("edge cluster service is instantiated", func() {
			When("CreateEdgeCluster is called with correct input parameters", func() {
				It("should call edge cluster repository CreateEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						CreateEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.CreateEdgeClusterRequest) {
							Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))
						}).
						Return(&repository.CreateEdgeClusterResponse{EdgeClusterID: cuid.New()}, nil)

					response, err := sut.CreateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})

				When("and edge cluster repository CreateEdgeCluster return EdgeClusterAlreadyExistError", func() {
					It("should return EdgeClusterAlreadyExistsError", func() {
						expectedError := repository.NewEdgeClusterAlreadyExistsError()
						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedError)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						assertEdgeClusterAlreadyExistsError(response.Err, expectedError)
					})
				})

				When("and edge cluster repository CreateEdgeCluster return any other error", func() {
					It("should return UnknownError", func() {
						expectedError := errors.New(cuid.New())
						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedError)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						assertUnknowError(expectedError.Error(), response.Err, expectedError)
					})
				})

				When("and edge cluster repository CreateEdgeCluster return no error", func() {
					It("should return the new edgeClusterID", func() {
						edgeClusterID := cuid.New()
						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&repository.CreateEdgeClusterResponse{EdgeClusterID: edgeClusterID}, nil)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						Ω(response.Err).Should(BeNil())
						Ω(response.EdgeClusterID).ShouldNot(BeNil())
						Ω(response.EdgeClusterID).Should(Equal(edgeClusterID))
					})
				})
			})
		})
	})

	Describe("ReadEdgeCluster", func() {
		var (
			request business.ReadEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.ReadEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("ReadEdgeCluster is called with correct input parameters", func() {
				It("should call edge cluster repository ReadEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.ReadEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
						}).
						Return(&repository.ReadEdgeClusterResponse{EdgeCluster: models.EdgeCluster{Name: cuid.New()}}, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("and edge cluster repository ReadEdgeCluster cannot find the edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertEdgeClusterNotFoundError(request.EdgeClusterID, response.Err, expectedError)
				})
			})

			When("and edge cluster repository ReadEdgeCluster return any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("and edge cluster repository ReadEdgeCluster return no error", func() {
				It("should return the edgeClusterID", func() {
					tenantID := cuid.New()
					edgeCluster := models.EdgeCluster{Name: cuid.New()}
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repository.ReadEdgeClusterResponse{
							TenantID:    tenantID,
							EdgeCluster: edgeCluster,
						}, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					Ω(response.EdgeCluster).ShouldNot(BeNil())
					Ω(response.TenantID).Should(Equal(tenantID))
					Ω(response.EdgeCluster.Name).Should(Equal(edgeCluster.Name))
				})
			})
		})
	})

	Describe("UpdateEdgeCluster", func() {
		var (
			request business.UpdateEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.UpdateEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
				EdgeCluster:   models.EdgeCluster{Name: cuid.New()},
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("UpdateEdgeCluster is called with correct input parameters", func() {
				It("should call edge cluster repository UpdateEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.UpdateEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
							Ω(mappedRequest.EdgeCluster.Name).Should(Equal(request.EdgeCluster.Name))
						}).
						Return(&repository.UpdateEdgeClusterResponse{}, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("and edge cluster repository UpdateEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertEdgeClusterNotFoundError(request.EdgeClusterID, response.Err, expectedError)
				})
			})

			When("and edge cluster repository UpdateEdgeCluster return any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("and edge cluster repository UpdateEdgeCluster return no error", func() {
				It("should return no error", func() {
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repository.UpdateEdgeClusterResponse{}, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})
		})
	})

	Describe("DeleteEdgeCluster is called.", func() {
		var (
			request business.DeleteEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.DeleteEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("input parameters are valid", func() {
				It("should call edge cluster repository DeleteEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repository.DeleteEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
						}).
						Return(&repository.DeleteEdgeClusterResponse{}, nil)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository DeleteEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repository.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertEdgeClusterNotFoundError(request.EdgeClusterID, response.Err, expectedError)
				})
			})

			When("edge cluster repository DeleteEdgeCluster is faced with any other error", func() {
				It("should return UnknownError", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertUnknowError(expectedError.Error(), response.Err, expectedError)
				})
			})

			When("edge cluster repository DeleteEdgeCluster completes successfully", func() {
				It("should return no error", func() {
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repository.DeleteEdgeClusterResponse{}, nil)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})
		})
	})
})

func assertArgumentNilError(expectedArgumentName, expectedMessage string, err error) {
	Ω(commonErrors.IsArgumentNilError(err)).Should(BeTrue())

	var argumentNilErr commonErrors.ArgumentNilError
	_ = errors.As(err, &argumentNilErr)

	if expectedArgumentName != "" {
		Ω(argumentNilErr.ArgumentName).Should(Equal(expectedArgumentName))
	}

	if expectedMessage != "" {
		Ω(strings.Contains(argumentNilErr.Error(), expectedMessage)).Should(BeTrue())
	}
}

func assertUnknowError(expectedMessage string, err error, nestedErr error) {
	Ω(business.IsUnknownError(err)).Should(BeTrue())

	var unknownErr business.UnknownError
	_ = errors.As(err, &unknownErr)

	Ω(strings.Contains(unknownErr.Error(), expectedMessage)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeClusterAlreadyExistsError(err error, nestedErr error) {
	Ω(business.IsEdgeClusterAlreadyExistsError(err)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeClusterNotFoundError(expectedEdgeClusterID string, err error, nestedErr error) {
	Ω(business.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

	var edgeClusterNotFoundErr business.EdgeClusterNotFoundError
	_ = errors.As(err, &edgeClusterNotFoundErr)

	Ω(edgeClusterNotFoundErr.EdgeClusterID).Should(Equal(expectedEdgeClusterID))
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}
