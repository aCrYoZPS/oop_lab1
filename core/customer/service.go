package customer

type CustomerService struct {
	repos CustomerRepository
}

func NewUserService(repos CustomerRepository) *CustomerService {
	return &CustomerService{repos: repos}
}
