package api_v1

import (
	"net/http"
	"oopLab1/config"
	"oopLab1/core/account"
	"oopLab1/core/transactions"
	"oopLab1/utils"
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
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if acc.CustomerID != userID && (role == "customer" || role == "company") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Access to other acc banking account is prohibited",
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

	if acc.Blocked {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "account blocked",
		})
	}

	delta.Blocked = acc.Blocked

	transact.ID = uuid.NewString()
	transact.MoneyDelta = delta.MoneyDelta
	transact.Blocked = delta.Blocked
	transact.ActorID = userID
	transact.SrcAccountID = acc_id
	transact.DestAccountID = acc_id
	if transact.MoneyDelta > 0 {
		transact.Type = transactions.TopUp
	} else {
		transact.Type = transactions.Withdrawal
	}
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

func CreateTransaction(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)
	acc_id := ctx.Param("acc_id")
	target_id := ctx.Param("target_id")

	acc, err := accountService.GetAccountByID(acc_id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if acc.CustomerID != userID {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Forbidden",
		})
	}

	if acc.Blocked {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "account blocked",
		})
	}

	_, err = accountService.GetAccountByID(target_id)
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

	if acc.Blocked {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "blocked",
		})
	}

	if delta.MoneyDelta > 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "unable to take money from other account",
		})
	}

	transact.ID = uuid.NewString()
	transact.MoneyDelta = delta.MoneyDelta
	transact.Blocked = delta.Blocked
	transact.ActorID = userID
	transact.SrcAccountID = acc_id
	transact.Type = transactions.MoneyTransfer
	transact.DestAccountID = target_id
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

func UndoRecentTransactions(ctx echo.Context) error {
	acc_id := ctx.Param("acc_id")
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if utils.GetPrivelegeLevel(role) < utils.GetPrivelegeLevel("manager") {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "only managers and admins can undo transactions",
		})
	}

	accTransactions, err := transactionService.GetAllByAccount(acc_id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	trans := accTransactions[0]
	err = transactionService.DeleteTransaction(trans.ID)
	if err != nil || trans.Type == transactions.Withdrawal || trans.Type == transactions.TopUp {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	trans = accTransactions[1]
	err = transactionService.DeleteTransaction(trans.ID)
	if err != nil || trans.Type == transactions.Withdrawal || trans.Type == transactions.TopUp {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "one transaction undone succesfully",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Transactions undone succesfully",
	})
}
