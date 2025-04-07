package salaryprojects

type SalaryProject struct {
	ID        string  `db:"id" json:"id,omitempty"`
	Salary    float64 `db:"salary" json:"salary,omitempty"`
	WorkerID  string  `db:"worker_id" json:"worker_id,omitempty"`
	CompanyID string  `db:"company_id" json:"company_id,omitempty"`
}

func UpdateSalaryProjectInfo(original *SalaryProject, updated *SalaryProject) {
	updated.ID = original.ID
	updated.WorkerID = original.WorkerID
	updated.CompanyID = original.CompanyID
}
