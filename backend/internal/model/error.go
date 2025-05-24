package model

const (
	INTERNAL_ERROR string = "internal_error"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
