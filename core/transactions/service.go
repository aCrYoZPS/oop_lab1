package transactions

import "oopLab1/config"

type TransactionService struct {
	repos TransactionRepository
}

func NewTransactionService(config *config.DBConfig) *TransactionService {
	return &TransactionService{repos: NewTransactionRepositoryPostgres(config)}
}

func (s *TransactionService) CreateTransaction(transaction *Transaction) error {
	return s.repos.Save(transaction)
}

func (s *TransactionService) GetTransactionByID(id string) (*Transaction, error) {
	return s.repos.GetById(id)
}

func (s *TransactionService) GetAllOutboundTransactions(transaction_id string) ([]Transaction, error) {
	return s.repos.GetAllOutbound(transaction_id)
}

func (s *TransactionService) GetAllInboundTransactions(transaction_id string) ([]Transaction, error) {
	return s.repos.GetAllInbound(transaction_id)
}

func (s *TransactionService) GetAllTransactions() ([]Transaction, error) {
	return s.repos.GetAll()
}

func (s *TransactionService) GetAllByActor(actor_id string) ([]Transaction, error) {
	return s.repos.GetAllByActor(actor_id)
}

func (s *TransactionService) DeleteTransaction(id string) error {
	return s.repos.DeleteById(id)
}
