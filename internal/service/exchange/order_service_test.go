package exchange_test

import (
	"errors"
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	mock_persistence "exchange-service/internal/persistence/mock"
	"exchange-service/internal/sdk"
	mock_sdk "exchange-service/internal/sdk/mock"
	"exchange-service/internal/service"
	"exchange-service/internal/service/exchange"
	mock_exchange "exchange-service/internal/service/exchange/mock"
	mock_service "exchange-service/internal/service/mock"

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
	var exSDK *mock_sdk.MockExchangeSDK

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		dao = mock_persistence.NewMockExchangeDAO(ctrl)
		pairs = mock_exchange.NewMockPairService(ctrl)
		exSDK = mock_sdk.NewMockExchangeSDK(ctrl)
		orderService = mock_service.NewMockOrderService(ctrl)
		orders = exchange.NewOrderService(dao, pairs, orderService)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Buy order", func() {
		Context("with invalid pair", func() {
			It("should return a not-found error for pair", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, persistence.RecordNotFoundError{})
				request := exchange.OrderRequest{
					Virtual: false,
					Type:    sdk.Buy,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).To(MatchError(service.NotFoundError{Type: model.Pair{}, Id: 0}))
			})
		})
		Context("with non-virtual order", func() {
			It("should return a not-implemented error", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				request := exchange.OrderRequest{
					Virtual: false,
					Type:    sdk.Buy,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err.Error()).To(Equal("not implemented"))
			})
		})
		Context("with failed price retrieval", func() {
			It("should return error", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).Return(decimal.NewFromFloat(10.0), errors.New(""))
				request := exchange.OrderRequest{
					Virtual: true,
					Type:    sdk.Buy,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).NotTo(BeNil())
			})
		})
		Context("with valid order", func() {
			It("should place order", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).Return(decimal.NewFromFloat(10.0), nil)
				orderService.EXPECT().SaveOrder(gomock.Any())
				request := exchange.OrderRequest{
					Virtual: true,
					Type:    sdk.Buy,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("Sell order", func() {
		Context("with invalid pair", func() {
			It("should return a not-found error for pair", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, persistence.RecordNotFoundError{})
				request := exchange.OrderRequest{
					Virtual: false,
					Type:    sdk.Sell,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).To(MatchError(service.NotFoundError{Type: model.Pair{}, Id: 0}))
			})
		})
		Context("with non-virtual order", func() {
			It("should return a not-implemented error", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				request := exchange.OrderRequest{
					Virtual: false,
					Type:    sdk.Sell,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err.Error()).To(Equal("not implemented"))
			})
		})
		Context("with failed price retrieval", func() {
			It("should return error", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).Return(decimal.NewFromFloat(10.0), errors.New(""))
				request := exchange.OrderRequest{
					Virtual: true,
					Type:    sdk.Sell,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).NotTo(BeNil())
			})
		})
		Context("with valid order", func() {
			It("should place order", func() {
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).Return(decimal.NewFromFloat(10.0), nil)
				orderService.EXPECT().SaveOrder(gomock.Any())
				request := exchange.OrderRequest{
					Virtual: true,
					Type:    sdk.Sell,
				}
				err := orders.Order(exSDK, 0, request)
				Expect(err).To(BeNil())
			})
		})
	})
})
