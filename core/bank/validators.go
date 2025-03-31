package bank

func IsValid(bank *Bank) bool {
	if bank.Name == "" || bank.Country == "" ||
		bank.BIC == "" {
		return false
	}

	return true
}
