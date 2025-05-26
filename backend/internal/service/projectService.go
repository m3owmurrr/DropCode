package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (ps *ProjectService) RunProject(ctx context.Context, req *model.RunRequest) (*model.RunResponse, error) {
	runID := fmt.Sprintf("run-%v", uuid.New().String())

	if err := ps.stor.Put(ctx, config.Cfg.S3.RunBucket, runID, bytes.NewReader(req.Project)); err != nil {
		return nil, err
	}

	routingKey := "run." + req.Language
	message := &model.RunMessage{RunId: runID}

	if err := ps.brok.Publish(ctx, "runs", routingKey, message); err != nil {
		return nil, err
	}

	resp := &model.RunResponse{RunID: runID}

	return resp, nil
}

func (ps *ProjectService) SaveProject(ctx context.Context, req *model.SaveRequest) (*model.SaveResponse, error) {
	projectID := fmt.Sprintf("project-%v", uuid.New().String())

	saveProject := req.ToSaveProject()

	data, err := json.Marshal(saveProject)
	if err != nil {
		return nil, err
	}

	if err := ps.stor.Put(ctx, config.Cfg.S3.SaveBucket, projectID, bytes.NewReader(data)); err != nil {
		return nil, err
	}

	resp := &model.SaveResponse{ProjectID: projectID}

	return resp, nil
}

func (ps *ProjectService) GetProject(ctx context.Context, id string) (io.Reader, error) {
	return nil, nil
}
