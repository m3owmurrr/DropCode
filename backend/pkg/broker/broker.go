package broker

import (
	"context"
)

type Producer interface {
	Publish(ctx context.Context, exchange, routingKey string, message any) error
}

type Consumer interface {
	Subscribe(ctx context.Context, queue string, handler func([]byte)) error
}

type Broker interface {
	Producer
	Consumer
}
