package utils

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret string
	once      sync.Once
)

func loadSecret() {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET is not set in environment variables")
	}
	jwtSecret = secret
}

func GetSecretKey() string {
	once.Do(loadSecret)
	return jwtSecret
}

func GenerateJWT(email string) (string, error) {
	jwtSecret := GetSecretKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) bool {
	jwtSecret := GetSecretKey()
	if jwtSecret == "" {
		fmt.Println("Failed to validate JWT: secret key is not set")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	return err == nil && token.Valid
}
