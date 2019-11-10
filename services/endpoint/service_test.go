package endpoint_test

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"testing"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	businessMock "github.com/decentralized-cloud/edge-cluster/services/business/mock"
	"github.com/decentralized-cloud/edge-cluster/services/endpoint"
	gokitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	clusterPublicIPAddress = "10.0.0.230"
)

func TestEndpointCreatorService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Endpoint Creator Service Tests")
}

var _ = Describe("Endpoint Creator Service Tests", func() {
	var (
		mockCtrl            *gomock.Controller
		sut                 endpoint.EndpointCreatorContract
		mockBusinessService *businessMock.MockBusinessContract
		ctx                 context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockBusinessService = businessMock.NewMockBusinessContract(mockCtrl)
		sut, _ = endpoint.NewEndpointCreatorService(mockBusinessService)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate EndpointCreatorService", func() {
		When("edge cluster business service is not provided and NewEndpointCreatorService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := endpoint.NewEndpointCreatorService(nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("businessService", "", err)
			})
		})

		When("all dependencies are resolved and NewEndpointCreatorService is called", func() {
			It("should instantiate the new EndpointCreatorService", func() {
				service, err := endpoint.NewEndpointCreatorService(mockBusinessService)
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
				endpoint gokitendpoint.Endpoint
				request  business.CreateEdgeClusterRequest
				response business.CreateEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.CreateEdgeClusterEndpoint()
				request = business.CreateEdgeClusterRequest{
					EdgeCluster: models.EdgeCluster{
						TenantID:               cuid.New(),
						Name:                   cuid.New(),
						ClusterPublicIPAddress: clusterPublicIPAddress,
					},
				}

				response = business.CreateEdgeClusterResponse{
					EdgeClusterID: cuid.New(),
				}
			})

			Context("CreateEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.CreateEdgeClusterRequest{
							EdgeCluster: models.EdgeCluster{
								Name: "",
							}}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service CreateEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							CreateEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.CreateEdgeClusterRequest) (*business.CreateEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.CreateEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service CreateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service CreateEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
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
				endpoint gokitendpoint.Endpoint
				request  business.ReadEdgeClusterRequest
				response business.ReadEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.ReadEdgeClusterEndpoint()
				request = business.ReadEdgeClusterRequest{
					EdgeClusterID: cuid.New(),
				}

				response = business.ReadEdgeClusterResponse{
					EdgeCluster: models.EdgeCluster{
						TenantID:               cuid.New(),
						Name:                   cuid.New(),
						ClusterPublicIPAddress: clusterPublicIPAddress,
					},
				}
			})

			Context("ReadEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.ReadEdgeClusterRequest{
							EdgeClusterID: "",
						}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service ReadEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							ReadEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.ReadEdgeClusterRequest) (*business.ReadEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.ReadEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service ReadEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							ReadEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service ReadEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
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
				endpoint gokitendpoint.Endpoint
				request  business.UpdateEdgeClusterRequest
				response business.UpdateEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.UpdateEdgeClusterEndpoint()
				request = business.UpdateEdgeClusterRequest{
					EdgeClusterID: cuid.New(),
					EdgeCluster: models.EdgeCluster{
						TenantID:               cuid.New(),
						Name:                   cuid.New(),
						ClusterPublicIPAddress: clusterPublicIPAddress,
					}}

				response = business.UpdateEdgeClusterResponse{}
			})

			Context("UpdateEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.UpdateEdgeClusterRequest{
							EdgeClusterID: "",
							EdgeCluster: models.EdgeCluster{
								Name: "",
							}}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service UpdateEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							UpdateEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.UpdateEdgeClusterRequest) (*business.UpdateEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.UpdateEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service UpdateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							UpdateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service UpdateEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
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
				endpoint gokitendpoint.Endpoint
				request  business.DeleteEdgeClusterRequest
				response business.DeleteEdgeClusterResponse
			)

			BeforeEach(func() {
				endpoint = sut.DeleteEdgeClusterEndpoint()
				request = business.DeleteEdgeClusterRequest{
					EdgeClusterID: cuid.New(),
				}

				response = business.DeleteEdgeClusterResponse{}
			})

			Context("DeleteEdgeClusterEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with invalid request", func() {
					It("should return ArgumentNilError", func() {
						invalidRequest := business.DeleteEdgeClusterRequest{
							EdgeClusterID: "",
						}
						returnedResponse, err := endpoint(ctx, &invalidRequest)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						validationErr := invalidRequest.Validate()
						assertArgumentError("request", validationErr.Error(), castedResponse.Err, validationErr)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service DeleteEdgeCluster method", func() {
						mockBusinessService.
							EXPECT().
							DeleteEdgeCluster(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.DeleteEdgeClusterRequest) (*business.DeleteEdgeClusterResponse, error) {
									Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.DeleteEdgeClusterResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service DeleteEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							DeleteEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service DeleteEdgeCluster returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
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

	Context("EndpointCreatorService is instantiated", func() {
		When("SearchEndpoint is called", func() {
			It("should return valid function", func() {
				endpoint := sut.SearchEndpoint()
				Ω(endpoint).ShouldNot(BeNil())
			})

			var (
				endpoint       gokitendpoint.Endpoint
				edgeClusterIDs []string
				tenantIDs      []string
				request        business.SearchRequest
				response       business.SearchResponse
			)

			BeforeEach(func() {
				endpoint = sut.SearchEndpoint()
				rand.Seed(42)
				edgeClusterIDs = []string{}
				for idx := 0; idx < rand.Intn(20)+1; idx++ {
					edgeClusterIDs = append(edgeClusterIDs, cuid.New())
				}

				tenantIDs = []string{}
				for idx := 0; idx < rand.Intn(20)+1; idx++ {
					tenantIDs = append(tenantIDs, cuid.New())
				}

				request = business.SearchRequest{
					Pagination: common.Pagination{
						After:  cuid.New(),
						First:  rand.Intn(1000),
						Before: cuid.New(),
						Last:   rand.Intn(1000),
					},
					SortingOptions: []common.SortingOptionPair{
						common.SortingOptionPair{
							Name:      cuid.New(),
							Direction: common.Ascending,
						},
						common.SortingOptionPair{
							Name:      cuid.New(),
							Direction: common.Descending,
						},
					},
					EdgeClusterIDs: edgeClusterIDs,
					TenantIDs:      tenantIDs,
				}

				edgeClusters := []models.EdgeClusterWithCursor{}

				for idx := 0; idx < rand.Intn(20)+1; idx++ {
					edgeClusters = append(edgeClusters, models.EdgeClusterWithCursor{
						EdgeClusterID: cuid.New(),
						EdgeCluster: models.EdgeCluster{
							Name:                   cuid.New(),
							TenantID:               cuid.New(),
							ClusterPublicIPAddress: clusterPublicIPAddress,
						},
						Cursor: cuid.New(),
					})
				}

				response = business.SearchResponse{EdgeClusters: edgeClusters}
			})

			Context("SearchEndpoint function is returned", func() {
				When("endpoint is called with nil context", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(nil, &request)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.SearchResponse)
						assertArgumentNilError("ctx", "", castedResponse.Err)
					})
				})

				When("endpoint is called with nil request", func() {
					It("should return ArgumentNilError", func() {
						returnedResponse, err := endpoint(ctx, nil)

						Ω(err).Should(BeNil())
						Ω(returnedResponse).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.SearchResponse)
						assertArgumentNilError("request", "", castedResponse.Err)
					})
				})

				When("endpoint is called with valid request", func() {
					It("should call business service Search method", func() {
						mockBusinessService.
							EXPECT().
							Search(ctx, gomock.Any()).
							DoAndReturn(
								func(
									_ context.Context,
									mappedRequest *business.SearchRequest) (*business.SearchResponse, error) {
									Ω(mappedRequest.Pagination).Should(Equal(request.Pagination))
									Ω(mappedRequest.SortingOptions).Should(Equal(request.SortingOptions))
									Ω(mappedRequest.EdgeClusterIDs).Should(Equal(request.EdgeClusterIDs))
									Ω(mappedRequest.TenantIDs).Should(Equal(request.TenantIDs))

									return &response, nil
								})

						returnedResponse, err := endpoint(ctx, &request)

						Ω(err).Should(BeNil())
						Ω(response).ShouldNot(BeNil())
						castedResponse := returnedResponse.(*business.SearchResponse)
						Ω(castedResponse.Err).Should(BeNil())
					})
				})

				When("business service Search returns error", func() {
					It("should return the same error", func() {
						expectedErr := errors.New(cuid.New())
						mockBusinessService.
							EXPECT().
							Search(gomock.Any(), gomock.Any()).
							Return(nil, expectedErr)

						_, err := endpoint(ctx, &request)

						Ω(err).Should(Equal(expectedErr))
					})
				})

				When("business service Search returns response", func() {
					It("should return the same response", func() {
						mockBusinessService.
							EXPECT().
							Search(gomock.Any(), gomock.Any()).
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
