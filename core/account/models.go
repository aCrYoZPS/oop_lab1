package account

var currencies = []string{"USD", "RUB", "BYN", "EUR"}

type Account struct {
	ID         string  `json:"id,omitempty" db:"id"`
	Currency   string  `json:"currency,omitempty" db:"currency"`
	Balance    float64 `json:"balance,omitempty" db:"balance"`
	CustomerID string  `json:"customer_id,omitempty" db:"customer_id"`
	BankID     string  `json:"bank_id,omitempty" db:"bank_id"`
	Blocked    bool    `json:"blocked,omitempty" db:"blocked"`
}
