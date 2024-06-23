package sdk

import (
	"exchange-service/internal/model"
	"net/http"
	"sort"
	"time"

	"github.com/shopspring/decimal"
	"github.com/wallexchange/wallex-go"
)

type wallexAPI struct {
	client   WallexClient
	exchange model.Exchange
}

type WallexClient interface {
	Markets() ([]*wallex.Market, error)
	Currencies() ([]*wallex.Currency, error)
	MarketOrders(symbol string) (ask []*wallex.MarketOrder, bid []*wallex.MarketOrder, _ error)
	Candles(symbol, resolution string, from, to time.Time) ([]*wallex.Candle, error)
}

func NewWallexClient(httpClient http.Client) WallexClient {
	opt := wallex.ClientOptions{
		APIKey:     "",
		HTTPClient: &httpClient,
	}
	return wallex.New(opt)
}

func newWallexAPIClient(exchange model.Exchange, wallex WallexClient) ExchangeAPIClient {
	return wallexAPI{
		client:   wallex,
		exchange: exchange,
	}
}

func (s wallexAPI) Currencies() ([]model.Currency, error) {
	currencies, err := s.client.Currencies()
	if err != nil {
		return make([]model.Currency, 0), err
	}
	return s.toCurrencies(currencies), nil
}

func (wallexAPI) toCurrencies(currencies []*wallex.Currency) []model.Currency {
	result := make([]model.Currency, len(currencies))
	for idx, currency := range currencies {
		result[idx] = model.Currency{
			Name:   currency.Name,
			Symbol: currency.Key,
		}
	}
	return result
}

func (s wallexAPI) Pairs() ([]model.Pair, error) {
	pairs, err := s.client.Markets()
	if err != nil {
		return make([]model.Pair, 0), err
	}
	return s.toPairs(pairs), nil
}

func (s wallexAPI) toPairs(markets []*wallex.Market) []model.Pair {
	result := make([]model.Pair, len(markets))
	for idx, market := range markets {
		result[idx] = model.Pair{
			Symbol: market.Symbol,
			Base: model.Currency{
				Symbol: market.BaseAsset,
			},
			Quote: model.Currency{
				Symbol: market.QuoteAsset,
			},
			ExchangeID: s.exchange.ID,
		}
	}
	return result
}

func (s wallexAPI) PriceFor(pair model.Pair, amount *decimal.Decimal, funds *decimal.Decimal, tradeType TradeType) (decimal.Decimal, error) {
	asks, bids, err := s.client.MarketOrders(pair.Symbol)
	if err != nil {
		return decimal.Decimal{}, err
	}
	var orders []*wallex.MarketOrder
	switch tradeType {
	case Buy:
		s.Sort(asks, true)
		orders = asks
	case Sell:
		s.Sort(bids, false)
		orders = bids
	default:
		return decimal.Decimal{}, InvalidTradeType{string(tradeType)}
	}
	if amount != nil {
		return s.calcPriceForAmount(pair.Symbol, orders, *amount)
	}
	if funds != nil {
		return s.CalcPriceForFunds(pair.Symbol, orders, *funds)
	}
	if len(orders) == 0 {
		return decimal.Decimal{}, InsufficientMarketOrderError{}
	}
	return decimal.NewFromString(string(orders[0].Price))
}

func (wallexAPI) calcPriceForAmount(pairSymbol string, orders []*wallex.MarketOrder, amount decimal.Decimal) (decimal.Decimal, error) {
	price := decimal.NewFromInt32(0)
	sum := decimal.NewFromInt32(0)
	for _, order := range orders {

		prc, err := decimal.NewFromString(string(order.Price))
		if err != nil {
			return decimal.Decimal{}, err
		}

		qty, err := decimal.NewFromString(string(order.Quantity))
		if err != nil {
			return decimal.Decimal{}, err
		}

		if amount.Sub(sum).Cmp(qty) <= 0 {
			price = price.Add(amount.Sub(sum).Mul(prc))
			sum = sum.Add(amount.Sub(sum))
		} else {
			price = price.Add(qty.Mul(prc))
			sum = sum.Add(qty)
		}

		if amount.Cmp(sum) == 0 {
			break
		}
	}
	if sum.Cmp(amount) < 0 {
		return decimal.Decimal{}, InsufficientMarketOrderError{PairSymbol: pairSymbol, Asked: amount, Available: sum}
	}
	return price.Div(amount), nil
}

func (wallexAPI) CalcPriceForFunds(pairSymbol string, orders []*wallex.MarketOrder, funds decimal.Decimal) (decimal.Decimal, error) {
	amount := decimal.NewFromInt32(0)
	sum := decimal.NewFromInt32(0)
	for _, order := range orders {

		prc, err := decimal.NewFromString(string(order.Price))
		if err != nil {
			return decimal.Decimal{}, err
		}

		qty, err := decimal.NewFromString(string(order.Quantity))
		if err != nil {
			return decimal.Decimal{}, err
		}
		if funds.Cmp(sum.Add(prc.Mul(qty))) >= 0 {
			amount = amount.Add(qty)
			sum = sum.Add(prc.Mul(qty))
		} else {
			percentage := funds.Sub(sum).Div(prc.Mul(qty))
			amount = amount.Add(qty.Mul(percentage))
			sum = sum.Add(prc.Mul(qty.Mul(percentage)))
		}

		if sum.Cmp(funds) == 0 {
			break
		}
	}
	if sum.Cmp(funds) < 0 {
		return decimal.Decimal{}, InsufficientMarketOrderError{PairSymbol: pairSymbol, Asked: funds, Available: sum}
	}
	return funds.Div(amount), nil
}

func (wallexAPI) Sort(orders []*wallex.MarketOrder, ascending bool) {
	sort.Slice(orders, func(i, j int) bool {
		if ascending {
			return orders[i].Price < orders[j].Price
		}
		return orders[i].Price > orders[j].Price
	})
}

func (s wallexAPI) HistoricPrice(pair model.Pair, dateTime time.Time) (decimal.Decimal, error) {
	from := dateTime.Add(-15 * time.Minute)
	candles, err := s.client.Candles(pair.Symbol, "1", from, dateTime)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal.NewFromString(string(candles[len(candles)-1].Close))
}

func (s wallexAPI) GetExchange() model.Exchange {
	return s.exchange
}
