package account

import "slices"

func IsValid(acc *Account) bool {
	if slices.Contains(currencies, acc.Currency) {
		return false
	}
	return true
}
