package service

import (
	"exchange-service/internal/dto"
	"exchange-service/internal/message"
)

type OrderService interface {
	SaveOrder(order dto.OrderDTO) error
}

type orderServiceImpl struct {
	msg message.MessageHandler
}

func NewOrderService(msg message.MessageHandler) OrderService {
	return orderServiceImpl{msg: msg}
}

func (s orderServiceImpl) SaveOrder(order dto.OrderDTO) error {
	return s.msg.SendMessage(order)
}
