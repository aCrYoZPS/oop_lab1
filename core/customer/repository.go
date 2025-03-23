package customer

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepository interface {
	Save(customer *Customer) error
	GetById(id string) (*Customer, error)
	GetByEmail(email string) (*Customer, error)
	GetAll() ([]Customer, error)
	DeleteById(id string) error
	Update(updatedCustomer *Customer) error
}

type CustomerRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *CustomerRepositoryPostgres) Save(customer *Customer) error {
	query := `INSERT INTO customers (id, name, phone_number, email, password, country, passport_number, passport_id)
	          VALUES (:id, :name, :phone_number, :email, :password, :country, :passport_number, :passport_id)`

	_, err := repos.db.NamedExec(query, customer)
	if err != nil {
		return err
	}

	return nil
}

func (repos *CustomerRepositoryPostgres) GetById(id string) (*Customer, error) {
	var customer *Customer = new(Customer)

	query := `SELECT id, name, phone_number, email, password, country, passport_number, passport_id, access_allowed 
              FROM customers WHERE id = $1`

	err := repos.db.Get(customer, query, id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repos *CustomerRepositoryPostgres) GetByEmail(email string) (*Customer, error) {
	var customer *Customer = new(Customer)

	query := `SELECT id, name, phone_number, email, password, country, passport_number, passport_id, access_allowed 
              FROM customers WHERE email = $1`

	err := repos.db.Get(customer, query, email)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (repos *CustomerRepositoryPostgres) GetAll() ([]Customer, error) {
	var customers []Customer

	query := "SELECT * FROM customers"

	err := repos.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (repos *CustomerRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM customers WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *CustomerRepositoryPostgres) Update(updatedCusomer *Customer) error {
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

func NewCustomerRepositoryPostgres(configuration *config.DBConfig) CustomerRepository {
	db := storage_postgres.GetPostgresDB(configuration)

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
		logger.Fatal(fmt.Sprintf("Failed creating customer db: %s", err.Error()))
	}
	return &CustomerRepositoryPostgres{db: db}
}
