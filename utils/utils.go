package utils

import (
	"oopLab1/core/account"
	"oopLab1/core/transactions"
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

func IsApplicable(transaction *transactions.Transaction, acc *account.Account) bool {
	return (acc.Balance + transaction.MoneyDelta) >= 0
}

func GetPrivelegeLevel(role string) int {
	switch role {
	case "customer":
		return 0
	case "company":
		return 1
	case "operator":
		return 2
	case "manager":
		return 3
	case "admin":
		return 4
	default:
		return -1
	}
}
