package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	HandleRoutes(e *echo.Echo)
}

type HelloHandler struct {
}

func NewHelloHandler() Handler {
	return HelloHandler{}
}

func (HelloHandler) HandleRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
