package core

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_SIGN string = "T7Ffm5oIGaR1Z4MFxNY1o0d68OYX07We"

func TestAuthTokenGen(t *testing.T) {
	os.Setenv("JWT_SIGN", JWT_SIGN)
	testMail := "user@mail.com"
	token, err := CreateAuthToken(testMail)
	if err != nil {
		t.Fatalf("Expected token gen err to be nil but instead it is %v", err)
	}
	var claims GopenClaims
	_, jwtParseErr := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		return []byte(JWT_SIGN), nil
	})
	if jwtParseErr != nil {
		t.Fatalf("Expected token parse err to be nil but instead it is %v", jwtParseErr)
	}
	if claims.Email != testMail {
		t.Fatalf("Expected token email to be %v but instead it is %v", testMail, claims.Email)
	}
	if claims.Issuer != "gopen-api" {
		t.Fatalf("Expected token issuer to be gopen-api but instead it is %v", claims.Issuer)
	}
}

func TestAuthMiddlewareOk(t *testing.T) {
	os.Setenv("JWT_SIGN", JWT_SIGN)
	testMail := "user@mail.com"
	token, _ := CreateAuthToken(testMail)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middlewareFunc := AuthMiddleware(testHandler)
	req := httptest.NewRequest("GET", "http://localhost:3000", nil)
	req.Header.Add("authorization", fmt.Sprintf("Bearer %v", token))
	res := httptest.NewRecorder()
	middlewareFunc.ServeHTTP(res, req)
	resStatus := res.Result().StatusCode
	if resStatus != http.StatusOK {
		t.Fatalf("Expected status to be 200 but instead it is %v", resStatus)
	}
}

func TestAuthMiddlewareNoToken(t *testing.T) {
	os.Setenv("JWT_SIGN", JWT_SIGN)
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middlewareFunc := AuthMiddleware(testHandler)
	req := httptest.NewRequest("GET", "http://localhost:3000", nil)
	res := httptest.NewRecorder()
	middlewareFunc.ServeHTTP(res, req)
	resStatus := res.Result().StatusCode
	if resStatus != http.StatusUnauthorized {
		t.Fatalf("Expected status to be 401 but instead it is %v", resStatus)
	}
}
