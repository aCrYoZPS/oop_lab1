package customer

type Customer struct {
	ID             string `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	PhoneNumber    string `json:"phone_number" db:"phone_number"`
	Email          string `json:"email" db:"email"`
	Password       string `json:"password" db:"password"`
	Country        string `json:"country" db:"country"`
	PassportNumber string `json:"passport_number" db:"passport_number"`
	PassportID     string `json:"passport_id" db:"passport_id"`
	AccessAllowed  bool   `json:"access_allowed" db:"access_allowed"`
}
