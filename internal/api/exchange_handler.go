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
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	PriceFor(c echo.Context) error
	HistoricPrice(c echo.Context) error
	Buy(c echo.Context) error
	Sell(c echo.Context) error
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
	e.POST("/exchanges", h.Create)
	e.GET("/exchanges/:id", h.GetById)
	e.PUT("/exchanges/:id", h.Update)
	e.DELETE("/exchanges/:id", h.Delete)
	e.GET("/exchanges/:id/price", h.PriceFor)
	e.GET("/exchanges/:id/historic-price", h.HistoricPrice)
	e.POST("/exchanges/:id/buy", h.Buy)
	e.POST("/exchanges/:id/sell", h.Sell)
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
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	exchange, err := h.service.GetById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, exchange)
}

func (h exchangeHandlerImpl) Create(c echo.Context) error {
	var exchangeModel *model.Exchange = new(model.Exchange)
	if err := c.Bind(exchangeModel); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(exchangeModel); err != nil {
		return err
	}
	exchanges, err := h.service.Create(*exchangeModel)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, exchanges)
}

func (h exchangeHandlerImpl) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
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

func (h exchangeHandlerImpl) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	err = h.service.Delete(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (h exchangeHandlerImpl) PriceFor(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	var request exchange.PriceCheckRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	exchangeModel, err := h.service.GetById(id)
	if err != nil {
		return err
	}
	exchangeSDK, err := h.sdkFactory.Create(exchangeModel)
	if err != nil {
		return err
	}
	price, err := h.service.PriceFor(exchangeSDK, request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, exchange.PriceCheckResponse{Price: price})
}

func (h exchangeHandlerImpl) HistoricPrice(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	var request exchange.PriceCheckRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	exchangeModel, err := h.service.GetById(id)
	if err != nil {
		return err
	}
	exchangeSDK, err := h.sdkFactory.Create(exchangeModel)
	if err != nil {
		return err
	}
	price, err := h.service.HistoricPrice(exchangeSDK, request)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, exchange.PriceCheckResponse{Price: price})
}

func (h exchangeHandlerImpl) Buy(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	var request exchange.OrderRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	exchangeModel, err := h.service.GetById(id)
	if err != nil {
		return err
	}
	exchangeSDK, err := h.sdkFactory.Create(exchangeModel)
	if err != nil {
		return err
	}
	err = h.service.Buy(exchangeSDK, request)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func (h exchangeHandlerImpl) Sell(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return InvalidIDError{Id: idStr, Type: model.Exchange{}}
	}
	var request exchange.OrderRequest
	if err := c.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(request); err != nil {
		return err
	}
	exchangeModel, err := h.service.GetById(id)
	if err != nil {
		return err
	}
	exchangeSDK, err := h.sdkFactory.Create(exchangeModel)
	if err != nil {
		return err
	}
	err = h.service.Buy(exchangeSDK, request)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}
