package account

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AccountRepository interface {
	Save(account *Account) error
	GetById(id string) (*Account, error)
	GetAll() ([]Account, error)
	GetAllByOwner(ownerID string) ([]Account, error)
	DeleteById(id string) error
	Update(updatedAccount *Account) error
}

type AccountRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *AccountRepositoryPostgres) Save(account *Account) error {
	query := `INSERT INTO accounts (id, currency, balance, customer_id, bank_id, blocked)
	          VALUES (:id, :currency, :balance, :customer_id, :bank_id, :blocked)`

	_, err := repos.db.NamedExec(query, account)

	if err != nil {
		return err
	}

	return nil
}

func (repos *AccountRepositoryPostgres) GetById(id string) (*Account, error) {
	var account *Account = new(Account)

	query := `SELECT id, currency, balance, customer_id, bank_id, blocked 
              FROM accounts WHERE id = $1`

	err := repos.db.Get(account, query, id)

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (repos *AccountRepositoryPostgres) GetAllByOwner(ownerID string) ([]Account, error) {
	var accounts []Account

	query := "SELECT * FROM accounts WHERE customer_id = $1"

	err := repos.db.Select(&accounts, query, ownerID)

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repos *AccountRepositoryPostgres) GetAll() ([]Account, error) {
	var accounts []Account

	query := "SELECT * FROM accounts"

	err := repos.db.Select(&accounts, query)

	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repos *AccountRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM accounts WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *AccountRepositoryPostgres) Update(updatedCusomer *Account) error {
	query := `UPDATE accounts
				SET balance = :balance,
					blocked = :blocked,
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedCusomer)

	if err != nil {
		return err
	}

	return nil
}

func NewAccountRepositoryPostgres(configuration *config.DBConfig) AccountRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS accounts (
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				currency VARCHAR(255) NOT NULL,
				balance DOUBLE PRECISION NOT NULL,
				customer_id VARCHAR(36) NOT NULL,
				bank_id VARCHAR(36) NOT NULL,
				blocked BOOLEAN DEFAULT FALSE
			  );`

	_, err := db.Exec(query)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating account db: %s", err.Error()))
	}
	return &AccountRepositoryPostgres{db: db}
}
