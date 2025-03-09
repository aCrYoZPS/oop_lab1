package entities

import "time"

type TransactionType int

const (
	MoneyTransfer TransactionType = iota
	Withdrawal

	AccountCreation
	AccountDeletion

	AccountBlock
	AccountUnblock

	AccountFreeze
	AccountUnfreeze
)

type AccountDelta struct {
	MoneyDelta float64
	Deleted    bool
	Created    bool
	Blocked    bool
	Frozen     bool
}

type Transaction struct {
	ID      string
	Type    TransactionType
	Date    time.Time
	ActorID string
}
