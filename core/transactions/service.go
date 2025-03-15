package transactions

type TransactionService struct {
	repos TransactionRepository
}

func NewUserService(repos TransactionRepository) *TransactionService {
	return &TransactionService{repos: repos}
}
