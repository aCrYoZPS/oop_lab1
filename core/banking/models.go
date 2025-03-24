package banking

import "time"

type Account struct {
	ID         string
	Currency   string
	Balance    float64
	CustomerID string
	BankID     string
	Blocked    bool
	Frozen     bool
}

type CompanyAccount struct {
	ID        string
	Currency  string
	Balance   float64
	CompanyID string
	BankID    string
	Blocked   bool
	Frozen    bool
}

type Loan struct {
	AccountID         string
	InterestRate      float64
	InitialSum        float64
	PayedSum          float64
	Duration          int
	CreationDate      time.Time
	NextPayment       time.Time
	RecieverAccountID string
}

type Installment struct {
	AccountID         string
	CreationDate      time.Time
	Duration          int
	InitialSum        float64
	PayedSum          float64
	NextPayment       time.Time
	RecieverAccountID string
}

type Transfer struct {
	ID                string
	SenderAccountID   string
	RecieverAccountID string
	Amount            float64
}

type ExchangeRate struct {
	Sell           float64
	Buy            float64
	TargetCurrency string
	LocalCurrence  string
	BankID         string
}

type Bank struct {
	ID      string
	Country string
	BIC     string
}
