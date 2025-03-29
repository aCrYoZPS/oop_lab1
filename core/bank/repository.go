package bank

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type BankRepository interface {
	Save(bank *Bank) error
	GetById(id string) (*Bank, error)
	GetAll() ([]Bank, error)
	DeleteById(id string) error
	Update(updatedBank *Bank) error
}

type BankRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *BankRepositoryPostgres) Save(bank *Bank) error {
	query := `INSERT INTO banks (id, name, country, bic)
			  VALUES (:id, :name, :country, :bic)`

	_, err := repos.db.NamedExec(query, bank)

	return err
}

func (repos *BankRepositoryPostgres) GetById(id string) (*Bank, error) {
	var bank Bank
	query := `SELECT * FROM banks WHERE id = $1`

	err := repos.db.Get(&bank, query, id)

	return &bank, err
}

func (repos *BankRepositoryPostgres) GetAll() ([]Bank, error) {
	var banks []Bank
	query := "SELECT * FROM banks"

	err := repos.db.Select(&banks, query)

	return banks, err
}

func (repos *BankRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM banks WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *BankRepositoryPostgres) Update(updatedBank *Bank) error {
	query := `UPDATE banks
				SET name = :name,
					country = :country,
					bic = :bic
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedBank)

	return err
}

func NewBankRepositoryPostgres(configuration *config.DBConfig) BankRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS banks(
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				name VARCHAR(255) NOT NULL,
				country VARCHAR(100) NOT NULL,
				bic VARCHAR(11) UNIQUE NOT NULL
			  );`

	_, err := db.Exec(query)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating bank db: %s", err.Error()))
	}

	return &BankRepositoryPostgres{db: db}
}
