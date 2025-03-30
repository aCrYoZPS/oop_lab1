package account

import "errors"

var currencies = []string{"USD", "RUB", "BYN", "EUR"}

type AccountRequest struct {
	Currency string `json:"currency"`
	BankID   string `json:"bank_id"`
}
type Account struct {
	ID         string  `json:"id" db:"id"`
	Currency   string  `json:"currency" db:"currency"`
	Balance    float64 `json:"balance" db:"balance"`
	CustomerID string  `json:"customer_id" db:"customer_id"`
	BankID     string  `json:"bank_id" db:"bank_id"`
	Blocked    bool    `json:"blocked" db:"blocked"`
}

func NewAccountFromRequest(request *AccountRequest, owner_id string) (*Account, error) {
	if !IsValid(request) {
		return nil, errors.New("Invalid account requesto")
	}

	return &Account{
		Currency:   request.Currency,
		BankID:     request.BankID,
		CustomerID: owner_id,
	}, nil
}
