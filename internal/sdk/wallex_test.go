package sdk_test

import (
	"exchange-service/internal/model"
	"exchange-service/internal/sdk"
	mock_sdk "exchange-service/internal/sdk/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"github.com/wallexchange/wallex-go"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Exchange API", Label("sdk"), func() {

	var exchange sdk.ExchangeAPIClient
	var ctrl *gomock.Controller
	var client *mock_sdk.MockWallexClient

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		client = mock_sdk.NewMockWallexClient(ctrl)
		exchange, _ = sdk.NewExchangeAPIClientFactory(client).Create(model.Exchange{Name: "Wallex"})
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Get buy price", func() {
		Context("with no amount or funds", func() {
			It("should return a lowest price", func() {
				client.EXPECT().MarketOrders(gomock.Any()).
					Return([]*wallex.MarketOrder{
						{
							Price:    "10",
							Quantity: "1",
						},
						{
							Price:    "12",
							Quantity: "1",
						},
					}, []*wallex.MarketOrder{}, nil)
				price, err := exchange.PriceFor(model.Pair{}, nil, nil, sdk.Buy)
				Expect(err).To(BeNil())
				Expect(price.String()).To(Equal(decimal.NewFromInt32(10).String()))
			})
		})
		Context("for amount", func() {
			Context("when multiple order needed to fill the amount", func() {
				It("should return a unit price for amount", func() {
					client.EXPECT().MarketOrders(gomock.Any()).
						Return([]*wallex.MarketOrder{
							{
								Price:    "10",
								Quantity: "1",
							},
							{
								Price:    "12",
								Quantity: "1",
							},
							{
								Price:    "13",
								Quantity: "1",
							},
						}, []*wallex.MarketOrder{}, nil)
					amount := decimal.NewFromInt32(2)
					price, err := exchange.PriceFor(model.Pair{}, &amount, nil, sdk.Buy)
					Expect(err).To(BeNil())
					Expect(price.String()).To(Equal(decimal.NewFromInt32(11).String()))
				})
			})
			Context("when multiple order needed and some partial", func() {
				It("should return a unit price for amount", func() {
					client.EXPECT().MarketOrders(gomock.Any()).
						Return([]*wallex.MarketOrder{
							{
								Price:    "10",
								Quantity: "1",
							},
							{
								Price:    "11",
								Quantity: "1",
							},
							{
								Price:    "12",
								Quantity: "2",
							},
						}, []*wallex.MarketOrder{}, nil)
					amount := decimal.NewFromInt32(3)
					price, err := exchange.PriceFor(model.Pair{}, &amount, nil, sdk.Buy)
					Expect(err).To(BeNil())
					Expect(price.String()).To(Equal(decimal.NewFromInt32(11).String()))
				})
			})
		})

	})
	Describe("Get sell price", func() {
		Context("with no amount or funds", func() {
			It("should return a lowest price", func() {
				client.EXPECT().MarketOrders(gomock.Any()).
					Return([]*wallex.MarketOrder{}, []*wallex.MarketOrder{
						{
							Price:    "10",
							Quantity: "1",
						},
						{
							Price:    "12",
							Quantity: "1",
						},
					}, nil)
				price, err := exchange.PriceFor(model.Pair{}, nil, nil, sdk.Sell)
				Expect(err).To(BeNil())
				Expect(price.String()).To(Equal(decimal.NewFromInt32(12).String()))
			})
		})
		Context("for amount", func() {
			Context("when multiple order needed to fill the amount", func() {
				It("should return a unit price for amount", func() {
					client.EXPECT().MarketOrders(gomock.Any()).
						Return([]*wallex.MarketOrder{}, []*wallex.MarketOrder{
							{
								Price:    "10",
								Quantity: "1",
							},
							{
								Price:    "11",
								Quantity: "1",
							},
							{
								Price:    "13",
								Quantity: "1",
							},
						}, nil)
					amount := decimal.NewFromInt32(2)
					price, err := exchange.PriceFor(model.Pair{}, &amount, nil, sdk.Sell)
					Expect(err).To(BeNil())
					Expect(price.String()).To(Equal(decimal.NewFromInt32(12).String()))
				})
			})
			Context("when multiple order needed and some partial", func() {
				It("should return a unit price for amount", func() {
					client.EXPECT().MarketOrders(gomock.Any()).
						Return([]*wallex.MarketOrder{}, []*wallex.MarketOrder{
							{
								Price:    "10",
								Quantity: "2",
							},
							{
								Price:    "11",
								Quantity: "1",
							},
							{
								Price:    "12",
								Quantity: "1",
							},
						}, nil)
					amount := decimal.NewFromInt32(3)
					price, err := exchange.PriceFor(model.Pair{}, &amount, nil, sdk.Sell)
					Expect(err).To(BeNil())
					Expect(price.String()).To(Equal(decimal.NewFromInt32(11).String()))
				})
			})
		})

	})

})
