package core

import (
	"time"

	"github.com/UltimateForm/gopen-api/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type GopenClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func CreateAuthToken(email string) (string, error) {
	claims := GopenClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gopen-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LoadAuthConfig().JwtSign))
}
