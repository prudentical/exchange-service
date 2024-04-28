package message

import (
	"encoding/json"
	"exchange-service/internal/configuration"
	"log/slog"
)

type MessageHandler interface {
	SendMessage(payload interface{}) error
}

type RabbitMQService struct {
	logger       *slog.Logger
	client       MessageQueueClient
	exchangeName string
	exchangeType string
	queue        string
}

func NewMessageHandler(logger *slog.Logger, client MessageQueueClient, config configuration.Config) MessageHandler {
	return &RabbitMQService{
		logger:       logger,
		client:       client,
		exchangeName: config.Messaging.Order.Exchange.Name,
		exchangeType: config.Messaging.Order.Exchange.Type,
		queue:        config.Messaging.Order.Queue,
	}
}

func (s *RabbitMQService) SendMessage(payload interface{}) error {
	str, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = s.client.Send(s.exchangeName, s.exchangeType, s.queue, str)
	if err != nil {
		return err
	}
	return nil
}
