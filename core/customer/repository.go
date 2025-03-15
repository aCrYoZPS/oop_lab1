package customer

type CustomerRepository interface {
	Save(customer Customer) error
	GetById(id string) (Customer, error)
	GetAll() ([]Customer, error)
	DeleteById(id string) error
}

type CustomerRepositoryPostgres struct{}

func (repos *CustomerRepositoryPostgres) Save(customer Customer) error {
	return nil
}

func (repos *CustomerRepositoryPostgres) GetById(id string) (Customer, error) {
	return Customer{}, nil
}

func (repos *CustomerRepositoryPostgres) GetAll() ([]Customer, error) {
	return nil, nil
}

func (repos *CustomerRepositoryPostgres) DeleteById(id string) error {
	return nil
}

func NewCustomerRepositoryPostgres() CustomerRepository {
	return &CustomerRepositoryPostgres{}
}
