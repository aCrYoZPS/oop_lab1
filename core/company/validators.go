package company

import "slices"

func IsValid(comp *Company) bool {
	if comp.Name == "" || comp.PhoneNumber == "" ||
		comp.Email == "" || comp.Password == "" ||
		comp.Country == "" || comp.Type == "" ||
		comp.BIC == "" || comp.Address == "" ||
		comp.ANP == "" {
		return false
	}

	if !slices.Contains(CompanyTypes, comp.Type) {
		return false
	}

	return true
}
