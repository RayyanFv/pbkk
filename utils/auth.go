package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"pos-system/models"

	"github.com/golang-jwt/jwt/v4"

	"github.com/joho/godotenv"
)

// Load .env variables
func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

// Claims represents the JWT claims
type Claims struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT token for a user
func GenerateJWT(user models.User) (string, error) {
	// Get the secret key from environment variables
	secretKey := os.Getenv("SECRET_KEY")

	// Set the expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims
	claims := &Claims{
		ID:   user.ID,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "pos-system",
		},
	}

	// Create a new token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWT parses the JWT token and returns the claims
func ParseJWT(tokenString string) (*Claims, error) {
	secretKey := os.Getenv("SECRET_KEY")

	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the token's signing method matches the expected one
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid and return the claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
