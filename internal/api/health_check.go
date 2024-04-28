package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheck interface {
	Check(c echo.Context) error
	HandleRoutes(e *echo.Echo)
}

func NewHealthCheck() HealthCheck {
	return healthCheckImpl{}
}

type healthCheckImpl struct{}

func (h healthCheckImpl) HandleRoutes(e *echo.Echo) {
	e.GET("/health", h.Check)
}

func (healthCheckImpl) Check(c echo.Context) error {
	status := struct {
		Status string `json:"status"`
	}{
		Status: "UP",
	}
	return c.JSON(http.StatusOK, status)
}
