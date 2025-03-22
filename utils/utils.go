package utils

import (
	"oopLab1/pkg/logger"
	"os"
)

func GetJWTSecret() []byte {
	secret, err := os.ReadFile("secret_key.txt")
	if err != nil {
		logger.Fatal("Error reading JWT secret")
	}
	return secret
}
