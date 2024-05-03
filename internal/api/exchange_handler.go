package api

import (
	"exchange-service/internal/model"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service/exchange"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ExchangeHandler interface {
	Handler
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Update(c echo.Context) error
}

type exchangeHandlerImpl struct {
	service    exchange.ExchangeService
	sdkFactory sdk.ExchangeSDKFactory
}

func NewExchangeHandler(service exchange.ExchangeService, sdkFactory sdk.ExchangeSDKFactory) ExchangeHandler {
	return exchangeHandlerImpl{service, sdkFactory}
}

func (h exchangeHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/exchanges", h.GetAll)
	e.GET("/exchanges/:id", h.GetById)
	e.PUT("/exchanges/:id", h.Update)
}

func (h exchangeHandlerImpl) GetAll(c echo.Context) error {
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	sizeStr := c.QueryParam("size")
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		size = 20
	}

	result, err := h.service.GetAllWithPage(page, size)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h exchangeHandlerImpl) GetById(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	exchange, err := h.service.GetById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, exchange)
}

func (h exchangeHandlerImpl) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	var exchangeModel model.Exchange
	if err := c.Bind(&exchangeModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(exchangeModel); err != nil {
		return err
	}
	exchangeModel, err = h.service.Update(id, exchangeModel)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, exchangeModel)
}
