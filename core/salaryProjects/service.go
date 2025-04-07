package salaryprojects

import "oopLab1/config"

type SalaryProjectService struct {
	repos SalaryProjectRepository
}

func NewSalaryProjectService(config *config.DBConfig) *SalaryProjectService {
	return &SalaryProjectService{repos: NewSalaryProjectRepositoryPostgres(config)}
}

func (s *SalaryProjectService) CreateSalaryProject(SalaryProject *SalaryProject) error {
	return s.repos.Save(SalaryProject)
}

func (s *SalaryProjectService) GetSalaryProjectByID(id string) (*SalaryProject, error) {
	return s.repos.GetById(id)
}

func (s *SalaryProjectService) GetAllSalaryProjects() ([]SalaryProject, error) {
	return s.repos.GetAll()
}

func (s *SalaryProjectService) GetAllSalaryProjectsByCompany(company_id string) ([]SalaryProject, error) {
	return s.repos.GetAllByCompany(company_id)
}

func (s *SalaryProjectService) DeleteSalaryProject(id string) error {
	return s.repos.DeleteById(id)
}

func (s *SalaryProjectService) UpdateSalaryProject(SalaryProject *SalaryProject) error {
	return s.repos.Update(SalaryProject)
}
