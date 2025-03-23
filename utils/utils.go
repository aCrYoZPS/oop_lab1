package utils

import (
	"oopLab1/core/customer"
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

func UpdateCustomerInfo(original *customer.Customer, updated *customer.Customer) {
	updated.ID = original.ID
	updated.AccessAllowed = original.AccessAllowed
}
