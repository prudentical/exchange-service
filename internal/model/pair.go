package model

type Pair struct {
	BaseEntity
	BaseID     int      `json:"baseId" validate:"required"`
	Base       Currency `json:"base"`
	QuoteID    int      `json:"quoteId" validate:"required"`
	Quote      Currency `json:"quote"`
	ExchangeID int      `json:"exchangeId" validate:"required"`
	Exchange   Exchange `json:"exchange"`
	Symbol     string   `json:"symbol" validate:"required"`
}
