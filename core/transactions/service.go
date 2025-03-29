package transactions

import "oopLab1/config"

type TransactionService struct {
	repos TransactionRepository
}

func NewTransactionService(config *config.DBConfig) *TransactionService {
	return &TransactionService{repos: NewTransactionRepositoryPostgres(config)}
}

func (s *TransactionService) CreateTransaction(account *Transaction) error {
	return s.repos.Save(account)
}

func (s *TransactionService) GetTransactionByID(id string) (*Transaction, error) {
	return s.repos.GetById(id)
}

func (s *TransactionService) GetAllOutboundTransactions(account_id string) ([]Transaction, error) {
	return s.repos.GetAllOutbound(account_id)
}

func (s *TransactionService) GetAllInboundTransactions(account_id string) ([]Transaction, error) {
	return s.repos.GetAllInbound(account_id)
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
