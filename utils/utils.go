package utils

import (
	"oopLab1/core/account"
	"oopLab1/core/company"
	"oopLab1/core/customer"
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

func UpdateCustomerInfo(original *customer.Customer, updated *customer.Customer) {
	updated.ID = original.ID
	updated.AccessAllowed = original.AccessAllowed
}

func UpdateCompanyInfo(original *company.Company, updated *company.Company) {
	updated.ID = original.ID
	updated.AccessAllowed = original.AccessAllowed
}

func IsApplicable(transaction *transactions.Transaction, acc *account.Account) bool {
	return (acc.Balance + transaction.MoneyDelta) >= 0
}
