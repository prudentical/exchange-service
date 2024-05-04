package exchange_test

import (
	"errors"
	"exchange-service/internal/model"
	mock_sdk "exchange-service/internal/sdk/mock"
	mock_currency "exchange-service/internal/service/currency/mock"
	"exchange-service/internal/service/exchange"
	mock_exchange "exchange-service/internal/service/exchange/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Exchange setup", Label("exchange"), func() {

	var setups exchange.ExchangeSetupService
	var ctrl *gomock.Controller
	var factory *mock_sdk.MockExchangeAPIClientFactory
	var exSDK *mock_sdk.MockExchangeAPIClient
	var currencies *mock_currency.MockCurrencyService
	var exchanges *mock_exchange.MockExchangeService
	var pairs *mock_exchange.MockPairService

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		factory = mock_sdk.NewMockExchangeAPIClientFactory(ctrl)
		exSDK = mock_sdk.NewMockExchangeAPIClient(ctrl)
		currencies = mock_currency.NewMockCurrencyService(ctrl)
		exchanges = mock_exchange.NewMockExchangeService(ctrl)
		pairs = mock_exchange.NewMockPairService(ctrl)
		setups = exchange.NewExchangeSetupService(factory, currencies, exchanges, pairs)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Setup", func() {
		Context("with no exchange", func() {
			It("should do nothing", func() {
				exchanges.EXPECT().GetAll().Return([]model.Exchange{}, nil)
				factory.EXPECT().Create(gomock.Any()).Times(0)
				setups.Setup()
			})
		})
		Context("with error while getting exchanges", func() {
			It("should return the error", func() {
				exchanges.EXPECT().GetAll().Return([]model.Exchange{}, errors.New(""))
				factory.EXPECT().Create(gomock.Any()).Times(0)
				err := setups.Setup()
				Expect(err).NotTo(BeNil())
			})
		})
		Context("with error while creating sdk service", func() {
			It("should return the error", func() {
				exchanges.EXPECT().GetAll().Return([]model.Exchange{{}}, nil)
				factory.EXPECT().Create(gomock.Any()).Return(exSDK, errors.New(""))
				exSDK.EXPECT().Currencies().Times(0)
				err := setups.Setup()
				Expect(err).NotTo(BeNil())
			})
		})
		Context("with error while adding exchange currencies", func() {
			It("should return the error", func() {
				exchanges.EXPECT().GetAll().Return([]model.Exchange{{}}, nil)
				factory.EXPECT().Create(gomock.Any()).Return(exSDK, nil)
				exSDK.EXPECT().Currencies().Return([]model.Currency{{}}, errors.New(""))
				exSDK.EXPECT().Pairs().Times(0)
				err := setups.Setup()
				Expect(err).NotTo(BeNil())
			})
		})
		Context("with error while adding exchange paris", func() {
			It("should return the error", func() {
				exchanges.EXPECT().GetAll().Return([]model.Exchange{{}}, nil)
				factory.EXPECT().Create(gomock.Any()).Return(exSDK, nil)
				exSDK.EXPECT().Currencies().Return([]model.Currency{{}}, nil)
				currencies.EXPECT().Merge(gomock.Any()).Times(1)
				exSDK.EXPECT().Pairs().Return([]model.Pair{{}}, errors.New(""))
				pairs.EXPECT().Merge(gomock.Any()).Times(0)
				err := setups.Setup()
				Expect(err).NotTo(BeNil())
			})
		})
		Context("with valid state", func() {
			It("should setup the exchange", func() {
				exchanges.EXPECT().GetAll().Return([]model.Exchange{{}}, nil)
				factory.EXPECT().Create(gomock.Any()).Return(exSDK, nil)
				exSDK.EXPECT().Currencies().Return([]model.Currency{{}}, nil)
				currencies.EXPECT().Merge(gomock.Any()).Times(1)
				exSDK.EXPECT().Pairs().Return([]model.Pair{{}}, nil)
				pairs.EXPECT().Merge(gomock.Any()).Times(1)
				err := setups.Setup()
				Expect(err).To(BeNil())
			})
		})
	})

})
