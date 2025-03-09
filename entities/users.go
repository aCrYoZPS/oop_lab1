package entities

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

type Role struct {
	ID          string
	Name        string
	AccessLevel int
}

type UserBase struct {
	ID          string
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	Country     string
}

type StaffMember struct {
	User   UserBase
	BankID string
	Role   Role
}

type Customer struct {
	User           UserBase
	PassportNumber string
	PassportID     string
}

func (ct CompanyType) String() string {
	return companyTypeName[ct]
}

type Company struct {
	User    UserBase
	Type    CompanyType
	BIC     string // Bank Identification Code (БИК)
	Address string
	ANP     string // Accounting Number of Payer (УНП)
}
