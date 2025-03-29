package staff

import (
	"fmt"
	"oopLab1/config"
	"oopLab1/pkg/logger"
	storage_postgres "oopLab1/storage/postgres"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type StaffMemberRepository interface {
	Save(staffMember *StaffMember) error
	GetById(id string) (*StaffMember, error)
	GetByEmail(email string) (*StaffMember, error)
	GetAll() ([]StaffMember, error)
	DeleteById(id string) error
	Update(updatedStaffMember *StaffMember) error
}

type StaffMemberRepositoryPostgres struct {
	db *sqlx.DB
}

func (repos *StaffMemberRepositoryPostgres) Save(staffMember *StaffMember) error {
	query := `INSERT INTO staff_members (id, name, phone_number, email, password, bank_id, role)
	          VALUES (:id, :name, :phone_number, :email, :password, :bank_id, :role)`

	_, err := repos.db.NamedExec(query, staffMember)

	if err != nil {
		return err
	}

	return nil
}

func (repos *StaffMemberRepositoryPostgres) GetById(id string) (*StaffMember, error) {
	var staffMember *StaffMember = new(StaffMember)

	query := `SELECT id, name, phone_number, email, password, bank_id, role 
              FROM staff_members WHERE id = $1`

	err := repos.db.Get(staffMember, query, id)

	if err != nil {
		return nil, err
	}

	return staffMember, nil
}

func (repos *StaffMemberRepositoryPostgres) GetByEmail(email string) (*StaffMember, error) {
	var staffMember *StaffMember = new(StaffMember)

	query := `SELECT id, name, phone_number, email, password, bank_id, role 
              FROM staff_members WHERE email = $1`

	err := repos.db.Get(staffMember, query, email)

	if err != nil {
		return nil, err
	}

	return staffMember, nil
}

func (repos *StaffMemberRepositoryPostgres) GetAll() ([]StaffMember, error) {
	var staffMembers []StaffMember

	query := "SELECT * FROM staff_members"

	err := repos.db.Select(&staffMembers, query)

	if err != nil {
		return nil, err
	}

	return staffMembers, nil
}

func (repos *StaffMemberRepositoryPostgres) DeleteById(id string) error {
	query := `DELETE FROM staff_members WHERE id = $1`

	_, err := repos.db.Exec(query, id)

	return err
}

func (repos *StaffMemberRepositoryPostgres) Update(updatedStaffMember *StaffMember) error {
	query := `UPDATE staff_members
				SET name = :name,
					phone_number = :phone_number,
					email = :email,
					password = :password,
					bank_id = :bank_id,
					role = role:
				WHERE id = :id`

	_, err := repos.db.NamedExec(query, updatedStaffMember)

	if err != nil {
		return err
	}

	return nil
}

func NewStaffMemberRepositoryPostgres(configuration *config.DBConfig) StaffMemberRepository {
	db := storage_postgres.GetPostgresDB(configuration)

	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto;
		      CREATE TABLE IF NOT EXISTS staff_members (
				id VARCHAR(36) PRIMARY KEY DEFAULT gen_random_uuid(),
				name VARCHAR(255) NOT NULL,
				phone_number VARCHAR(20),
				email VARCHAR(255) UNIQUE NOT NULL,
				password VARCHAR(255) NOT NULL,
				bank_id VARCHAR(100) NOT NULL,
				role VARCHAR(50) NOT NULL,
			  );`

	_, err := db.Exec(query)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed creating staffMember db: %s", err.Error()))
	}
	return &StaffMemberRepositoryPostgres{db: db}
}
