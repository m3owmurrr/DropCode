package storage

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, bucket string, id string, data io.Reader) error
	Get(ctx context.Context, bucket string, id string) (io.Reader, error)
}
