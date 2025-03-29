package company

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CompanyRepository interface {
	Save(company *Company) error
	GetById(id string) (*Company, error)
	GetByEmail(email string) (*Company, error)
	GetAll() ([]Company, error)
	DeleteById(id string) error
	Update(updatedCompany *Company) error
}

type CompanyRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *CompanyRepositoryPostgres) Save(company *Company) error {
	query := `INSERT INTO companies (id, name, phone_number, email, password, country, type, bic, address, anp, access_allowed)
			  VALUES (:id, :name, :phone_number, :email, :password, :country, :type, :bic, :address, :anp, :access_allowed)`

	_, err := repos.db.NamedExec(query, company)
	if err != nil {
		return err
	}

	return nil
}

func (repos *CompanyRepositoryPostgres) GetById(id string) (*Company, error) {
	var company *Company = new(Company)

	query := `SELECT id, name, phone_number, email, password, country, type, bic, address, anp, access_allowed 
              FROM companies WHERE id = $1`

	err := repos.db.Get(company, query, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (repos *CompanyRepositoryPostgres) GetByEmail(email string) (*Company, error) {
	var company *Company = new(Company)

	query := `SELECT id, name, phone_number, email, password, country, type, bic, address, anp, access_allowed 
              FROM companies WHERE email = $1`

	err := repos.db.Get(company, query, email)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (repos *CompanyRepositoryPostgres) GetAll() ([]Company, error) {
	var companies []Company

	query := "SELECT * FROM companies"

	err := repos.db.Select(&companies, query)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (repos *CompanyRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM companies WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *CompanyRepositoryPostgres) Update(updatedCompany *Company) error {
	query := `UPDATE companies
				SET name = :name,
					phone_number = :phone_number,
					email = :email,
					password = :password,
					country = :country,
					type = :type,
					bic = :bic,
					address = :address,
					anp = :anp,
					access_allowed = :access_allowed
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedCompany)

	if err != nil {
		return err
	}

	return nil
}

func NewCompanyRepositoryPostgres(configuration *config.DBConfig) CompanyRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS companies(
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				name VARCHAR(255) NOT NULL,
				phone_number VARCHAR(20),
				email VARCHAR(255) UNIQUE NOT NULL,
				password VARCHAR(255) NOT NULL,
				country VARCHAR(100) NOT NULL,
				type VARCHAR(32) NOT NULL,
				bic VARCHAR(36) NOT NULL,
				address VARCHAR(255) NOT NULL,
				anp VARCHAR(36) NOT NULL,
				access_allowed BOOLEAN DEFAULT FALSE
			  );`

	_, err := db.Exec(query)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating company db: %s", err.Error()))
	}
	return &CompanyRepositoryPostgres{db: db}
}
