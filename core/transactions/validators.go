package transactions

import "oopLab1/core/account"

func IsApplicable(transaction *Transaction, acc *account.Account) bool {
	return acc.Balance+transaction.Delta.MoneyDelta >= 0
}
