package model

import (
	"encoding/json"
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

type SaveResponse struct {
	ProjectID string `json:"project_id"`
}
