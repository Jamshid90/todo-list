package models

import (
	"net/http"

	"github.com/go-chi/render"
)

type ResponseError struct {
	HTTPStatusCode int    `json:"status"`
	Code           string `json:"code"`
	Message        string `json:"message"`
}

func NewResponseError(httpStatusCode int, code string, message string) *ResponseError {
	return &ResponseError{
		HTTPStatusCode: httpStatusCode,
		Code:           code,
		Message:        message,
	}
}

func (e *ResponseError) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}
