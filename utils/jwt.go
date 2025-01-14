package utils

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var (
	jwtSecret string
	once      sync.Once
)

func loadSecret() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		fmt.Println("JWT_SECRET is not set in environment variables")
	} else {
		fmt.Println("JWT_SECRET successfully loaded")
	}
	jwtSecret = secret
}

func GetSecretKey() string {
	once.Do(loadSecret)
	return jwtSecret
}

func GenerateJWT(email string) (string, error) {
	jwtSecret := []byte(GetSecretKey())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Set expiration time
	})

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) bool {
	jwtSecret := GetSecretKey()
	if jwtSecret == "" {
		fmt.Println("Failed to validate JWT: secret key is not set")
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		fmt.Printf("JWT validation error: %v\n", err)
		return false
	}

	return token.Valid
}
