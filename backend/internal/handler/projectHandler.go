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
	var runReq model.RunRequest

	err := json.NewDecoder(r.Body).Decode(&runReq)
	if err != nil {
		msg := "invalid JSON in request body"
		writeError(w, http.StatusBadRequest, model.INVALID_JSON, msg)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	runResp, err := ph.serv.RunProject(ctx, &runReq)
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

func (ph *ProjectHandler) SaveProjectHandler(w http.ResponseWriter, r *http.Request) {
	var saveReq model.SaveRequest

	err := json.NewDecoder(r.Body).Decode(&saveReq)
	if err != nil {
		msg := "invalid JSON in request body"
		writeError(w, http.StatusBadRequest, model.INVALID_JSON, msg)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()
	runResp, err := ph.serv.SaveProject(ctx, &saveReq)
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
