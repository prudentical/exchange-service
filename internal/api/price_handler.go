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

type PriceHandler interface {
	Handler
	GetPrice(c echo.Context) error
}

type priceHandlerImpl struct {
	price    exchange.PriceService
	exchange exchange.ExchangeService
	factory  sdk.ExchangeAPIClientFactory
	log      *slog.Logger
}

func NewPriceHandler(price exchange.PriceService, exchange exchange.ExchangeService,
	factory sdk.ExchangeAPIClientFactory, log *slog.Logger) PriceHandler {
	return priceHandlerImpl{price, exchange, factory, log}
}

func (h priceHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/exchanges/:exchange_id/pairs/:pair_id/price", h.GetPrice)
}

func (h priceHandlerImpl) GetPrice(c echo.Context) error {
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

	var request exchange.PriceCheckRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return err
	}

	exchangeModel, err := h.exchange.GetById(exchangeId)
	if err != nil {
		return err
	}

	exchangeSDK, err := h.factory.Create(exchangeModel)
	if err != nil {
		return err
	}

	price, err := h.price.GetPrice(exchangeSDK, pairId, request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, exchange.PriceCheckResponse{Price: price})
}
