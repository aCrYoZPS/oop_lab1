package api_v1

import (
	"net/http"
	"oopLab1/config"
	"oopLab1/core/account"
	"oopLab1/core/transactions"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var transactionService = transactions.NewTransactionService(config.GetConfig().Database)

func CreateSelfTransaction(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)
	acc_id := ctx.Param("acc_id")

	acc, err := accountService.GetAccountByID(acc_id)

	if acc.CustomerID != userID && (role == "customer" || role == "company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Access to other acc banking account is prohibited",
		})
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	trans, err := transactionService.GetAllByAccount(acc.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	for _, transaction := range trans {
		account.ApplyTransaction(acc, &transaction)
	}

	transact := &transactions.Transaction{}
	delta := &transactions.AccountDelta{}

	if err = ctx.Bind(delta); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	delta.Blocked = acc.Blocked

	transact.ID = uuid.NewString()
	transact.MoneyDelta = delta.MoneyDelta
	transact.Blocked = delta.Blocked
	transact.ActorID = userID
	transact.SrcAccountID = acc_id
	transact.DestAccountID = acc_id
	transact.Date = time.Now()

	if acc.Balance+transact.MoneyDelta < 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Inapplicable transaction",
		})
	}

	err = transactionService.CreateTransaction(transact)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed to create transaction",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Transaction succesfull",
	})
}
