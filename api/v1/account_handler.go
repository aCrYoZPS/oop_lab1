package api_v1

import (
	"fmt"
	"net/http"
	"oopLab1/auth"
	"oopLab1/config"
	"oopLab1/core/account"
	myJWT "oopLab1/pkg/jwt"
	"oopLab1/pkg/logger"
	"oopLab1/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var accountService = account.NewAccountService(config.GetConfig().Database)

func GetAccount(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	acc, err := accountService.GetAccountByID(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != acc.CustomerID && (role == "customer" || role == "company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Access to other acc banking account is prohibited",
		})
	}

	return ctx.JSON(http.StatusOK, acc)
}

func GetAllAccountsByOwner(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	user_id := claims["user_id"].(string)
	role := claims["role"].(string)

	if id != user_id && (role == "customer" || role == "accany") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Access to others accounts is prohibited",
		})
	}

	accounts, err := accountService.GetAllAccountsByOwner(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
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
	return ctx.JSON(http.StatusOK, accounts)
}

func DeleteAccount(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)

	acc, err := accountService.GetAccountByID(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if userID != acc.CustomerID && role == "account" {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Acces to other acc banking account is prohibited",
		})
	}

	err = accountService.DeleteAccount(id)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return ctx.NoContent(http.StatusNoContent)
}
