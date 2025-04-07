package loans

import (
	"time"

	"github.com/google/uuid"
)

type LoanRequest struct {
	Sum     float64   `db:"sum" json:"sum,omitempty"`
	Percent int32     `db:"percent" json:"percent,omitempty"`
	EndDate time.Time `db:"end_date" json:"end_date"`
}

type Loan struct {
	ID          string    `db:"id" json:"id,omitempty"`
	CustomerID  string    `db:"customer_id" json:"customer_id,omitempty"`
	Sum         float64   `db:"sum" json:"sum,omitempty"`
	PaidSum     float64   `db:"paid_sum" json:"paid_sum,omitempty"`
	Percent     int32     `db:"percent" json:"percent,omitempty"`
	PaymentDate time.Time `db:"payment_date" json:"payment_date"`
	EndDate     time.Time `db:"end_date" json:"end_date"`
}

func NewLoanFromRequest(request *LoanRequest) *Loan {
	return &Loan{
		ID:          uuid.NewString(),
		Sum:         request.Sum,
		Percent:     request.Percent,
		PaymentDate: time.Now().Add(time.Hour * 24 * 30),
		EndDate:     request.EndDate,
	}
}

func UpdateLoanInfo(original *Loan, updated *Loan) {
	updated.ID = original.ID
	updated.CustomerID = original.CustomerID
	updated.Sum = original.Sum
	updated.Percent = original.Percent
	updated.EndDate = original.EndDate
}
