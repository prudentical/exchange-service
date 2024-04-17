package model

type Pair struct {
	BaseEntity
	BaseID     int      `json:"base_id" validate:"required"`
	Base       Currency `json:"base"`
	QuoteID    int      `json:"quote_id" validate:"required"`
	Quote      Currency `json:"quote"`
	ExchangeID int      `json:"exchange_id" validate:"required"`
	Exchange   Exchange `json:"exchange"`
	Symbol     string   `json:"symbol" validate:"required"`
}
