package company

func IsValid(company *Company) bool {
	if company.Name == "" || company.PhoneNumber == "" ||
		company.Email == "" || company.Password == "" ||
		company.Country == "" || company.Type == "" ||
		company.BIC == "" || company.Address == "" ||
		company.ANP == "" {
		return false
	}
	return true
}
