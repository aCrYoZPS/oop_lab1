package api_v1

import (
	"net/http"
	"oopLab1/config"
	"oopLab1/core/transactions"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var transactionService = transactions.NewTransactionService(config.GetConfig().Database)

func CreateSelfTransaction(ctx echo.Context) error {
	id := ctx.Param("id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	role := claims["role"].(string)
	acc_id := ctx.Param("acc_id")

	acc, err := accountService.GetAccountByID(acc_id)

	if (id != userID || acc.CustomerID != userID) && (role == "customer" || role == "company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Access to other acc banking account is prohibited",
		})
	}

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	var transact = &transactions.Transaction{}

	delta := &transactions.AccountDelta{}

	if err = ctx.Bind(delta); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	transact.ID = uuid.NewString()
	transact.Delta = *delta
	transact.ActorID = userID
	transact.SrcAccountID = acc_id
	transact.DestAccountID = acc_id
	transact.Date = time.Now()

	if !transactions.IsApplicable(transact, acc) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Transaction inapplicable",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Transaction succesfull",
	})

}
