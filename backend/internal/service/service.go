package service

import (
	"context"
	"io"

	"github.com/m3owmurrr/dropcode/backend/internal/model"
)

type Service interface {
	RunProject(ctx context.Context, data io.Reader) (*model.RunResponse, error)
	SaveProject(ctx context.Context, data io.Reader) (*model.SaveResponse, error)
	GetProject(ctx context.Context, id string) (io.Reader, error)
}
