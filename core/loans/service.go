package loans

import "oopLab1/config"

type LoanService struct {
	repos LoanRepository
}

func NewLoanService(config *config.DBConfig) *LoanService {
	return &LoanService{repos: NewLoanRepositoryPostgres(config)}
}

func (s *LoanService) CreateLoan(transaction *Loan) error {
	return s.repos.Save(transaction)
}

func (s *LoanService) GetLoanByID(id string) (*Loan, error) {
	return s.repos.GetById(id)
}

func (s *LoanService) GetAllLoans() ([]Loan, error) {
	return s.repos.GetAll()
}

func (s *LoanService) GetAllByOwner(owner_id string) ([]Loan, error) {
	return s.repos.GetAllByOwner(owner_id)
}

func (s *LoanService) DeleteLoan(id string) error {
	return s.repos.DeleteById(id)
}

func (s *LoanService) UpdateLoan(updatedLoan *Loan) error {
	return s.repos.Update(updatedLoan)
}
