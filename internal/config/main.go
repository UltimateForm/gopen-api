package config

import (
	"errors"
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

func LoadAuthConfig() AuthConfig {
	jwtSign := os.Getenv("JWT_SIGN")
	if jwtSign == "" {
		panic(errors.New("JWT_SIGN not in environment variables"))
	}
	return AuthConfig{JwtSign: jwtSign}
}
