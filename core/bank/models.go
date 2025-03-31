package bank

import "oopLab1/core/staff"

type Bank struct {
	ID        string `db:"id" json:"id,omitempty"`
	Name      string `db:"name" json:"name,omitempty"`
	Country   string `db:"country" json:"country,omitempty"`
	BIC       string `db:"bic" json:"bic,omitempty"`
	AccountID string `db:"account_id" json:"account_id,omitempty"`
}

type BankRegistrationRequest struct {
	Bank  Bank              `json:"bank"`
	Admin staff.StaffMember `json:"admin"`
}
