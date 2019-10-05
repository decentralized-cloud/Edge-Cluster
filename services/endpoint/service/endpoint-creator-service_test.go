package service_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/lucsky/cuid"

	"github.com/decentralized-cloud/edge-cluster/models"
	businessContract "github.com/decentralized-cloud/edge-cluster/services/business/contract"
	businessMock "github.com/decentralized-cloud/edge-cluster/services/business/mock"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint/contract"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint/service"
	"github.com/golang/mock/gomock"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestEndpointCreatorService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EndpointCreatorService Tests")
}

var _ = Describe("EndpointCreatorService Tests", func() {
	var (
		mockCtrl                       *gomock.Controller
		sut                            contract.EndpointCreatorContract
		mockEdgeClusterBusinessService *businessMock.MockEdgeClusterServiceContract
		ctx                            context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockEdgeClusterBusinessService = businessMock.NewMockEdgeClusterServiceContract(mockCtrl)
		sut, _ = service.NewEndpointCreatorService(mockEdgeClusterBusinessService)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate EndpointCreatorService", func() {
		When("edge cluster business service is not provided and NewEndpointCreatorService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := service.NewEndpointCreatorService(nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("businessService", "", err)
			})
		})

		When("all dependencies are resolved and NewEndpointCreatorService is called", func() {
			It("should instantiate the new EndpointCreatorService", func() {
				service, err := service.NewEndpointCreatorService(mockEdgeClusterBusinessService)
				Ω(err).Should(BeNil())
				Ω(service).ShouldNot(BeNil())
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("CreateEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.CreateEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint endpoint.Endpoint
				request  businessContract.CreateEdgeClusterRequest
				response businessContract.CreateEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.CreateEdgeClusterEndpoint()
				request = businessContract.CreateEdgeClusterRequest{
					TenantID: cuid.New(),
					EdgeCluster: models.EdgeCluster{
						Name: cuid.New(),
					},
				}

				response = businessContract.CreateEdgeClusterResponse{
					EdgeClusterID: cuid.New(),
				}
			})

			Context("CreateEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.CreateEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.CreateEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := businessContract.CreateEdgeClusterRequest{
							TenantID: "",
							EdgeCluster: models.EdgeCluster{
								Name: "",
							}}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.CreateEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service CreateEdgeCluster method", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							CreateEdgeCluster(ctx, gomock.Any()).
							Do(func(_ context.Context, mappedRequest *businessContract.CreateEdgeClusterRequest) {
								Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))
							}).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.CreateEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service CreateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockEdgeClusterBusinessService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service CreateEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("ReadEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.ReadEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint endpoint.Endpoint
				request  businessContract.ReadEdgeClusterRequest
				response businessContract.ReadEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.ReadEdgeClusterEndpoint()
				request = businessContract.ReadEdgeClusterRequest{
					TenantID:      cuid.New(),
					EdgeClusterID: cuid.New(),
				}

				response = businessContract.ReadEdgeClusterResponse{
					EdgeCluster: models.EdgeCluster{
						Name: cuid.New(),
					},
				}
			})

			Context("ReadEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.ReadEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.ReadEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := businessContract.ReadEdgeClusterRequest{
							TenantID:      "",
							EdgeClusterID: "",
						}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.ReadEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service ReadEdgeCluster method", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							ReadEdgeCluster(ctx, gomock.Any()).
							Do(func(_ context.Context, mappedRequest *businessContract.ReadEdgeClusterRequest) {
								Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
							}).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.ReadEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service ReadEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockEdgeClusterBusinessService.
							EXPECT().
							ReadEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service ReadEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							ReadEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("UpdateEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.UpdateEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint endpoint.Endpoint
				request  businessContract.UpdateEdgeClusterRequest
				response businessContract.UpdateEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.UpdateEdgeClusterEndpoint()
				request = businessContract.UpdateEdgeClusterRequest{
					TenantID:      cuid.New(),
					EdgeClusterID: cuid.New(),
					EdgeCluster: models.EdgeCluster{
						Name: cuid.New(),
					}}

				response = businessContract.UpdateEdgeClusterResponse{}
			})

			Context("UpdateEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.UpdateEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.UpdateEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := businessContract.UpdateEdgeClusterRequest{
							TenantID:      "",
							EdgeClusterID: "",
							EdgeCluster: models.EdgeCluster{
								Name: "",
							}}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.UpdateEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service UpdateEdgeCluster method", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							UpdateEdgeCluster(ctx, gomock.Any()).
							Do(func(_ context.Context, mappedRequest *businessContract.UpdateEdgeClusterRequest) {
								Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
							}).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.UpdateEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service UpdateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockEdgeClusterBusinessService.
							EXPECT().
							UpdateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service UpdateEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							UpdateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
				})
			})
		})
	})

	Context("EndpointCreatorService is instantiated", func() {
		When("DeleteEdgeClusterEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.DeleteEdgeClusterEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint endpoint.Endpoint
				request  businessContract.DeleteEdgeClusterRequest
				response businessContract.DeleteEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.DeleteEdgeClusterEndpoint()
				request = businessContract.DeleteEdgeClusterRequest{
					TenantID:      cuid.New(),
					EdgeClusterID: cuid.New(),
				}

				response = businessContract.DeleteEdgeClusterResponse{}
			})

			Context("DeleteEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.DeleteEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.DeleteEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := businessContract.DeleteEdgeClusterRequest{
							TenantID:      "",
							EdgeClusterID: "",
						}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.DeleteEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service DeleteEdgeCluster method", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							DeleteEdgeCluster(ctx, gomock.Any()).
							Do(func(_ context.Context, mappedRequest *businessContract.DeleteEdgeClusterRequest) {
								Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
							}).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*businessContract.DeleteEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service DeleteEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockEdgeClusterBusinessService.
							EXPECT().
							DeleteEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service DeleteEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockEdgeClusterBusinessService.
							EXPECT().
							DeleteEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&response, nil)

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).Should(Equal(&response))
					})
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

func assertArgumentError(expectedArgumentName, expectedMessage string, err error, nestedErr error) {
	Ω(commonErrors.IsArgumentError(err)).Should(BeTrue())

	var argumentErr commonErrors.ArgumentError
	_ = errors.As(err, &argumentErr)

	Ω(argumentErr.ArgumentName).Should(Equal(expectedArgumentName))
	Ω(strings.Contains(argumentErr.Error(), expectedMessage)).Should(BeTrue())
	Ω(errors.Unwrap(err)).Should(Equal(nestedErr))
}
