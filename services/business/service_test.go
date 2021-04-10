package business_test

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/decentralized-cloud/edge-cluster/models"
	"github.com/decentralized-cloud/edge-cluster/services/business"
	edgeClusterTypes "github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	edgeClusterFactoryMock "github.com/decentralized-cloud/edge-cluster/services/edgecluster/types/mock"
	repository "github.com/decentralized-cloud/edge-cluster/services/repository"
	repsoitoryMock "github.com/decentralized-cloud/edge-cluster/services/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/lucsky/cuid"
	"github.com/micro-business/go-core/common"
	commonErrors "github.com/micro-business/go-core/system/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBusinessService(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	RegisterFailHandler(Fail)
	RunSpecs(t, "Business Service Tests")
}

var _ = Describe("Business Service Tests", func() {
	var (
		mockCtrl                          *gomock.Controller
		sut                               business.BusinessContract
		mockRepositoryService             *repsoitoryMock.MockRepositoryContract
		mockEdgeClusterProvisionerService *edgeClusterFactoryMock.MockEdgeClusterProvisionerContract
		mockEdgeClusterFactoryService     *edgeClusterFactoryMock.MockEdgeClusterFactoryContract
		ctx                               context.Context
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())

		mockRepositoryService = repsoitoryMock.NewMockRepositoryContract(mockCtrl)

		mockEdgeClusterProvisionerService = edgeClusterFactoryMock.NewMockEdgeClusterProvisionerContract(mockCtrl)
		mockEdgeClusterProvisionerService.
			EXPECT().
			CreateProvision(ctx, gomock.Any()).
			Return(&edgeClusterTypes.CreateProvisionResponse{}, nil).
			AnyTimes()
		mockEdgeClusterProvisionerService.
			EXPECT().
			UpdateProvisionWithRetry(ctx, gomock.Any()).
			Return(&edgeClusterTypes.UpdateProvisionResponse{}, nil).
			AnyTimes()
		mockEdgeClusterProvisionerService.
			EXPECT().
			DeleteProvision(ctx, gomock.Any()).
			Return(&edgeClusterTypes.DeleteProvisionResponse{}, nil).
			AnyTimes()

		mockEdgeClusterProvisionerService.
			EXPECT().
			GetProvisionDetails(ctx, gomock.Any()).
			Return(&edgeClusterTypes.GetProvisionDetailsResponse{}, nil).
			AnyTimes()

		mockEdgeClusterFactoryService = edgeClusterFactoryMock.NewMockEdgeClusterFactoryContract(mockCtrl)
		mockEdgeClusterFactoryService.
			EXPECT().
			Create(ctx, models.K3S).
			Return(mockEdgeClusterProvisionerService, nil).
			AnyTimes()

		sut, _ = business.NewBusinessService(
			mockRepositoryService,
			mockEdgeClusterFactoryService,
		)
		ctx = context.Background()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("user tries to instantiate BusinessService", func() {
		When("edge cluster repository service is not provided and NewBusinessService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := business.NewBusinessService(
					nil,
					mockEdgeClusterFactoryService)
				Ω(service).Should(BeNil())
				assertArgumentNilError("repositoryService", "", err)
			})
		})

		When("edge cluster factory service is not provided and NewBusinessService is called", func() {
			It("should return ArgumentNilError", func() {
				service, err := business.NewBusinessService(
					mockRepositoryService,
					nil)
				Ω(service).Should(BeNil())
				assertArgumentNilError("edgeClusterFactoryService", "", err)
			})
		})

		When("all dependencies are resolved and NewBusinessService is called", func() {
			It("should instantiate the new BusinessService", func() {
				service, err := business.NewBusinessService(
					mockRepositoryService,
					mockEdgeClusterFactoryService)
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
				EdgeCluster: models.EdgeCluster{
					ProjectID:     cuid.New(),
					Name:          cuid.New(),
					ClusterSecret: cuid.New(),
					ClusterType:   models.K3S,
				}}
		})

		Context("edge cluster service is instantiated", func() {
			When("CreateEdgeCluster is called", func() {
				It("should call edge cluster repository CreateEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						CreateEdgeCluster(ctx, gomock.Any()).
						DoAndReturn(
							func(
								_ context.Context,
								mappedRequest *repository.CreateEdgeClusterRequest) (*repository.CreateEdgeClusterResponse, error) {
								Ω(mappedRequest.EdgeCluster).Should(Equal(request.EdgeCluster))

								return &repository.CreateEdgeClusterResponse{}, nil
							})

					response, err := sut.CreateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})

				When("edge cluster repository CreateEdgeCluster returns error", func() {
					It("should return the same error", func() {
						expectedError := errors.New(cuid.New())
						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(nil, expectedError)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						Ω(response.Err).Should(Equal(expectedError))
					})
				})

				When("edge cluster repository CreateEdgeCluster return no error", func() {
					It("should return expected details", func() {
						expectedResponse := repository.CreateEdgeClusterResponse{
							EdgeClusterID: cuid.New(),
							EdgeCluster: models.EdgeCluster{
								ProjectID:     cuid.New(),
								Name:          cuid.New(),
								ClusterSecret: cuid.New(),
								ClusterType:   models.K3S,
							},
							Cursor: cuid.New(),
						}

						mockRepositoryService.
							EXPECT().
							CreateEdgeCluster(gomock.Any(), gomock.Any()).
							Return(&expectedResponse, nil)

						response, err := sut.CreateEdgeCluster(ctx, &request)
						Ω(err).Should(BeNil())
						Ω(response.Err).Should(BeNil())
						Ω(response.EdgeClusterID).ShouldNot(BeNil())
						Ω(response.EdgeClusterID).Should(Equal(expectedResponse.EdgeClusterID))
						Ω(response.EdgeCluster).Should(Equal(expectedResponse.EdgeCluster))
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
			When("ReadEdgeCluster is called", func() {
				It("should call edge cluster repository ReadEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(ctx, gomock.Any()).
						DoAndReturn(
							func(
								_ context.Context,
								mappedRequest *repository.ReadEdgeClusterRequest) (*repository.ReadEdgeClusterResponse, error) {
								Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))

								return &repository.ReadEdgeClusterResponse{
									EdgeCluster: models.EdgeCluster{
										Name:          cuid.New(),
										ProjectID:     cuid.New(),
										ClusterSecret: cuid.New(),
										ClusterType:   models.K3S,
									}}, nil
							})

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository ReadEdgeCluster returns error", func() {
				It("should return the same error", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(Equal(expectedError))
				})
			})

			When("edge cluster repository ReadEdgeCluster return no error", func() {
				It("should return the edgeClusterID", func() {
					expectedResponse := repository.ReadEdgeClusterResponse{
						EdgeCluster: models.EdgeCluster{
							Name: cuid.New(),
						}}
					mockRepositoryService.
						EXPECT().
						ReadEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&expectedResponse, nil)

					response, err := sut.ReadEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					Ω(response.EdgeCluster).Should(Equal(expectedResponse.EdgeCluster))
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
				EdgeCluster: models.EdgeCluster{
					Name:          cuid.New(),
					ProjectID:     cuid.New(),
					ClusterSecret: cuid.New(),
					ClusterType:   models.K3S,
				},
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("UpdateEdgeCluster is called", func() {
				It("should call edge cluster repository UpdateEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(ctx, gomock.Any()).
						DoAndReturn(
							func(
								_ context.Context,
								mappedRequest *repository.UpdateEdgeClusterRequest) (*repository.UpdateEdgeClusterResponse, error) {
								Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
								Ω(mappedRequest.EdgeCluster.Name).Should(Equal(request.EdgeCluster.Name))
								Ω(mappedRequest.EdgeCluster.ProjectID).Should(Equal(request.EdgeCluster.ProjectID))
								Ω(mappedRequest.EdgeCluster.ClusterSecret).Should(Equal(request.EdgeCluster.ClusterSecret))
								Ω(mappedRequest.EdgeCluster.ClusterType).Should(Equal(request.EdgeCluster.ClusterType))

								return &repository.UpdateEdgeClusterResponse{}, nil
							})

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository UpdateEdgeCluster returns error", func() {
				It("should return the same error", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(Equal(expectedError))
				})
			})

			When("edge cluster repository UpdateEdgeCluster return no error", func() {
				It("should return expected details", func() {
					expectedResponse := repository.UpdateEdgeClusterResponse{
						EdgeCluster: models.EdgeCluster{
							Name:          cuid.New(),
							ProjectID:     cuid.New(),
							ClusterSecret: cuid.New(),
							ClusterType:   models.K3S,
						},
						Cursor: cuid.New(),
					}
					mockRepositoryService.
						EXPECT().
						UpdateEdgeCluster(gomock.Any(), gomock.Any()).
						Return(&expectedResponse, nil)

					response, err := sut.UpdateEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					Ω(response.EdgeCluster).Should(Equal(expectedResponse.EdgeCluster))
				})
			})
		})
	})

	Describe("DeleteEdgeCluster is called", func() {
		var (
			request business.DeleteEdgeClusterRequest
		)

		BeforeEach(func() {
			request = business.DeleteEdgeClusterRequest{
				EdgeClusterID: cuid.New(),
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("DeleteEdgeCluster is called", func() {
				It("should call edge cluster repository DeleteEdgeCluster method", func() {
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(ctx, gomock.Any()).
						DoAndReturn(
							func(
								_ context.Context,
								mappedRequest *repository.DeleteEdgeClusterRequest) (*repository.DeleteEdgeClusterResponse, error) {
								Ω(mappedRequest.EdgeClusterID).Should(Equal(request.EdgeClusterID))
								return &repository.DeleteEdgeClusterResponse{}, nil
							})

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository DeleteEdgeCluster returns error", func() {
				It("should return the same error", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						DeleteEdgeCluster(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.DeleteEdgeCluster(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(Equal(expectedError))
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

	Describe("ListEdgeClusters is called", func() {
		var (
			request        business.ListEdgeClustersRequest
			edgeClusterIDs []string
			projectIDs     []string
		)

		BeforeEach(func() {
			edgeClusterIDs = []string{}
			for idx := 0; idx < rand.Intn(20)+1; idx++ {
				edgeClusterIDs = append(edgeClusterIDs, cuid.New())
			}

			projectIDs = []string{}
			for idx := 0; idx < rand.Intn(20)+1; idx++ {
				projectIDs = append(projectIDs, cuid.New())
			}

			request = business.ListEdgeClustersRequest{
				Pagination: common.Pagination{
					After:  convertStringToPointer(cuid.New()),
					First:  convertIntToPointer(rand.Intn(1000)),
					Before: convertStringToPointer(cuid.New()),
					Last:   convertIntToPointer(rand.Intn(1000)),
				},
				SortingOptions: []common.SortingOptionPair{
					{
						Name:      cuid.New(),
						Direction: common.Ascending,
					},
					{
						Name:      cuid.New(),
						Direction: common.Descending,
					},
				},
				EdgeClusterIDs: edgeClusterIDs,
				ProjectIDs:     projectIDs,
			}
		})

		Context("edge cluster service is instantiated", func() {
			When("ListEdgeClusters is called", func() {
				It("should call edge cluster repository ListEdgeClusters method", func() {
					mockRepositoryService.
						EXPECT().
						ListEdgeClusters(ctx, gomock.Any()).
						DoAndReturn(
							func(
								_ context.Context,
								mappedRequest *repository.ListEdgeClustersRequest) (*repository.ListEdgeClustersResponse, error) {
								Ω(mappedRequest.Pagination).Should(Equal(request.Pagination))
								Ω(mappedRequest.SortingOptions).Should(Equal(request.SortingOptions))
								Ω(mappedRequest.EdgeClusterIDs).Should(Equal(request.EdgeClusterIDs))
								Ω(mappedRequest.ProjectIDs).Should(Equal(request.ProjectIDs))

								return &repository.ListEdgeClustersResponse{}, nil
							})

					response, err := sut.ListEdgeClusters(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
				})
			})

			When("edge cluster repository ListEdgeClusters returns error", func() {
				It("should return the same error", func() {
					expectedError := errors.New(cuid.New())
					mockRepositoryService.
						EXPECT().
						ListEdgeClusters(gomock.Any(), gomock.Any()).
						Return(nil, expectedError)

					response, err := sut.ListEdgeClusters(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(Equal(expectedError))
				})
			})

			When("edge cluster repository ListEdgeClusters completes successfully", func() {
				It("should return the list of matched edgeClusterIDs", func() {
					edgeClusters := []models.EdgeClusterWithCursor{}

					for idx := 0; idx < rand.Intn(20)+1; idx++ {
						edgeClusters = append(edgeClusters, models.EdgeClusterWithCursor{
							EdgeClusterID: cuid.New(),
							EdgeCluster: models.EdgeCluster{
								ProjectID:     cuid.New(),
								Name:          cuid.New(),
								ClusterSecret: cuid.New(),
								ClusterType:   models.K3S,
							},
							Cursor: cuid.New(),
						})
					}

					expectedResponse := repository.ListEdgeClustersResponse{
						HasPreviousPage: (rand.Intn(10) % 2) == 0,
						HasNextPage:     (rand.Intn(10) % 2) == 0,
						TotalCount:      rand.Int63n(1000),
						EdgeClusters:    edgeClusters,
					}

					mockRepositoryService.
						EXPECT().
						ListEdgeClusters(gomock.Any(), gomock.Any()).
						Return(&expectedResponse, nil)

					response, err := sut.ListEdgeClusters(ctx, &request)
					Ω(err).Should(BeNil())
					Ω(response.Err).Should(BeNil())
					Ω(response.HasPreviousPage).Should(Equal(expectedResponse.HasPreviousPage))
					Ω(response.HasNextPage).Should(Equal(expectedResponse.HasNextPage))
					Ω(response.TotalCount).Should(Equal(expectedResponse.TotalCount))
					Ω(response.EdgeClusters).Should(Equal(expectedResponse.EdgeClusters))
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
func convertStringToPointer(str string) *string {
	return &str
}

func convertIntToPointer(i int) *int {
	return &i
}
