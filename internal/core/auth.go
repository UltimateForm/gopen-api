package core

import (
	"net/http"
	"strings"
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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			RespondError(res, req, Unauthorized())
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		_, parseErr := jwt.Parse(token, func(t *jwt.Token) (any, error) {
			return []byte(config.LoadAuthConfig().JwtSign), nil
		})
		if parseErr != nil {
			RespondError(res, req, Unauthorized())
			return
		}
		next.ServeHTTP(res, req)
	})
}
