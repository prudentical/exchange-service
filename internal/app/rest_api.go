package app

import (
	"exchange-service/internal/api"
	"exchange-service/internal/util"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

type RESTApp interface {
	setup() error
	server() *echo.Echo
}

type EchoApp struct {
	e            *echo.Echo
	validator    util.Validator
	handlers     []api.Handler
	errorHandler api.HTTPErrorHandler
	log          *slog.Logger
}

func NewRestApp(handlers []api.Handler, log *slog.Logger, e *echo.Echo, errorHandler api.HTTPErrorHandler, validator util.Validator) RESTApp {
	log.Info("Creating REST App")
	return EchoApp{e, validator, handlers, errorHandler, log}
}

func (app EchoApp) setup() error {
	app.e.Use(slogecho.New(app.log))
	app.e.Use(middleware.Recover())
	app.e.Validator = app.validator
	for _, handler := range app.handlers {
		handler.HandleRoutes(app.server())
	}
	app.e.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}
		code, response := app.errorHandler.Handle(err)
		c.JSON(code, response)
	}
	return nil
}

func (app EchoApp) server() *echo.Echo {
	return app.e
}

func ProvideEcho() *echo.Echo {
	return echo.New()
}
