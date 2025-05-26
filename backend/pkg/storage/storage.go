package storage

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, bucket, key string, data io.Reader) error
	Get(ctx context.Context, bucket, key string) (io.Reader, error)
	Delete(ctx context.Context, bucket, key string) error
}
