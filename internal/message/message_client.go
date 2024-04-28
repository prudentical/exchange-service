package message

import (
	"context"
	"exchange-service/internal/configuration"
	"fmt"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueueClient interface {
	Send(exchangeName string, exchangeType string, queue string, payload []byte) error
	Done()
	Shutdown() error
}

type RabbitMQClient struct {
	tag        string
	done       chan error
	logger     *slog.Logger
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewMessageQueueClient(config configuration.Config, logger *slog.Logger) (MessageQueueClient, error) {
	var url = fmt.Sprintf("%s://%s:%s@%s:%d",
		config.Messaging.Protocol,
		config.Messaging.User,
		config.Messaging.Password,
		config.Messaging.Host,
		config.Messaging.Port,
	)

	logger.Debug("Obtaining AMQP connection")
	conn, err := amqp.Dial(url)
	if err != nil {
		return &RabbitMQClient{}, err
	}

	logger.Debug("Opening AMQP channel")
	ch, err := conn.Channel()
	if err != nil {
		return &RabbitMQClient{}, err
	}

	return &RabbitMQClient{logger: logger, connection: conn, channel: ch}, nil
}

func (c *RabbitMQClient) Send(exchangeName string, exchangeType string, queue string, payload []byte) error {
	err := c.setup(exchangeName, exchangeName, queue)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = c.channel.PublishWithContext(ctx,
		exchangeName, // exchange
		queue,        // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        payload,
		})
	return err
}

func (c *RabbitMQClient) Shutdown() error {

	if err := c.channel.Cancel(c.tag, true); err != nil {
		return err
	}

	if err := c.connection.Close(); err != nil {
		return err
	}

	defer c.logger.Info("RabbitMQ client successfully shutdown")

	// wait for client users to exit
	return <-c.done
}

func (c *RabbitMQClient) Done() {
	c.done <- nil
}

func (c *RabbitMQClient) setup(exchangeName string, exchangeType string, queue string) error {
	err := c.channel.ExchangeDeclare(
		exchangeName, // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	_, err = c.channel.QueueDeclare(
		queue, // name of the queue
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	err = c.channel.QueueBind(
		queue,        // name of the queue
		queue,        // bindingKey
		exchangeName, // sourceExchange
		false,        // noWait
		nil,          // arguments
	)
	return err
}
