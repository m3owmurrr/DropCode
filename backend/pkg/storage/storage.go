package storage

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, bucket string, key string, data io.Reader) error
	Get(ctx context.Context, bucket string, key string) (io.Reader, error)
}
