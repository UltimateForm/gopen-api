package api

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/authrep"
	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/go-chi/render"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (src *LoginResponse) Render(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (src *LoginRequest) Bind(r *http.Request) error {
	if src.Email == "" || src.Password == "" {
		return core.BadRequest("missing either email or password field")
	}
	return nil
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	var reqData LoginRequest
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
	token, tokenCreationErr := core.CreateAuthToken(reqData.Email)
	if tokenCreationErr != nil {
		core.RespondError(res, req, tokenCreationErr)
	}
	tokenRest := LoginResponse{Token: token}
	// NOTE: maybe we dont need render here at all, we could just decode it
	renderErr := render.Render(res, req, &tokenRest)
	if renderErr != nil {
		core.RespondError(res, req, renderErr)
	}
}
