package company

import "oopLab1/config"

type CompanyService struct {
	repos CompanyRepository
}

func NewCompanyService(config *config.DBConfig) *CompanyService {
	return &CompanyService{repos: NewCompanyRepositoryPostgres(config)}
}

func (s *CompanyService) CreateCompany(company *Company) error {
	return s.repos.Save(company)
}

func (s *CompanyService) GetCompanyByID(id string) (*Company, error) {
	return s.repos.GetById(id)
}

func (s *CompanyService) GetCompanyByEmail(email string) (*Company, error) {
	return s.repos.GetByEmail(email)
}

func (s *CompanyService) GetAllCompanies() ([]Company, error) {
	return s.repos.GetAll()
}

func (s *CompanyService) DeleteCompany(id string) error {
	return s.repos.DeleteById(id)
}

func (s *CompanyService) UpdateCompany(company *Company) error {
	return s.repos.Update(company)
}
