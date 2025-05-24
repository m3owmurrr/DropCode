package service

import (
	"context"
	"io"

	"github.com/m3owmurrr/dropcode/backend/internal/model"
	"github.com/m3owmurrr/dropcode/backend/internal/repository"
	"github.com/m3owmurrr/dropcode/backend/pkg/broker"
	"github.com/m3owmurrr/dropcode/backend/pkg/storage"
)

type ProjectService struct {
	repo repository.Repository
	stor storage.Storage
	brok broker.Broker
}

func NewProjectService(repo repository.Repository, stor storage.Storage, brok broker.Broker) *ProjectService {
	return &ProjectService{
		repo: repo,
		stor: stor,
		brok: brok,
	}
}

func (ps *ProjectService) RunProject(ctx context.Context, data io.Reader) (*model.RunResponse, error) {
	return nil, nil
}

func (ps *ProjectService) SaveProject(ctx context.Context, data io.Reader) (*model.SaveResponse, error) {
	return nil, nil
}

func (ps *ProjectService) GetProject(ctx context.Context, id string) (io.Reader, error) {
	return nil, nil
}
