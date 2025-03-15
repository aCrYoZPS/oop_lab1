package transactions

import "oopLab1/storage_postgres"

type TransactionRepository interface {
	Save(customer Transaction) error
	GetById(id string) (Transaction, error)
	GetAll() ([]Transaction, error)
	DeleteById(id string) error
}

type TransactionRepositoryPostgres struct {
}

func (repos *TransactionRepositoryPostgres) Save(transaction Transaction) error {
	return nil
}

func (repos *TransactionRepositoryPostgres) GetById(id string) (Transaction, error) {
	return Transaction{}, nil
}

func (repos *TransactionRepositoryPostgres) GetAll() ([]Transaction, error) {
	return nil, nil
}

func (repos *TransactionRepositoryPostgres) DeleteById(id string) error {
	return nil
}

func NewTransactionRepositoryPostgres() TransactionRepository {
	return &TransactionRepositoryPostgres{}
}
