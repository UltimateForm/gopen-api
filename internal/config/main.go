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

func LoadAuthConfig() AuthConfig {
	jwtSign, err := os.ReadFile("/run/secrets/jwt_sign")
	if err != nil {
		if os.IsNotExist(err) {
			// TODO: remove this fallback secret management is set
			// will need to change tests too
			jwtSign = []byte(os.Getenv("JWT_SIGN"))
		} else {
			panic(err)
		}
	}
	return AuthConfig{JwtSign: string(jwtSign)}
}
