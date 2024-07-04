package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaim struct {
	ID    string `json:"id"`
	User  string `json:"user"`
	Admin bool   `json:"role"`
	jwt.RegisteredClaims
}

func CreateNewAuthToken(id string, email string, isAdmin bool) (string, error) {
	cliams := AuthClaim{
		ID:    id,
		User:  email,
		Admin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "searchengine.com",
		},
	}

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	secretKey := os.Getenv("SECRET_KEY")

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("error signing the token")
	}
	return signedToken, nil
}
