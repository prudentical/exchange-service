package api

import (
	"log/slog"
	"net/http"
)

type HTTPErrorHandler interface {
	Handle(err error) (int, interface{})
}
type echoErrorHandlerImpl struct {
	logger *slog.Logger
}

func NewHTTPErrorHandler(logger *slog.Logger) HTTPErrorHandler {
	return echoErrorHandlerImpl{logger}
}

func (h echoErrorHandlerImpl) Handle(err error) (int, interface{}) {
	return http.StatusBadRequest, struct {
		Message string
	}{
		Message: err.Error(),
	}
}
