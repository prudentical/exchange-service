package sdk

type InsufficientMarketOrderError struct{}

func (InsufficientMarketOrderError) Error() string {
	return "Insufficient market orders"
}
