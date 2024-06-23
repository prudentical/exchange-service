package api

import (
	"exchange-service/internal/model"
	"exchange-service/internal/service/exchange"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PairHandler interface {
	Handler
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type pairHandlerImpl struct {
	pairs exchange.PairService
	log   *slog.Logger
}

func NewPairHandler(pairs exchange.PairService, log *slog.Logger) PairHandler {
	return pairHandlerImpl{pairs, log}
}

func (h pairHandlerImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/exchanges/:exchange_id/pairs", h.GetAll)
	e.POST("/exchanges/:exchange_id/pairs", h.Create)
	e.GET("/exchanges/:exchange_id/pairs/:id", h.GetById)
	e.PUT("/exchanges/:exchange_id/pairs/:id", h.Update)
	e.DELETE("/exchanges/:exchange_id/pairs/:id", h.Delete)
}

func (h pairHandlerImpl) GetAll(c echo.Context) error {
	exchangeIdStr := c.Param("exchange_id")
	exchangeId, err := strconv.ParseInt(exchangeIdStr, 10, 64)
	if err != nil {
		return err
	}
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
	result, err := h.pairs.GetAll(exchangeId, page, size)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h pairHandlerImpl) GetById(c echo.Context) error {
	exchangeIdStr := c.Param("exchange_id")
	exchangeId, err := strconv.ParseInt(exchangeIdStr, 10, 64)
	if err != nil {
		return err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	pair, err := h.pairs.GetById(exchangeId, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, pair)
}

func (h pairHandlerImpl) Create(c echo.Context) error {
	exchangeIdStr := c.Param("exchange_id")
	exchangeId, err := strconv.ParseInt(exchangeIdStr, 10, 64)
	if err != nil {
		return err
	}

	var pair *model.Pair = new(model.Pair)
	if err := c.Bind(pair); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO: add pair id to location header
	created, err := h.pairs.Create(exchangeId, *pair)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, created)
}

func (h pairHandlerImpl) Update(c echo.Context) error {
	exchangeIdStr := c.Param("exchange_id")
	exchangeId, err := strconv.ParseInt(exchangeIdStr, 10, 64)
	if err != nil {
		return err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}

	var pair *model.Pair = new(model.Pair)
	if err := c.Bind(pair); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updated, err := h.pairs.Update(exchangeId, id, *pair)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, updated)
}

func (h pairHandlerImpl) Delete(c echo.Context) error {
	exchangeIdStr := c.Param("exchange_id")
	exchangeId, err := strconv.ParseInt(exchangeIdStr, 10, 64)
	if err != nil {
		return err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return err
	}
	err = h.pairs.Delete(exchangeId, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
