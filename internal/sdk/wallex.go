package sdk

import (
	"errors"
	"exchange-service/internal/model"
	"net/http"
	"sort"
	"time"

	"github.com/shopspring/decimal"
	"github.com/wallexchange/wallex-go"
)

type wallexSDK struct {
	client   wallex.Client
	exchange model.Exchange
}

func newWallexSDK(exchange model.Exchange) ExchangeSDK {
	opt := wallex.ClientOptions{
		APIKey: "",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return wallexSDK{
		client:   *wallex.New(opt),
		exchange: exchange,
	}
}

func (s wallexSDK) Currencies() ([]model.Currency, error) {
	currencies, err := s.client.Currencies()
	if err != nil {
		return make([]model.Currency, 0), err
	}
	return s.toCurrencies(currencies), nil
}

func (wallexSDK) toCurrencies(currencies []*wallex.Currency) []model.Currency {
	result := make([]model.Currency, len(currencies))
	for idx, currency := range currencies {
		result[idx] = model.Currency{
			Name:   currency.Name,
			Symbol: currency.Key,
		}
	}
	return result
}

func (s wallexSDK) Pairs() ([]model.Pair, error) {
	pairs, err := s.client.Markets()
	if err != nil {
		return make([]model.Pair, 0), err
	}
	return s.toPairs(pairs), nil
}

func (s wallexSDK) toPairs(markets []*wallex.Market) []model.Pair {
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

func (s wallexSDK) PriceFor(pair model.Pair, amount decimal.Decimal, tradeType TradeType) (decimal.Decimal, error) {
	asks, bids, err := s.client.MarketOrders(pair.Symbol)
	if err != nil {
		return decimal.Decimal{}, err
	}
	switch tradeType {
	case Buy:
		s.Sort(asks, true)
		return s.CalcPriceForAmount(asks, amount)
	case Sell:
		s.Sort(asks, false)
		return s.CalcPriceForAmount(bids, amount)
	default:
		return decimal.Decimal{}, InvalidTradeType{string(tradeType)}
	}
}
func (wallexSDK) CalcPriceForAmount(orders []*wallex.MarketOrder, amount decimal.Decimal) (decimal.Decimal, error) {
	price := decimal.NewFromInt32(0)
	sum := decimal.NewFromInt32(0)
	for _, ask := range orders {

		prc, err := decimal.NewFromString(string(ask.Price))
		if err != nil {
			return decimal.Decimal{}, err
		}

		qty, err := decimal.NewFromString(string(ask.Quantity))
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
		// Todo: return a specific error type
		return decimal.Decimal{}, errors.New("insufficient market orders")
	}
	return price.Div(amount), nil
}

func (wallexSDK) Sort(orders []*wallex.MarketOrder, ascending bool) {
	sort.Slice(orders, func(i, j int) bool {
		if ascending {
			return orders[i].Price < orders[j].Price
		}
		return orders[i].Price > orders[j].Price
	})
}
func (s wallexSDK) HistoricPrice(pair model.Pair, dateTime time.Time) (decimal.Decimal, error) {
	from := dateTime.Add(-15 * time.Minute)
	candles, err := s.client.Candles(pair.Symbol, "1", from, dateTime)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal.NewFromString(string(candles[len(candles)-1].Close))
}

func (s wallexSDK) GetExchange() model.Exchange {
	return s.exchange
}
