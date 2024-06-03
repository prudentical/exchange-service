package dto

import (
	"exchange-service/internal/sdk"
	"time"

	"github.com/shopspring/decimal"
)

type OrderDTO struct {
	BotId        int64           `json:"botId" validate:"required"`
	InternalId   string          `json:"internalId"`
	Amount       decimal.Decimal `json:"amount" validate:"required"`
	FilledAmount decimal.Decimal `json:"filledAmount"`
	Price        decimal.Decimal `json:"price"`
	Type         sdk.TradeType   `json:"type" validate:"required"`
	Status       OrderStatus     `json:"status"`
	DateTime     *time.Time      `json:"datetime"`
	IsVirtual    bool            `json:"virtual"`
}

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Partial   OrderStatus = "Partial"
	Fulfilled OrderStatus = "Fulfilled"
	Canceled  OrderStatus = "Canceled"
)
