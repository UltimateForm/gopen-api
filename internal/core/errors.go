package core

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"

	"github.com/UltimateForm/gopen-api/internal/repository"
)

type ErrorResponse struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func BadRequest(msg string) ErrorResponse {
	return ErrorResponse{Code: 400, Message: msg}
}

func InternalServerError() ErrorResponse {
	return ErrorResponse{Code: 500, Message: "Internal Server Error"}
}

func NotFound() ErrorResponse {
	return ErrorResponse{Code: 404, Message: "Resource Not Found"}
}

func Unauthorized() ErrorResponse {
	return ErrorResponse{Code: 401, Message: "Unauthorized"}
}

func Forbidden() ErrorResponse {
	return ErrorResponse{Code: 403, Message: "Forbidden"}
}

func RespondKnownError(res http.ResponseWriter, err ErrorResponse) {
	// TODO: use http.Error instead
	res.WriteHeader(int(err.Code))
	json.NewEncoder(res).Encode(err)
}

func mapRepoError(err repository.RepoError) ErrorResponse {
	switch err.Code {
	case repository.EmptyCollectionRepoError:
		return NotFound()
	case repository.ValidationError:
		return BadRequest("Db validation error occured, if you were trying to create something it's likely it already exists")
	case repository.UnknownRepoError:
		fallthrough
	default:
		return InternalServerError()
	}
}

func RespondError(res http.ResponseWriter, req *http.Request, err error) {
	switch bindErr := err.(type) {
	case ErrorResponse:
		RespondKnownError(res, bindErr)
	case repository.RepoError:
		RespondKnownError(res, mapRepoError(bindErr))
	default:
		log.Printf("Non HTTP error of type %T occured: %v", err, err)
		RespondKnownError(res, InternalServerError())

	}
}

func ErrorHandlingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			recovered := recover()
			if recovered != nil {
				const size = 64 << 10
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				log.Printf("UNHANDLED ERROR WHILE SERVING %v: %v\n%s", r.RemoteAddr, recovered, buf)
				RespondError(w, r, InternalServerError())
			}
		}()
		next.ServeHTTP(w, r)
	})
}
