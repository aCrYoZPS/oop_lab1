package api_v1

import (
	"net/http"
	"oopLab1/config"
	"oopLab1/core/account"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var accountService = account.NewAccountService(config.GetConfig().Database)

func CreateAccount(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	accRequest := &account.AccountRequest{}

	if err := ctx.Bind(accRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request",
		})
	}

	_, err := bankService.GetBankByID(accRequest.BankID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request",
		})
	}

	acc, err := account.NewAccountFromRequest(accRequest, userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	acc.ID = uuid.New().String()

	err = accountService.CreateAccount(acc)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"id": acc.ID,
	})
}

func GetAccount(ctx echo.Context) error {
	acc_id := ctx.Param("acc_id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	acc, err := accountService.GetAccountByID(acc_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != acc.CustomerID && (role == "customer" || role == "company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Access to other acc banking account is prohibited",
		})
	}

	trans, err := transactionService.GetAllByAccount(acc.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	for _, transaction := range trans {
		account.ApplyTransaction(acc, &transaction)
	}

	return ctx.JSON(http.StatusOK, acc)
}

func GetAllAccountsByOwner(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	accounts, err := accountService.GetAllAccountsByOwner(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	for index, acc := range accounts {
		trans, err := transactionService.GetAllByAccount(acc.ID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}
		for _, transaction := range trans {
			account.ApplyTransaction(&accounts[index], &transaction)
		}
	}

	return ctx.JSON(http.StatusOK, accounts)
}

func GetAllAccounts(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role == "customer" || role == "company" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other acc banking account is prohibited",
		})
	}

	accounts, err := accountService.GetAllAccounts()
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	for index, acc := range accounts {
		trans, err := transactionService.GetAllByAccount(acc.ID)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}
		for _, transaction := range trans {
			account.ApplyTransaction(&accounts[index], &transaction)
		}
	}

	return ctx.JSON(http.StatusOK, accounts)
}

func DeleteAccount(ctx echo.Context) error {
	acc_id := ctx.Param("acc_id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	acc, err := accountService.GetAccountByID(acc_id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != acc.CustomerID && (role == "customer" || role == "company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other acc banking account is prohibited",
		})
	}

	err = accountService.DeleteAccount(acc_id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}
