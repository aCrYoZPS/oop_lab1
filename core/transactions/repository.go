package transactions

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TransactionRepository interface {
	Save(Transaction *Transaction) error
	GetById(id string) (*Transaction, error)
	GetAll() ([]Transaction, error)
	DeleteById(id string) error
	GetAllInbound(account_id string) ([]Transaction, error)
	GetAllOutbound(account_id string) ([]Transaction, error)
	GetAllByActor(actor_id string) ([]Transaction, error)
}

type TransactionRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *TransactionRepositoryPostgres) Save(transaction *Transaction) error {
	query := `
        INSERT INTO transactions 
            (id, type, date, actor_id, src_account_id, dest_account_id, money_delta, blocked)
        VALUES 
            (:id, :type, :date, :actor_id, :src_account_id, :dest_account_id, :money_delta, :blocked)
    `
	_, err := repos.db.NamedExec(query, transaction)
	return err
}

func (repos *TransactionRepositoryPostgres) GetById(id string) (*Transaction, error) {
	query := `SELECT * FROM transactions WHERE id = $1`
	transaction := &Transaction{}

	err := repos.db.Get(transaction, query, id)

	return transaction, err
}

func (repos *TransactionRepositoryPostgres) GetAll() ([]Transaction, error) {
	var transactions []Transaction
	query := "SELECT * FROM transactions"

	err := repos.db.Select(&transactions, query)

	return transactions, err
}

func (repos *TransactionRepositoryPostgres) GetAllOutbound(account_id string) ([]Transaction, error) {
	var transactions []Transaction
	query := "SELECT * FROM transactions WHERE src_account_id = $1"

	err := repos.db.Select(&transactions, query, account_id)

	return transactions, err
}

func (repos *TransactionRepositoryPostgres) GetAllInbound(account_id string) ([]Transaction, error) {
	var transactions []Transaction
	query := "SELECT * FROM transactions WHERE dest_account_id = $1"

	err := repos.db.Select(&transactions, query, account_id)

	return transactions, err
}

func (repos *TransactionRepositoryPostgres) GetAllByActor(actor_id string) ([]Transaction, error) {
	var transactions []Transaction
	query := "SELECT * FROM transactions WHERE actor_id = $1"

	err := repos.db.Select(&transactions, query, actor_id)

	return transactions, err
}

func (repos *TransactionRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM transactions WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func NewTransactionRepositoryPostgres(configuration *config.DBConfig) TransactionRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS transactions(
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				type INT NOT NULL,
				date TIMESTAMP NOT NULL,
				actor_id VARCHAR(36) NOT NULL,
				src_account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
				dest_account_id VARCHAR(36) NOT NULL REFERENCES accounts(id),
				money_delta DOUBLE PRECISION NOT NULL,
				blocked BOOL NOT NULL,
			  );`

	_, err := db.Exec(query)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating transaction db: %s", err.Error()))
	}
	return &TransactionRepositoryPostgres{db: db}
}
