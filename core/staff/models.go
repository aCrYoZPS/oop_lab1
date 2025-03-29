package staff

var staffRoles = []string{
	"operator", "admin",
}

type StaffMember struct {
	ID          string `db:"id" json:"id,omitempty"`
	Name        string `db:"name" json:"name,omitempty"`
	PhoneNumber string `db:"phone_number" json:"phone_number,omitempty"`
	Email       string `db:"email" json:"email,omitempty"`
	Password    string `db:"password" json:"password,omitempty"`
	BankID      string `db:"bank_id" json:"bank_id,omitempty"`
	Role        string `db:"role" json:"role,omitempty"`
}
