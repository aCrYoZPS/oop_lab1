package bank

type Bank struct {
	ID      string `json:"id,omitempty" db:"id"`
	Name    string `json:"name,omitempty" db:"name"`
	Country string `json:"country,omitempty" db:"country"`
	BIC     string `json:"bic,omitempty" db:"bic"`
}
