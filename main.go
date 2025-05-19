package main

import (
	"log"
	"net/http"

	"github.com/UltimateForm/gopen-api/cmd/api"
	_ "github.com/UltimateForm/gopen-api/internal/config"
)

func main() {
	log.Println("SERVER START")
	log.Fatal(http.ListenAndServe(":3000", api.Start()))
}
