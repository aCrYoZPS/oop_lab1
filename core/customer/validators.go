package customer

func IsValid(customer *Customer) bool {
	if customer.Name == "" || customer.PhoneNumber == "" ||
		customer.Email == "" || customer.Password == "" ||
		customer.Country == "" || customer.PassportNumber == "" ||
		customer.PassportID == "" {
		return false
	}
	return true
}
