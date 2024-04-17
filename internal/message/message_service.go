package message

import (
	"encoding/json"
	"exchange-service/configuration"
	"log/slog"
)

type MessageService interface {
	SendMessage(payload interface{}) error
}

type RabbitMQService struct {
	logger *slog.Logger
	client MessageQueueClient
	queue  string
}

func NewMessageService(logger *slog.Logger, client MessageQueueClient, config configuration.Config) MessageService {
	return &RabbitMQService{
		logger: logger,
		client: client,
		queue:  config.Messaging.Queue.Name,
	}
}

func (s *RabbitMQService) SendMessage(payload interface{}) error {
	str, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = s.client.Send(s.queue, str)
	if err != nil {
		return err
	}
	return nil
}
