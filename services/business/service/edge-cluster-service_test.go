package service_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business/contract"
	"github.com/decentralized-cloud/edge-cluster/services/business/service"
	repositoryContract "github.com/decentralized-cloud/edge-cluster/services/repository/contract"
	repsoitoryMock "github.com/decentralized-cloud/edge-cluster/services/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EdgeClusterService Tests", func() {
	var (
		mockCtrl                         *gomock.Controller
		sut                              contract.EdgeClusterServiceContract
		mockEdgeClusterRepositoryService *repsoitoryMock.MockEdgeClusterRepositoryServiceContract
		ctx                              context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockEdgeClusterRepositoryService = repsoitoryMock.NewMockEdgeClusterRepositoryServiceContract(mockCtrl)
		sut, _ = service.NewEdgeClusterService(mockEdgeClusterRepositoryService)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate EdgeClusterService", func() {
		When("edge cluster repository service is not provided and NewEdgeClusterService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := service.NewEdgeClusterService(nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("repositoryService", "", err)
			})
		})

		When("all dependencies are resolved and NewEdgeClusterService is called", func() {
			It("should instantiate the new EdgeClusterService", func() {
				service, err := service.NewEdgeClusterService(mockEdgeClusterRepositoryService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Describe("CreateEdgeCluster", func() {
		var (
			request contract.CreateEdgeClusterRequest
		)

		BeforeEach(func() {
			request = contract.CreateEdgeClusterRequest{
				TenantID: cuid.New(),
				EdgeCluster: models.EdgeCluster{
					Name: cuid.New(),
				}}
		})

		Context("edge cluster service is instantiated", func() {
			When("CreateEdgeCluster is called without context", func() {
				It("should return ArgumentNilError", func() {
					response, err := sut.CreateEdgeCluster(nil, &request)
					Ω(err).Should(BeNil())
					assertArgumentNilError("ctx", "", response.Err)
				})
			})

			When("CreateEdgeCluster is called without request", func() {
				It("should return ArgumentNilError", func() {
					response, err := sut.CreateEdgeCluster(ctx, nil)
					Ω(err).Should(BeNil())
					assertArgumentNilError("request", "", response.Err)
				})
			})

			When("CreateEdgeCluster is called with invalid request", func() {
				It("should return ArgumentNilError", func() {
					invalidRequest := contract.CreateEdgeClusterRequest{
						EdgeCluster: models.EdgeCluster{
							Name: "",
						}}

					response, err := sut.CreateEdgeCluster(ctx, &invalidRequest)
					Ω(err).Should(BeNil())
					validationErr := invalidRequest.Validate()
					assertArgumentError("request", validationErr.Error(), response.Err, validationErr)
				})
			})

			When("CreateEdgeCluster is called with correct input parameters", func() {
				It("should call edge cluster repository CreateEdgeCluster method", func() {
					mockEdgeClusterRepositoryService.
						EXPECT().
						CreateEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.CreateEdgeClusterRequest) {
							Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))
						}).
						Return(&repositoryContract.CreateEdgeClusterResponse{EdgeClusterID: cuid.New()}, nil)

					response, err := sut.CreateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})

				When("and edge cluster repository CreateEdgeCluster return EdgeClusterAlreadyExistError", func() {
					It("should return EdgeClusterAlreadyExistsError", func() {
						expectedError := repositoryContract.NewEdgeClusterAlreadyExistsError()
						mockEdgeClusterRepositoryService.
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
						mockEdgeClusterRepositoryService.
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
						mockEdgeClusterRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&repositoryContract.CreateEdgeClusterResponse{EdgeClusterID: edgeClusterID}, nil)

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
			request contract.ReadEdgeClusterRequest
		)

		BeforeEach(func() {
			request = contract.ReadEdgeClusterRequest{
				TenantID:      cuid.New(),
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("ReadEdgeCluster is called without context", func() {
				It("should return ArgumentNilError", func() {
					response, err := sut.ReadEdgeCluster(nil, &request)
					Ω(err).Should(BeNil())
					assertArgumentNilError("ctx", "", response.Err)
				})
			})

			When("ReadEdgeCluster is called without request", func() {
				It("should return ArgumentNilError", func() {
					response, err := sut.ReadEdgeCluster(ctx, nil)
					Ω(err).Should(BeNil())
					assertArgumentNilError("request", "", response.Err)
				})
			})

			When("ReadEdgeCluster is called with invalid request", func() {
				It("should return ArgumentNilError", func() {
					invalidRequest := contract.ReadEdgeClusterRequest{
						EdgeClusterID: "",
					}

					response, err := sut.ReadEdgeCluster(ctx, &invalidRequest)
					Ω(err).Should(BeNil())
					validationErr := invalidRequest.Validate()
					assertArgumentError("request", validationErr.Error(), response.Err, validationErr)
				})
			})

			When("ReadEdgeCluster is called with correct input parameters", func() {
				It("should call edge cluster repository ReadEdgeCluster method", func() {
					mockEdgeClusterRepositoryService.
						EXPECT().
						ReadEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.ReadEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
						}).
						Return(&repositoryContract.ReadEdgeClusterResponse{EdgeCluster: models.EdgeCluster{Name: cuid.New()}}, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("and edge cluster repository ReadEdgeCluster cannot find the tenant", func() {
				It("should return TenantNotFoundError", func() {
					expectedError := repositoryContract.NewTenantNotFoundError(request.TenantID)
					mockEdgeClusterRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertTenantNotFoundError(request.TenantID, response.Err, expectedError)
				})
			})

			When("and edge cluster repository ReadEdgeCluster cannot find the edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repositoryContract.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockEdgeClusterRepositoryService.
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
					mockEdgeClusterRepositoryService.
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
					edgeCluster := models.EdgeCluster{Name: cuid.New()}
					mockEdgeClusterRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repositoryContract.ReadEdgeClusterResponse{EdgeCluster: edgeCluster}, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					Ω(response.EdgeCluster).ShouldNot(BeNil())
					Ω(response.EdgeCluster.Name).Should(Equal(edgeCluster.Name))
				})
			})
		})
	})

	Describe("UpdateEdgeCluster", func() {
		var (
			request contract.UpdateEdgeClusterRequest
		)

		BeforeEach(func() {
			request = contract.UpdateEdgeClusterRequest{
				TenantID:      cuid.New(),
				EdgeClusterID: cuid.New(),
				EdgeCluster:   models.EdgeCluster{Name: cuid.New()},
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("UpdateEdgeCluster is called without context", func() {
				It("should return ArgumentNilError", func() {
					response, err := sut.UpdateEdgeCluster(nil, &request)
					Ω(err).Should(BeNil())
					assertArgumentNilError("ctx", "", response.Err)
				})
			})

			When("UpdateEdgeCluster is called without request", func() {
				It("should return ArgumentNilError", func() {
					response, err := sut.UpdateEdgeCluster(ctx, nil)
					Ω(err).Should(BeNil())
					assertArgumentNilError("request", "", response.Err)
				})
			})

			When("UpdateEdgeCluster is called with invalid request", func() {
				It("should return ArgumentNilError", func() {
					invalidRequest := contract.UpdateEdgeClusterRequest{
						EdgeClusterID: "",
						EdgeCluster:   models.EdgeCluster{Name: ""},
					}

					response, err := sut.UpdateEdgeCluster(ctx, &invalidRequest)
					Ω(err).Should(BeNil())
					validationErr := invalidRequest.Validate()
					assertArgumentError("request", validationErr.Error(), response.Err, validationErr)
				})
			})

			When("UpdateEdgeCluster is called with correct input parameters", func() {
				It("should call edge cluster repository UpdateEdgeCluster method", func() {
					mockEdgeClusterRepositoryService.
						EXPECT().
						UpdateEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.UpdateEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
							Ω(mappedRequest.EdgeCluster.Name).Should(Equal(request.EdgeCluster.Name))
						}).
						Return(&repositoryContract.UpdateEdgeClusterResponse{}, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("and edge cluster repository UpdateEdgeCluster cannot find provided tenant", func() {
				It("should return TenantNotFoundError", func() {
					expectedError := repositoryContract.NewTenantNotFoundError(request.TenantID)
					mockEdgeClusterRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertTenantNotFoundError(request.TenantID, response.Err, expectedError)
				})
			})

			When("and edge cluster repository UpdateEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repositoryContract.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockEdgeClusterRepositoryService.
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
					mockEdgeClusterRepositoryService.
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
					mockEdgeClusterRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repositoryContract.UpdateEdgeClusterResponse{}, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})
		})
	})

	Describe("DeleteEdgeCluster is called.", func() {
		var (
			request contract.DeleteEdgeClusterRequest
		)

		BeforeEach(func() {
			request = contract.DeleteEdgeClusterRequest{
				TenantID:      cuid.New(),
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("context is null", func() {
				It("should return ArgumentNilError and ArgumentName matches the context argument name", func() {
					response, err := sut.DeleteEdgeCluster(nil, &request)
					Ω(err).Should(BeNil())
					assertArgumentNilError("ctx", "", response.Err)
				})
			})

			When("request is null", func() {
				It("should return ArgumentNilError and ArgumentName matches the request argument name", func() {
					response, err := sut.DeleteEdgeCluster(ctx, nil)
					Ω(err).Should(BeNil())
					assertArgumentNilError("request", "", response.Err)
				})
			})

			When("request is invalid", func() {
				It("should return ArgumentNilError and both ArgumentName and ErrorMessage are matched", func() {
					invalidRequest := contract.DeleteEdgeClusterRequest{
						EdgeClusterID: "",
					}

					response, err := sut.DeleteEdgeCluster(ctx, &invalidRequest)
					Ω(err).Should(BeNil())
					validationErr := invalidRequest.Validate()
					assertArgumentError("request", validationErr.Error(), response.Err, validationErr)
				})
			})

			When("input parameters are valid", func() {
				It("should call edge cluster repository DeleteEdgeCluster method", func() {
					mockEdgeClusterRepositoryService.
						EXPECT().
						DeleteEdgeCluster(ctx, gomock.Any()).
						Do(func(_ context.Context, mappedRequest *repositoryContract.DeleteEdgeClusterRequest) {
							Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
						}).
						Return(&repositoryContract.DeleteEdgeClusterResponse{}, nil)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository DeleteEdgeCluster cannot find provided tenant", func() {
				It("should return TenantNotFoundError", func() {
					expectedError := repositoryContract.NewTenantNotFoundError(request.TenantID)
					mockEdgeClusterRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					assertTenantNotFoundError(request.TenantID, response.Err, expectedError)
				})
			})

			When("edge cluster repository DeleteEdgeCluster cannot find provided edge cluster", func() {
				It("should return EdgeClusterNotFoundError", func() {
					expectedError := repositoryContract.NewEdgeClusterNotFoundError(request.EdgeClusterID)
					mockEdgeClusterRepositoryService.
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
					mockEdgeClusterRepositoryService.
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
					mockEdgeClusterRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&repositoryContract.DeleteEdgeClusterResponse{}, nil)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})
		})
	})
})

func TestEdgeClusterService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EdgeClusterService Tests")
}

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

func assertArgumentError(expectedArgumentName, expectedMessage string, err error, nestedErr error) {
	Ω(commonErrors.IsArgumentError(err)).Should(BeTrue())

	var argumentErr commonErrors.ArgumentError
	_ = errors.As(err, &argumentErr)

	Ω(argumentErr.ArgumentName).Should(Equal(expectedArgumentName))
	Ω(strings.Contains(argumentErr.Error(), expectedMessage)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertUnknowError(expectedMessage string, err error, nestedErr error) {
	Ω(contract.IsUnknownError(err)).Should(BeTrue())

	var unknownErr contract.UnknownError
	_ = errors.As(err, &unknownErr)

	Ω(strings.Contains(unknownErr.Error(), expectedMessage)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeClusterAlreadyExistsError(err error, nestedErr error) {
	Ω(contract.IsEdgeClusterAlreadyExistsError(err)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertTenantNotFoundError(expectedTenantID string, err error, nestedErr error) {
	Ω(contract.IsTenantNotFoundError(err)).Should(BeTrue())

	var tenantNotFoundErr contract.TenantNotFoundError
	_ = errors.As(err, &tenantNotFoundErr)

	Ω(tenantNotFoundErr.TenantID).Should(Equal(expectedTenantID))
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}

func assertEdgeClusterNotFoundError(expectedEdgeClusterID string, err error, nestedErr error) {
	Ω(contract.IsEdgeClusterNotFoundError(err)).Should(BeTrue())

	var edgeClusterNotFoundErr contract.EdgeClusterNotFoundError
	_ = errors.As(err, &edgeClusterNotFoundErr)

	Ω(edgeClusterNotFoundErr.EdgeClusterID).Should(Equal(expectedEdgeClusterID))
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}
