package broker

import "context"

type Producer interface {
	Publish(ctx context.Context) error
}

type Consumer interface {
	Subscribe(ctx context.Context) error
}

type Broker interface {
	Producer
	Consumer
}
