package company

import "time"

var CompanyTypes = []string{
	"IE",
	"LLC",
	"CJSC",
	"JSC",
}

type Company struct {
	ID            string `json:"id,omitempty" db:"id"`
	Name          string `json:"name,omitempty" db:"name"`
	PhoneNumber   string `json:"phone_number,omitempty" db:"phone_number"`
	Email         string `json:"email,omitempty" db:"email"`
	Password      string `json:"password,omitempty" db:"password"`
	Country       string `json:"country,omitempty" db:"country"`
	Type          string `json:"type,omitempty" db:"type"`
	BIC           string `json:"bic,omitempty" db:"bic"` // Bank Identification Code (БИК)
	Address       string `json:"address,omitempty" db:"address"`
	ANP           string `json:"anp,omitempty" db:"anp"` // Accounting Number of Payer (УНП)
	AccessAllowed bool   `json:"access_allowed,omitempty" db:"access_allowed"`
}

type SallaryProject struct {
	EmployeeAccountID string
	CompanyAccountID  string
	Sallary           float64
	PaymentDate       time.Time
}
