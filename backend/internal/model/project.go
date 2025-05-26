package model

import (
	"encoding/json"
	"time"
)

type RunRequest struct {
	SessionID string          `json:"session_id"`
	Language  string          `json:"language"`
	Project   json.RawMessage `json:"project"`
}

type RunResponse struct {
	RunID string `json:"run_id"`
}

type RunMessage struct {
	RunId string `jsom:"run_id"`
}

type SaveRequest struct {
	TimeToLive time.Duration   `json:"time_to_live,omitempty"`
	Language   string          `json:"language"`
	Project    json.RawMessage `json:"project"`
}

type SaveResponse struct {
	ProjectID string `json:"project_id"`
}

type SaveProject struct {
	Language string          `json:"language"`
	Project  json.RawMessage `json:"project"`
}

func (sr *SaveRequest) ToSaveProject() *SaveProject {
	return &SaveProject{
		Language: sr.Language,
		Project:  sr.Project,
	}
}
