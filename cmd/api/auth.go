package api

import (
	"log"
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/authrep"
	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/go-chi/render"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (src *LoginDto) Bind(r *http.Request) error {
	if src.Email == "" || src.Password == "" {
		return core.BadRequest("missing either email or password field")
	}
	return nil
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	var reqData LoginDto
	bindErr := render.Bind(req, &reqData)
	if bindErr != nil {
		switch bindErr := bindErr.(type) {
		case *core.ErrorResponse:
			render.Render(res, req, bindErr)
		default:
			render.Render(res, req, core.InternalServerError())
		}
		return
	}
	userExists, dbError := authrep.UserExistsAtomically(
		req.Context(),
		authrep.User{
			Email:    reqData.Email,
			Password: reqData.Password,
		},
	)
	if dbError != nil {
		log.Println(dbError.Error())
		render.Render(res, req, core.InternalServerError())
		return
	}
	if !userExists {
		render.Render(res, req, core.NotFound())
		return
	}
	render.Status(req, 204)
}
