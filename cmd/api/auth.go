package api

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/authrep"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (src *LoginRequest) Validate() error {
	if src.Email == "" || src.Password == "" {
		return core.BadRequest("missing either email or password field")
	}
	return nil
}

func authHandler(res http.ResponseWriter, req *http.Request) {
	var reqData LoginRequest
	bodyErr := core.ParseBody(req, &reqData)
	if bodyErr != nil {
		core.RespondError(res, req, bodyErr)
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
	core.RespondOk(res, tokenRest)
}
