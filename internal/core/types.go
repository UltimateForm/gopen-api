package core

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func (e *ErrorResponse) Render(res http.ResponseWriter, req *http.Request) error {
	render.Status(req, int(e.Code))
	return nil
}

func BadRequest(msg string) *ErrorResponse {
	return &ErrorResponse{Code: 400, Message: msg}
}

func InternalServerError() *ErrorResponse {
	return &ErrorResponse{Code: 500, Message: "Internal Server Error"}
}

func NotFound() *ErrorResponse {
	return &ErrorResponse{Code: 404, Message: "Resource Not Found"}
}
