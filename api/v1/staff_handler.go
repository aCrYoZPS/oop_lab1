package api_v1

import (
	"oopLab1/config"
	"oopLab1/core/staff"
)

var staffService = staff.NewStaffMemberService(config.GetConfig().Database)
