package core

import (
	"log"
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

func Unauthorized() *ErrorResponse {
	return &ErrorResponse{Code: 401, Message: "Unauthorized"}
}

func Forbidden() *ErrorResponse {
	return &ErrorResponse{Code: 403, Message: "Forbidden"}
}

func RespondError(res http.ResponseWriter, req *http.Request, err error) {
	switch bindErr := err.(type) {
	case *ErrorResponse:
		render.Render(res, req, bindErr)
	default:
		log.Printf("Non HTTP error occured: %v", err)
		render.Render(res, req, InternalServerError())
	}
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			recovered := recover()
			if recovered != nil {
				log.Printf("UNHANDLED ERROR %v", recovered)
				RespondError(w, r, InternalServerError())
			}
		}()
		next.ServeHTTP(w, r)
	})
}
