package model

const (
	INTERNAL_ERROR string = "internal_error"
	INVALID_JSON          = "invalid_json"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
