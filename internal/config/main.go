package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	log.Println("CONFIG PACKAGE INIT")
}

func LoadDbConfig() DbConfig {
	return DbConfig{
		Uri:      os.Getenv("NEO4J_URI"),
		Username: os.Getenv("NEO4J_USERNAME"),
		Password: os.Getenv("NEO4J_PASSWORD")}
}
