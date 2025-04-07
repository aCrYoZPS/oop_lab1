package api_v1

import (
	"net/http"
	"oopLab1/config"
	"oopLab1/core/loans"
	"oopLab1/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var loanService = loans.NewLoanService(config.GetConfig().Database)

func CreateLoan(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	if role != "customer" && role != "company" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	loanRequest := &loans.LoanRequest{}

	if err := ctx.Bind(loanRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	loan := loans.NewLoanFromRequest(loanRequest)

	loan.CustomerID = userID

	err := loanService.CreateLoan(loan)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create loan",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Transaction succesfull",
	})
}

func GetLoan(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	loan, err := loanService.GetLoanByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if loan.CustomerID != userID && role == "customer" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	return ctx.JSON(http.StatusOK, loan)
}

func GetAllLoansByOwner(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	loans, err := loanService.GetAllByOwner(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, loans)
}

func GetAllLoans(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manger") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "forbidden",
		})
	}

	loans, err := loanService.GetAllLoans()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, loans)
}

func UpdateLoan(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	loan, err := loanService.GetLoanByID(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manger") {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	updatedLoan := &loans.Loan{}

	if err := ctx.Bind(updatedLoan); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	loans.UpdateLoanInfo(loan, updatedLoan)

	if !loans.IsValid(updatedLoan) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid customer data",
		})
	}

	err = loanService.UpdateLoan(updatedLoan)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, updatedLoan)
}

func DeleteLoan(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manger") {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	err := loanService.DeleteLoan(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}
