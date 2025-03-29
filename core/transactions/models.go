package transactions

import "time"

type TransactionType int

const (
	MoneyTransfer TransactionType = iota
	Withdrawal
	TopUp

	AccountBlock
	AccountUnblock
)

type AccountDelta struct {
	MoneyDelta float64 `db:"money_delta" json:"money_delta,omitempty"`
	Blocked    bool    `db:"blocked" json:"blocked,omitempty"`
}

type Transaction struct {
	ID            string          `db:"id" json:"id,omitempty"`
	Type          TransactionType `db:"type" json:"type,omitempty"`
	Date          time.Time       `db:"date" json:"date"`
	ActorID       string          `db:"actor_id" json:"actor_id,omitempty"`
	SrcAccountID  string          `db:"src_account_id" json:"src_account_id,omitempty"`
	DestAccountID string          `db:"dest_account_id" json:"dest_account_id,omitempty"`
	Delta         AccountDelta    `json:"delta"`
}
