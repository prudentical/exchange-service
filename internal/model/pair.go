package model

type Pair struct {
	BaseEntity
	BaseID     int64    `json:"baseId" validate:"required"`
	Base       Currency `json:"base"`
	QuoteID    int64    `json:"quoteId" validate:"required"`
	Quote      Currency `json:"quote"`
	ExchangeID int64    `json:"exchangeId" validate:"required"`
	Exchange   Exchange `json:"exchange"`
	Symbol     string   `json:"symbol" validate:"required"`
}
