package service

import (
	"exchange-service/internal/message"
	"exchange-service/internal/sdk"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	BotId    int             `json:"botId"`
	Amount   decimal.Decimal `json:"amount"`
	Price    decimal.Decimal `json:"price"`
	Type     sdk.TradeType   `json:"type"`
	DateTime time.Time       `json:"datetime"`
}

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
