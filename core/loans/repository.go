package loans

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type LoanRepository interface {
	Save(loan *Loan) error
	GetById(id string) (*Loan, error)
	GetAll() ([]Loan, error)
	GetAllByOwner(id string) ([]Loan, error)
	DeleteById(id string) error
	Update(updatedLoan *Loan) error
}

type LoanRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *LoanRepositoryPostgres) Save(loan *Loan) error {
	query := `INSERT INTO loans (id, customer_id, sum, paid_sum, percent, payment_date, end_date)
	          VALUES (:id, :customer_id, :sum, :paid_sum, :percent, :payment_date, :end_date)`

	_, err := repos.db.NamedExec(query, loan)
	if err != nil {
		return err
	}

	return nil
}

func (repos *LoanRepositoryPostgres) GetById(id string) (*Loan, error) {
	var loan *Loan = new(Loan)

	query := `SELECT id, customer_id, sum, paid_sum, percent, payment_date, end_date 
              FROM loans WHERE id = $1`

	err := repos.db.Get(loan, query, id)
	if err != nil {
		return nil, err
	}

	return loan, nil
}

func (repos *LoanRepositoryPostgres) GetAll() ([]Loan, error) {
	var loans []Loan

	query := "SELECT * FROM loans"

	err := repos.db.Select(&loans, query)
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (repos *LoanRepositoryPostgres) GetAllByOwner(id string) ([]Loan, error) {
	var loans []Loan

	query := "SELECT * FROM loans WHERE customer_id = $1"

	err := repos.db.Select(&loans, query, id)
	if err != nil {
		return nil, err
	}

	return loans, nil
}

func (repos *LoanRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM loans WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *LoanRepositoryPostgres) Update(updatedLoan *Loan) error {
	query := `UPDATE loans 
				SET paid_sum = :paid_sum,
					percent = :percent,
					payment_date = :payment_date,
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedLoan)
	if err != nil {
		return err
	}

	return nil
}

func NewLoanRepositoryPostgres(configuration *config.DBConfig) LoanRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS loans (
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				customer_id VARCHAR(36) NOT NULL,
				sum DOUBLE PRECISION NOT NULL,
				paid_sum DOUBLE PRECISION NOT NULL,
				percent INT NOT NULL,
				payment_date TIMESTAMP NOT NULL,
				end_date TIMESTAMP NOT NULL
			  );`

	_, err := db.Exec(query)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating loans db: %s", err.Error()))
	}
	return &LoanRepositoryPostgres{db: db}
}
