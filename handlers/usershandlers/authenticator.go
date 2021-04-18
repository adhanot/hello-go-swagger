package usershandlers

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(pseudo, email, fname, lname string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["pseudo"] = pseudo
	claims["email"] = email
	claims["fname"] = fname
	claims["lname"] = lname
	claims["date_created_at"] = time.Now()
	claims["date_expired_at"] = time.Now().Add(time.Hour * 2)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
