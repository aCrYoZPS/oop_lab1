package customer

import "oopLab1/config"

type CustomerService struct {
	repos CustomerRepository
}

func NewCustomerService(config *config.DBConfig) *CustomerService {
	return &CustomerService{repos: NewCustomerRepositoryPostgres(config)}
}

func (s *CustomerService) CreateCustomer(customer *Customer) error {
	return s.repos.Save(customer)
}

func (s *CustomerService) GetCustomerByID(id string) (*Customer, error) {
	return s.repos.GetById(id)
}

func (s *CustomerService) GetCustomerByEmail(email string) (*Customer, error) {
	return s.repos.GetByEmail(email)
}

func (s *CustomerService) GetAllCustomers() ([]Customer, error) {
	return s.repos.GetAll()
}

func (s *CustomerService) DeleteCustomer(id string) error {
	return s.repos.DeleteById(id)
}

func (s *CustomerService) UpdateCustomer(customer *Customer) error {
	return s.repos.Update(customer)
}
