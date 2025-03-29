package bank

import "oopLab1/config"

type BankService struct {
	repos BankRepository
}

func NewBankService(config *config.DBConfig) *BankService {
	return &BankService{repos: NewBankRepositoryPostgres(config)}
}

func (s *BankService) CreateBank(bank *Bank) error {
	return s.repos.Save(bank)
}

func (s *BankService) GetBankByID(id string) (*Bank, error) {
	return s.repos.GetById(id)
}

func (s *BankService) GetAllBanks() ([]Bank, error) {
	return s.repos.GetAll()
}

func (s *BankService) DeleteBank(id string) error {
	return s.repos.DeleteById(id)
}

func (s *BankService) UpdateBank(bank *Bank) error {
	return s.repos.Update(bank)
}
