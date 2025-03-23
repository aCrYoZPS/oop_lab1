package customer

type Customer struct {
	ID             string `json:"id,omitempty" db:"id"`
	Name           string `json:"name,omitempty" db:"name"`
	PhoneNumber    string `json:"phone_number,omitempty" db:"phone_number"`
	Email          string `json:"email,omitempty" db:"email"`
	Password       string `json:"password,omitempty" db:"password"`
	Country        string `json:"country,omitempty" db:"country"`
	PassportNumber string `json:"passport_number,omitempty" db:"passport_number"`
	PassportID     string `json:"passport_id,omitempty" db:"passport_id"`
	AccessAllowed  bool   `json:"access_allowed,omitempty" db:"access_allowed"`
}
