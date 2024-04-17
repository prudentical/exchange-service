package api

import (
	"exchange-service/internal/model"
	"exchange-service/internal/service/currency"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CurrencyHandler interface {
	Handler
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type currencyHandlerImpl struct {
	currency currency.CurrencyService
	log      *slog.Logger
}

func NewCurrencyHandler(currency currency.CurrencyService, log *slog.Logger) CurrencyHandler {
	return currencyHandlerImpl{currency, log}
}

func (h currencyHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/currencies", h.GetAll)
	e.POST("/currencies", h.Create)
	e.GET("/currencies/:id", h.GetById)
	e.PUT("/currencies/:id", h.Update)
	e.DELETE("/currencies/:id", h.Delete)
}

func (h currencyHandlerImpl) GetAll(c echo.Context) error {
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
	h.log.Debug("Get all currencies", "page", page, "size", size)
	result, err := h.currency.GetAll(page, size)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h currencyHandlerImpl) GetById(c echo.Context) error {
	idStr := c.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	pair, err := h.currency.GetById(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pair)
}

func (h currencyHandlerImpl) Create(c echo.Context) error {
	var currency *model.Currency = new(model.Currency)
	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: add pair id to location header
	created, err := h.currency.Create(*currency)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, created)
}

func (h currencyHandlerImpl) Update(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	var currency *model.Currency = new(model.Currency)
	if err := c.Bind(currency); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.currency.Update(id, *currency)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, updated)
}

func (h currencyHandlerImpl) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	err = h.currency.Delete(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
