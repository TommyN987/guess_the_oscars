package service

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/TommyN987/guess_the_oscars/backend/internal/domain"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	rootDir := filepath.Join("..", ".env")

	if err := godotenv.Load(rootDir); err != nil {
		log.Println("No .env file found, loading environment variables from system.")
	}
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is not set.")
	}
	jwtSecret = []byte(secret)
}

func GenerateJWT(user domain.User) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
}
