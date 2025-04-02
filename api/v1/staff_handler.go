package api_v1

import (
	"fmt"
	"net/http"
	"oopLab1/auth"
	"oopLab1/config"
	"oopLab1/core/staff"
	myJWT "oopLab1/pkg/jwt"
	"oopLab1/pkg/logger"
	"oopLab1/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var staffService = staff.NewStaffMemberService(config.GetConfig().Database)

func RegisterStaffMember(ctx echo.Context) error {
	staffMember := new(staff.StaffMember)

	if err := ctx.Bind(staffMember); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if !staff.IsValid(staffMember) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid staff data",
		})
	}

	staffMember.ID = uuid.NewString()
	logger.Info(fmt.Sprintf("Created staff member with id %s", staffMember.ID))

	token, err := myJWT.GenerateJWT(staffMember.ID, staffMember.Role)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	err = staffService.CreateStaffMember(staffMember)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"token": token,
		"id":    staffMember.ID,
	})
}

func LoginStaffMember(ctx echo.Context) error {
	var loginRequest = new(auth.LoginRequest)

	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid login request",
		})
	}

	staffMember, err := staffService.GetStaffMemberByEmail(loginRequest.Username)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": "Staff member with such username doesn't exist",
		})
	}

	if staffMember.Password != loginRequest.Password {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Wrong password",
		})
	}

	token, err := myJWT.GenerateJWT(staffMember.ID, staffMember.Role)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"token": token,
		"id":    staffMember.ID,
	})
}

func GetStaffMember(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	staffMember, err := staffService.GetStaffMemberByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != id && utils.GetPrivelegeLevel(role) <= utils.GetPrivelegeLevel(staffMember.Role) {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other staff members is prohibited",
		})
	}

	return ctx.JSON(http.StatusOK, staffMember)
}

func DeleteStaffMember(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	staffMember, err := staffService.GetStaffMemberByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != id && utils.GetPrivelegeLevel(role) <= utils.GetPrivelegeLevel(staffMember.Role) {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other staff members is prohibited",
		})
	}

	err = staffService.DeleteStaffMember(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func UpdateStaffMember(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	staffMember, err := staffService.GetStaffMemberByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != id && utils.GetPrivelegeLevel(role) <= utils.GetPrivelegeLevel(staffMember.Role) {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other staff members is prohibited",
		})
	}

	updatedStaffMember := &staff.StaffMember{}

	if err := ctx.Bind(updatedStaffMember); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	staff.UpdateStaffMemberInfo(staffMember, updatedStaffMember)

	if !staff.IsValid(updatedStaffMember) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid staff data",
		})
	}

	err = staffService.UpdateStaffMember(updatedStaffMember)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedStaffMember)
}
