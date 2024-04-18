package service

import (
	"exchange-service/internal/message"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	BotId    int             `json:"bot_id"`
	Amount   decimal.Decimal `json:"amount"`
	Price    decimal.Decimal `json:"price"`
	Type     OrderType       `json:"type"`
	DateTime time.Time       `json:"date_time"`
}
type OrderType string

const (
	BuyOrder  OrderType = "Buy"
	SellOrder OrderType = "Sell"
)

type OrderService interface {
	SaveOrder(order Order) error
}

type orderServiceImpl struct {
	msg message.MessageHandler
}

func NewOrderService(msg message.MessageHandler) OrderService {
	return orderServiceImpl{msg: msg}
}

func (s orderServiceImpl) SaveOrder(order Order) error {
	return s.msg.SendMessage(order)
}
