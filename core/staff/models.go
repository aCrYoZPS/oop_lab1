package staff

type Role struct {
	ID          string
	Name        string
	AccessLevel int
}

type StaffMember struct {
	ID          string
	Name        string
	PhoneNumber string
	Email       string
	Password    string
	BankID      string
	Role        Role
}
