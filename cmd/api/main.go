package api

import (
	"log"
	"net/http"

	"github.com/UltimateForm/gopen-api/cmd/api/charactersapi"
	"github.com/UltimateForm/gopen-api/cmd/api/loginapi"
	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(core.ErrorHandlingMiddleware)
	router.Post("/login", loginapi.HandlePostLogin)
	router.Route("/characters", func(r chi.Router) {
		r.Use(core.AuthMiddleware)
		r.Get("/", charactersapi.HandleGetCharacters)
	})
	log.Println("ROUTER CREATED")
	return router
}
