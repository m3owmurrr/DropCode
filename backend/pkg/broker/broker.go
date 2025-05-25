package broker

import (
	"context"

	"github.com/m3owmurrr/dropcode/backend/internal/model"
)

type Producer interface {
	Publish(ctx context.Context, exchange, routingKey string, message *model.RunMessage) error
}

type Consumer interface {
	Subscribe(ctx context.Context, queue string, handler func([]byte)) error
}

type Broker interface {
	Producer
	Consumer
}
