package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Start() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/login", authHandler)
	log.Println("ROUTER CREATED")
	return router
}
