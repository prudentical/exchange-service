package sdk

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type InsufficientMarketOrderError struct {
	Asked      decimal.Decimal
	Available  decimal.Decimal
	PairSymbol string
}

func (i InsufficientMarketOrderError) Error() string {
	return fmt.Sprintf("Insufficient market orders[pair=%s, asked=%s, available=%s]", i.PairSymbol, i.Asked, i.Available)
}
