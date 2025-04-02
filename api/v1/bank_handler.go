package api_v1

import (
	"fmt"
	"net/http"
	"oopLab1/config"
	"oopLab1/core/bank"
	"oopLab1/core/staff"
	myJWT "oopLab1/pkg/jwt"
	"oopLab1/pkg/logger"
	"oopLab1/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var bankService = bank.NewBankService(config.GetConfig().Database)

func RegisterBank(ctx echo.Context) error {
	bankRegRequest := &bank.BankRegistrationRequest{}

	if err := ctx.Bind(bankRegRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	bankInstance := &bankRegRequest.Bank
	bankInstance.ID = uuid.New().String()

	bankAdmin := &bankRegRequest.Admin
	bankAdmin.ID = uuid.NewString()
	bankAdmin.Role = "admin"
	bankAdmin.BankID = bankInstance.ID

	if !bank.IsValid(bankInstance) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid bank data",
		})
	}

	if !staff.IsValid(bankAdmin) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid admin data",
		})
	}

	err := bankService.CreateBank(bankInstance)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	err = staffService.CreateStaffMember(bankAdmin)
	if err != nil {
		bankService.DeleteBank(bankInstance.ID)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	logger.Info(fmt.Sprintf("Created bank with id %s", bankInstance.ID))

	token, err := myJWT.GenerateJWT(bankAdmin.ID, "admin")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"token": token,
		"id":    bankAdmin.ID,
	})
}

func DeleteBank(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if role != "admin" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Only administrator can make that decision",
		})
	}

	admin, err := staffService.GetStaffMemberByID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	err = bankService.DeleteBank(admin.BankID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func UpdateBank(ctx echo.Context) error {
	bank_id := ctx.Param("bank_id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	admin, err := staffService.GetStaffMemberByID(userID)
	if admin.Role != "admin" || admin.BankID != bank_id {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Only administrator can make that decision",
		})
	}

	currentBank, err := bankService.GetBankByID(bank_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	updatedBank := &bank.Bank{}
	ctx.Bind(updatedBank)

	if !bank.IsValid(updatedBank) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid customer data",
		})
	}

	bank.UpdateBankInfo(currentBank, updatedBank)

	err = bankService.UpdateBank(updatedBank)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedBank)
}

func GetBankByID(ctx echo.Context) error {
	bank_id := ctx.Param("bank_id")

	bank, err := bankService.GetBankByID(bank_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, bank)
}

func GetAllBanks(ctx echo.Context) error {
	banks, err := bankService.GetAllBanks()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, banks)
}
