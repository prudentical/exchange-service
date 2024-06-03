package exchange_test

import (
	"errors"
	"exchange-service/internal/dto"
	"exchange-service/internal/model"
	mock_persistence "exchange-service/internal/persistence/mock"
	"exchange-service/internal/sdk"
	mock_sdk "exchange-service/internal/sdk/mock"
	"exchange-service/internal/service"
	"exchange-service/internal/service/exchange"
	mock_exchange "exchange-service/internal/service/exchange/mock"
	mock_service "exchange-service/internal/service/mock"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Exchange order", Label("exchange"), func() {

	var orders exchange.OrderService
	var ctrl *gomock.Controller
	var dao *mock_persistence.MockExchangeDAO
	var orderService *mock_service.MockOrderService
	var pairs *mock_exchange.MockPairService
	var exSDK *mock_sdk.MockExchangeAPIClient

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		dao = mock_persistence.NewMockExchangeDAO(ctrl)
		pairs = mock_exchange.NewMockPairService(ctrl)
		exSDK = mock_sdk.NewMockExchangeAPIClient(ctrl)
		orderService = mock_service.NewMockOrderService(ctrl)
		orders = exchange.NewOrderService(dao, orderService, pairs)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Virtual order", func() {
		Context("with zero amount", func() {
			It("should return a not-found error for pair", func() {
				request := dto.OrderDTO{
					IsVirtual: false,
					Amount:    decimal.Zero,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).To(MatchError(service.OrderAmountRequiredError{}))
			})
		})
		Context("with invalid pair", func() {
			It("should return a not-found error for pair", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).
					Return(model.Pair{}, service.NotFoundError{})
				request := dto.OrderDTO{
					IsVirtual: false,
					Amount:    decimal.NewFromInt(100),
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).To(MatchError(service.NotFoundError{}))
			})
		})
		Context("with historic price", func() {
			When("price retrieval fails", func() {
				It("should return error", func() {
					exSDK.EXPECT().GetExchange().Return(model.Exchange{})
					pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
					exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).
						Return(decimal.NewFromFloat(10.0), errors.New(""))
					now := time.Now()
					request := dto.OrderDTO{
						IsVirtual: true,
						Amount:    decimal.NewFromInt(1),
						DateTime:  &now,
					}
					err := orders.Order(exSDK, 0, request)
					Expect(err).NotTo(BeNil())
				})
			})
			When("when order is valid", func() {
				It("should place the order", func() {
					exSDK.EXPECT().GetExchange().Return(model.Exchange{})
					pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
					exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).
						Return(decimal.NewFromFloat(10.0), nil)
					var savedOrder dto.OrderDTO
					orderService.EXPECT().SaveOrder(gomock.Any()).
						Do(func(arg dto.OrderDTO) { savedOrder = arg })
					now := time.Now()
					request := dto.OrderDTO{
						IsVirtual: true,
						Amount:    decimal.NewFromInt(1),
						DateTime:  &now,
					}
					err := orders.Order(exSDK, 0, request)
					Expect(err).To(BeNil())
					Expect(savedOrder.Status).To(Equal(dto.Fulfilled))
					Expect(savedOrder.InternalId).To(Equal("VIRT"))
					Expect(savedOrder.FilledAmount.String()).To(Equal(request.Amount.String()))
				})
			})
		})
		Context("with realtime price", func() {
			When("price retrieval fails", func() {
				It("should return error", func() {
					exSDK.EXPECT().GetExchange().Return(model.Exchange{})
					pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
					exSDK.EXPECT().PriceFor(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
						Return(decimal.NewFromFloat(10.0), errors.New(""))
					request := dto.OrderDTO{
						IsVirtual: true,
						Amount:    decimal.NewFromInt(1),
					}
					err := orders.Order(exSDK, 0, request)
					Expect(err).NotTo(BeNil())
				})
			})
			When("when order is valid", func() {
				It("should place the order", func() {
					exSDK.EXPECT().GetExchange().Return(model.Exchange{})
					pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
					exSDK.EXPECT().
						PriceFor(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
						Return(decimal.NewFromFloat(10.0), nil)
					var savedOrder dto.OrderDTO
					orderService.EXPECT().SaveOrder(gomock.Any()).
						Do(func(arg dto.OrderDTO) { savedOrder = arg })
					request := dto.OrderDTO{
						IsVirtual: true,
						Amount:    decimal.NewFromInt(1),
					}
					err := orders.Order(exSDK, 0, request)
					Expect(err).To(BeNil())
					Expect(savedOrder.Status).To(Equal(dto.Fulfilled))
					Expect(savedOrder.InternalId).To(Equal("VIRT"))
					Expect(savedOrder.FilledAmount.String()).To(Equal(request.Amount.String()))
				})
			})
		})
	})
	Describe("Actual order", func() {
		It("should return a not-implemented error", func() {
			exSDK.EXPECT().GetExchange().Return(model.Exchange{})
			pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
			request := dto.OrderDTO{
				IsVirtual: false,
				Type:      sdk.Buy,
				Amount:    decimal.NewFromInt(1),
			}
			err := orders.Order(exSDK, 0, request)
			Expect(err.Error()).To(Equal("not implemented"))
		})
	})
})
