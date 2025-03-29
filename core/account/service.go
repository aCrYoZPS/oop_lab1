package account

import "oopLab1/config"

type AccountService struct {
	repos AccountRepository
}

func NewAccountService(config *config.DBConfig) *AccountService {
	return &AccountService{repos: NewAccountRepositoryPostgres(config)}
}

func (s *AccountService) CreateAccount(account *Account) error {
	return s.repos.Save(account)
}

func (s *AccountService) GetAccountByID(id string) (*Account, error) {
	return s.repos.GetById(id)
}

func (s *AccountService) GetAllAccountsByOwner(ownerID string) ([]Account, error) {
	return s.repos.GetAllByOwner(ownerID)
}

func (s *AccountService) GetAllAccounts() ([]Account, error) {
	return s.repos.GetAll()
}

func (s *AccountService) DeleteAccount(id string) error {
	return s.repos.DeleteById(id)
}

func (s *AccountService) UpdateAccount(account *Account) error {
	return s.repos.Update(account)
}
