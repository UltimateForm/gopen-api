package api

import (
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
		core.RespondError(res, req, bindErr)
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
		core.RespondError(res, req, dbError)
		return
	}
	if !userExists {
		core.RespondError(res, req, core.Unauthorized())
		return
	}
	render.Status(req, 204)
}
