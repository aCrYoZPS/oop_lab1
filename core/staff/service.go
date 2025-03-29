package staff

import "oopLab1/config"

type StaffMemberService struct {
	repos StaffMemberRepository
}

func NewStaffMemberService(config *config.DBConfig) *StaffMemberService {
	return &StaffMemberService{repos: NewStaffMemberRepositoryPostgres(config)}
}

func (s *StaffMemberService) CreateStaffMember(StaffMember *StaffMember) error {
	return s.repos.Save(StaffMember)
}

func (s *StaffMemberService) GetStaffMemberByID(id string) (*StaffMember, error) {
	return s.repos.GetById(id)
}

func (s *StaffMemberService) GetStaffMemberByEmail(email string) (*StaffMember, error) {
	return s.repos.GetByEmail(email)
}

func (s *StaffMemberService) GetAllStaffMembers() ([]StaffMember, error) {
	return s.repos.GetAll()
}

func (s *StaffMemberService) DeleteStaffMember(id string) error {
	return s.repos.DeleteById(id)
}

func (s *StaffMemberService) UpdateStaffMember(StaffMember *StaffMember) error {
	return s.repos.Update(StaffMember)
}
