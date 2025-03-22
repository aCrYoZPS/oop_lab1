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
	Update(updatedTransaction *Transaction) error
}

type TransactionRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *TransactionRepositoryPostgres) Save(Transaction *Transaction) error {
	query := `INSERT INTO transactions (name, phone_number, email, password, country, passport_number, passport_id, access_allowed)
              VALUES (:name, :phone_number, :email, :password, :country, :passport_number, :passport_id, :access_allowed)`

	_, err := repos.db.NamedExec(query, Transaction)
	if err != nil {
		return err
	}

	return nil
}

func (repos *TransactionRepositoryPostgres) GetById(id string) (*Transaction, error) {
	var Transaction *Transaction = new(Transaction)

	query := `SELECT id, name, phone_number, email, password, country, passport_number, passport_id, access_allowed 
              FROM customers WHERE id = $1`

	err := repos.db.Get(Transaction, query, id)
	if err != nil {
		return nil, err
	}

	return Transaction, nil
}

func (repos *TransactionRepositoryPostgres) GetAll() ([]Transaction, error) {
	var transactions []Transaction

	query := "SELECT * FROM customers"

	err := repos.db.Select(&transactions, query)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (repos *TransactionRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM transactions WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *TransactionRepositoryPostgres) Update(updatedCusomer *Transaction) error {
	query := `UPDATE customers
				SET name = :name,
					phone_number = :phone_number,
					email = :email,
					password = :password,
					country = :country,
					passport_number = :passport_number,
					passport_id = :passport_id,
					access_allowed = :access_allowed
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedCusomer)

	if err != nil {
		return err
	}

	return nil
}

func NewTransactionRepositoryPostgres(configuration config.DBConfig) TransactionRepository {
	db := storage_postgres.GetPostgresDB(&configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS customers (
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				name VARCHAR(255) NOT NULL,
				phone_number VARCHAR(20),
				email VARCHAR(255) UNIQUE NOT NULL,
				password VARCHAR(255) NOT NULL,
				country VARCHAR(100),
				passport_number VARCHAR(50) UNIQUE NOT NULL,
				passport_id VARCHAR(50) UNIQUE NOT NULL,
				access_allowed BOOLEAN DEFAULT FALSE
			  );`

	_, err := db.Exec(query)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating transaction db: %s", err.Error()))
	}
	return &TransactionRepositoryPostgres{db: db}
}
