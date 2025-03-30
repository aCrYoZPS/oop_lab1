package jwt

import (
	"oopLab1/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID string, role string) (string, error) {
	secretKey := utils.GetJWTSecret()

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}
