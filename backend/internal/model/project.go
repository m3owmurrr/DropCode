package model

import "io"

type RunRequest struct {
	SessionID string
	Language  string
	Project   io.Reader
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
