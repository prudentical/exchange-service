package exchange_test

import (
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	mock_sdk "exchange-service/internal/sdk/mock"
	"exchange-service/internal/service"
	"exchange-service/internal/service/exchange"
	mock_exchange "exchange-service/internal/service/exchange/mock"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Exchange price", Label("exchange"), func() {

	var prices exchange.PriceService
	var ctrl *gomock.Controller
	var pairs *mock_exchange.MockPairService
	var exSDK *mock_sdk.MockExchangeAPIClient

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		pairs = mock_exchange.NewMockPairService(ctrl)
		exSDK = mock_sdk.NewMockExchangeAPIClient(ctrl)
		prices = exchange.NewPriceService(pairs)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Check current prices", func() {
		Context("with invalid pair", func() {
			It("should return a not-found error", func() {
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, persistence.RecordNotFoundError{})
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				request := exchange.PriceCheckRequest{}
				_, err := prices.GetPrice(exSDK, 0, request)
				Expect(err).To(MatchError(service.NotFoundError{Type: model.Pair{}, Id: 0}))
			})
		})
		Context("with valid request", func() {
			It("should return the result", func() {
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				exSDK.EXPECT().PriceFor(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(decimal.NewFromFloat(10.0), nil)
				request := exchange.PriceCheckRequest{}
				result, err := prices.GetPrice(exSDK, 0, request)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(decimal.NewFromFloat(10.0)))

			})
		})
	})
	Describe("Check historic prices", func() {
		Context("with invalid pair", func() {
			It("should return a not-found error", func() {
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, persistence.RecordNotFoundError{})
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				now := time.Now()
				request := exchange.PriceCheckRequest{
					DateTime: &now,
				}
				_, err := prices.GetPrice(exSDK, 0, request)
				Expect(err).To(MatchError(service.NotFoundError{Type: model.Pair{}, Id: 0}))
			})
		})
		Context("with valid request", func() {
			It("should return the result", func() {
				pairs.EXPECT().GetById(gomock.Any(), gomock.Any()).Return(model.Pair{}, nil)
				exSDK.EXPECT().GetExchange().Return(model.Exchange{})
				exSDK.EXPECT().HistoricPrice(gomock.Any(), gomock.Any()).Return(decimal.NewFromFloat(10.0), nil)
				now := time.Now()
				request := exchange.PriceCheckRequest{
					DateTime: &now,
				}
				result, err := prices.GetPrice(exSDK, 0, request)
				Expect(err).To(BeNil())
				Expect(result).To(Equal(decimal.NewFromFloat(10.0)))
			})
		})
	})

})
