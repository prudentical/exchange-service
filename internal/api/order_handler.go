package api

import (
	"exchange-service/internal/model"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service/exchange"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler interface {
	Handler
	Order(c echo.Context) error
}

type orderHandlerImpl struct {
	order    exchange.OrderService
	exchange exchange.ExchangeService
	factory  sdk.ExchangeSDKFactory
	log      *slog.Logger
}

func NewOrderHandler(order exchange.OrderService, exchange exchange.ExchangeService,
	factory sdk.ExchangeSDKFactory, log *slog.Logger) OrderHandler {
	return orderHandlerImpl{order, exchange, factory, log}
}

func (h orderHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/exchanges/:exchange_id/pairs/:pair_id/order", h.Order)
}

func (h orderHandlerImpl) Order(c echo.Context) error {
	exchangeIdStr := c.Param("exchange_id")
	exchangeId, err := strconv.ParseInt(exchangeIdStr, 10, 64)
	if err != nil || exchangeId <= 0 {
		return InvalidIDError{Id: exchangeIdStr, Type: model.Exchange{}}
	}

	pairIdStr := c.Param("pair_id")
	pairId, err := strconv.ParseInt(pairIdStr, 10, 64)
	if err != nil || pairId <= 0 {
		return InvalidIDError{Id: pairIdStr, Type: model.Pair{}}
	}

	var request exchange.OrderRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		return err
	}

	exchangeModel, err := h.exchange.GetById(pairId)
	if err != nil {
		return err
	}

	exchangeSDK, err := h.factory.Create(exchangeModel)
	if err != nil {
		return err
	}

	err = h.order.Order(exchangeSDK, pairId, request)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
