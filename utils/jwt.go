package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var ExpirationTime = time.Now().Add(240 * time.Minute)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID *string) (*string, error) {
	claims := &Claims{
		UserID: *userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(ExpirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("bananacomaçucar"))
	if err != nil {
		return nil, err
	}

	return PString(tokenString), nil
}

func ValidationBearerToken(token string) (*string, error) {
	//Validar Token
	if len(token) <= 0 {
		return nil, errors.New("invalid token")
	}

	x := strings.Split(token, "Bearer ")
	if len(x) < 2 {
		return nil, errors.New("invalid token")
	}

	return &x[1], nil
}
