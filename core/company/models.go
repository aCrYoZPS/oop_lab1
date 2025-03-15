package company

import "time"

type CompanyType int

const (
	IE   CompanyType = iota //ИП
	LLC                     //ООО
	CJSC                    //ЗАО
	JSC                     //ОАО
)

var companyTypeName = map[CompanyType]string{
	IE:   "IE",
	LLC:  "LLC",
	CJSC: "CJSC",
	JSC:  "JSC",
}

func (ct CompanyType) String() string {
	return companyTypeName[ct]
}

type Company struct {
	ID           string
	Name         string
	PhoneNumber  string
	Email        string
	Password     string
	Country      string
	Type         CompanyType
	BIC          string // Bank Identification Code (БИК)
	Address      string
	ANP          string // Accounting Number of Payer (УНП)
	AccesAllowed bool
	AccountID    string
}

type SallaryProject struct {
	EmployeeAccountID string
	EmployerAccountID string
	Sallary           float64
	PaymentDate       time.Time
}
