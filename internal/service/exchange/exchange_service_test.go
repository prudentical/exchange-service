package exchange_test

import (
	"exchange-service/internal/model"
	"exchange-service/internal/persistence"
	mock_persistence "exchange-service/internal/persistence/mock"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"
	"exchange-service/internal/service/exchange"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Exchange manage", Label("exchange"), func() {

	var manager exchange.ExchangeService
	var ctrl *gomock.Controller
	var dao *mock_persistence.MockExchangeDAO

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		dao = mock_persistence.NewMockExchangeDAO(ctrl)
		manager = exchange.NewExchangeService(dao, sdk.NewExchangeSDKFactory())
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Get exchange by id", func() {
		Context("non-existing exchange", func() {
			It("should return not-found error", func() {
				dao.EXPECT().Get(gomock.Any()).Return(model.Exchange{}, persistence.RecordNotFoundError{})
				_, err := manager.GetById(1)
				Expect(err).To(MatchError(service.NotFoundError{Type: model.Exchange{}, Id: 1}))
			})
			Context("existing exchange", func() {
				It("should return the exchange", func() {
					dao.EXPECT().Get(gomock.Any()).Return(model.Exchange{}, nil)
					_, err := manager.GetById(1)
					Expect(err).To(BeNil())
				})
			})
		})
	})
	Describe("Update exchange", func() {
		Context("non-existing exchange", func() {
			It("should return not-found error", func() {
				exchange := model.Exchange{
					Name: "Wallex",
				}
				dao.EXPECT().Get(gomock.Any()).Return(exchange, persistence.RecordNotFoundError{})
				_, err := manager.Update(1, exchange)
				Expect(err).To(MatchError(service.NotFoundError{Type: model.Exchange{}, Id: 1}))
			})
		})
		Context("existing exchange", func() {
			It("should update the exchange", func() {
				exchange := model.Exchange{
					Name: "Wallex",
				}
				dao.EXPECT().Get(gomock.Any()).Return(exchange, nil)
				dao.EXPECT().Update(gomock.Any()).Return(exchange, nil)
				_, err := manager.Update(1, exchange)
				Expect(err).To(BeNil())
			})
		})
	})

})
