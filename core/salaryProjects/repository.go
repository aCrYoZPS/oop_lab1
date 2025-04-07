package salaryprojects

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SalaryProjectRepository interface {
	Save(salaryProject *SalaryProject) error
	GetById(id string) (*SalaryProject, error)
	GetAll() ([]SalaryProject, error)
	GetAllByCompany(company_id string) ([]SalaryProject, error)
	DeleteById(id string) error
	Update(updatedSalaryProject *SalaryProject) error
}

type SalaryProjectRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *SalaryProjectRepositoryPostgres) Save(salaryProject *SalaryProject) error {
	query := `INSERT INTO salary_projects (id, salary, worker_id, company_id)
			  VALUES (:id, :salary, :worker_id, :company_id)`

	_, err := repos.db.NamedExec(query, salaryProject)
	if err != nil {
		return err
	}

	return nil
}

func (repos *SalaryProjectRepositoryPostgres) GetById(id string) (*SalaryProject, error) {
	var salaryProject *SalaryProject = new(SalaryProject)

	query := `SELECT id, salary, worker_id, company_id
              FROM salary_projects WHERE id = $1`

	err := repos.db.Get(salaryProject, query, id)
	if err != nil {
		return nil, err
	}

	return salaryProject, nil
}

func (repos *SalaryProjectRepositoryPostgres) GetAll() ([]SalaryProject, error) {
	var salaryProjects []SalaryProject

	query := "SELECT * FROM salary_projects"

	err := repos.db.Select(&salaryProjects, query)
	if err != nil {
		return nil, err
	}

	return salaryProjects, nil
}

func (repos *SalaryProjectRepositoryPostgres) GetAllByCompany(company_id string) ([]SalaryProject, error) {
	var salaryProjects []SalaryProject

	query := "SELECT * FROM salary_projects WHERE company_id = $1"

	err := repos.db.Select(&salaryProjects, query, company_id)
	if err != nil {
		return nil, err
	}

	return salaryProjects, nil
}

func (repos *SalaryProjectRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM salary_projects WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *SalaryProjectRepositoryPostgres) Update(updatedSalaryProject *SalaryProject) error {
	query := `UPDATE salary_projects 
				SET salary = :salary,
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedSalaryProject)
	if err != nil {
		return err
	}

	return nil
}

func NewSalaryProjectRepositoryPostgres(configuration *config.DBConfig) SalaryProjectRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS salary_projects(
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				salary DOUBLE PRECISION NOT NULL,
				worker_id VARCHAR(36) NOT NULL,
				company_id VARCHAR(36) NOT NULL
			  );`

	_, err := db.Exec(query)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating salary_projects db: %s", err.Error()))
	}

	return &SalaryProjectRepositoryPostgres{db: db}
}
