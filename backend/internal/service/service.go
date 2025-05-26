package service

import (
	"context"
	"io"

	"github.com/m3owmurrr/dropcode/backend/internal/model"
)

type Service interface {
	RunProject(ctx context.Context, req *model.RunRequest) (*model.RunResponse, error)
	SaveProject(ctx context.Context, req *model.SaveRequest) (*model.SaveResponse, error)
	GetProject(ctx context.Context, id string) (io.Reader, error)
}
