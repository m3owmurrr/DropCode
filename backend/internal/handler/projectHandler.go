package v1

import (
	"encoding/json"
	"net/http"

	"github.com/m3owmurrr/dropcode/backend/internal/model"
	"github.com/m3owmurrr/dropcode/backend/internal/service"
)

type ProjectHandler struct {
	serv service.Service
}

func NewProjectHandler(serv service.Service) *ProjectHandler {
	return &ProjectHandler{
		serv: serv,
	}
}

func (ph *ProjectHandler) RunProjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	data := &model.RunRequest{
		SessionID: r.Header.Get("X-Session-ID"),
		Language:  r.Header.Get("X-Language"),
		Project:   r.Body,
	}

	runResp, err := ph.serv.RunProject(ctx, data)
	if err != nil {
		// TODO: после реализации сервисного слоя
		return
	}

	resp, err := json.Marshal(runResp)
	if err != nil {
		msg := "failed to encode response"
		writeError(w, http.StatusInternalServerError, model.INTERNAL_ERROR, msg)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(resp)
}

func writeError(w http.ResponseWriter, statusCode int, errCode, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.ErrorResponse{
		Error:   errCode,
		Message: msg,
	})
}
