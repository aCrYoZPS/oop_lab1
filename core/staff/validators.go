package staff

import "slices"

func IsValid(staffMem *StaffMember) bool {
	if staffMem.Name == "" || staffMem.PhoneNumber == "" ||
		staffMem.Email == "" || staffMem.Password == "" ||
		!slices.Contains(staffRoles, staffMem.Role) {
		return false
	}
	return true
}
