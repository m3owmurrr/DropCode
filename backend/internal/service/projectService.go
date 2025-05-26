package service

import (
	"bytes"
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/m3owmurrr/dropcode/backend/internal/config"
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

func (ps *ProjectService) RunProject(ctx context.Context, data *model.RunRequest) (*model.RunResponse, error) {
	runID := "run-" + uuid.New().String()

	if err := ps.stor.Put(ctx, config.Cfg.S3.RunBucket, runID, bytes.NewReader(data.Project)); err != nil {
		return nil, err
	}

	routingKey := "run." + data.Language
	message := &model.RunMessage{RunId: runID}

	if err := ps.brok.Publish(ctx, "runs", routingKey, message); err != nil {
		return nil, err
	}

	resp := &model.RunResponse{RunID: runID}

	return resp, nil
}

func (ps *ProjectService) SaveProject(ctx context.Context, data io.Reader) (*model.SaveResponse, error) {
	return nil, nil
}

func (ps *ProjectService) GetProject(ctx context.Context, id string) (io.Reader, error) {
	return nil, nil
}
