package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/m3owmurrr/dropcode/backend/internal/config"
	"github.com/m3owmurrr/dropcode/backend/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitBroker struct {
	conn  *amqp.Connection
	pubCh *amqp.Channel
	subCh *amqp.Channel
}

func NewRabbitBroker() (*RabbitBroker, error) {
	conn, err := amqp.Dial(config.Cfg.Rabbit.URL)
	if err != nil {
		log.Printf("can't connect to rabbitmq: %v", err)
		return nil, err
	}

	initCh, err := conn.Channel()
	if err != nil {
		log.Printf("can't create publish channel: %v", err)
		return nil, err
	}
	defer initCh.Close()

	if err := mustSetup(initCh); err != nil {
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

func (rb *RabbitBroker) Publish(ctx context.Context, exchange, routingKey string, message *model.RunMessage) error {
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

func setup(ch *amqp.Channel) error {
	if err := ch.ExchangeDeclare("runs", "topic", true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.ExchangeDeclare("results", "direct", true, false, false, false, nil); err != nil {
		return err
	}

	if _, err := ch.QueueDeclare("runs-go", true, false, false, false, nil); err != nil {
		return err
	}

	if _, err := ch.QueueDeclare("runs-python", true, false, false, false, nil); err != nil {
		return err
	}

	if _, err := ch.QueueDeclare("runs-javascript", true, false, false, false, nil); err != nil {
		return err
	}

	if _, err := ch.QueueDeclare("results", true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.QueueBind("runs-go", "run.go", "runs", false, nil); err != nil {
		return err
	}

	if err := ch.QueueBind("runs-python", "run.python", "runs", false, nil); err != nil {
		return err
	}

	if err := ch.QueueBind("runs-javascript", "run.javascript", "runs", false, nil); err != nil {
		return err
	}

	if err := ch.QueueBind("results", "result", "results", false, nil); err != nil {
		return err
	}

	return nil
}

func mustSetup(ch *amqp.Channel) error {
	backoff := 1 * time.Second
	const maxTries = 10

	for i := 1; i <= maxTries; i++ {
		if err := setup(ch); err == nil {
			return nil
		} else {
			log.Printf("broker topology setup failed: %v, try: %v", err, i)
			time.Sleep(backoff)
		}
	}

	log.Printf("can't initialize broker topology")
	return fmt.Errorf("can't initialize topology")
}
