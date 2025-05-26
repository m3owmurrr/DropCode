package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitBroker struct {
	conn  *amqp.Connection
	pubCh *amqp.Channel
	subCh *amqp.Channel
}

func NewRabbitBroker(cfg RabbitConfig) (*RabbitBroker, error) {
	conn, err := amqp.Dial(cfg.URL)
	if err != nil {
		log.Printf("can't connect to rabbitmq: %v", err)
		return nil, err
	}

	pubCh, err := conn.Channel()
	if err != nil {
		log.Printf("can't create publish channel: %v", err)
		return nil, err
	}

	subCh, err := conn.Channel()
	if err != nil {
		log.Printf("can't create subscribe  channel: %v", err)
		return nil, err
	}

	broker := &RabbitBroker{
		conn:  conn,
		pubCh: pubCh,
		subCh: subCh,
	}

	return broker, nil
}

func (rb *RabbitBroker) Publish(ctx context.Context, exchange, routingKey string, message any) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = rb.pubCh.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	return err
}

func (rb *RabbitBroker) Subscribe(ctx context.Context, queue string, handler func([]byte)) error {
	msgs, err := rb.subCh.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("consumer context done, stopping...")
				return
			case d, ok := <-msgs:
				if !ok {
					log.Println("consumer channel closed")
					return
				}
				handler(d.Body)
			}
		}
	}()

	return nil
}
