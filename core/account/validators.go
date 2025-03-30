package account

import "slices"

func IsValid(request *AccountRequest) bool {
	if !slices.Contains(currencies, request.Currency) {
		return false
	}
	return true
}
