package storage

import "oopLab1/entities"

type BankStorage interface {
	GetBanks() ([]entities.Bank, error)
	GetBankByID(string) (entities.Bank, error)
	SaveBank(entities.Bank) error
	DeleteBank(string) error
}

type CustomerStorage interface {
	GetCustomers() ([]entities.Customer, error)
	GetCustomerByID(string) (entities.Customer, error)
	SaveCustomer(entities.Customer) error
	DeleteCustomer(string) error
}

type CompanyStorage interface {
	GetCompanies() ([]entities.Company, error)
	GetCompanyByID(string) (entities.Company, error)
	SaveCompany(entities.Company) error
	DeleteCompany(string) error
}

type AccountStorage interface {
	GetAccountByID(string) (entities.Account, error)
	GetAccountsByOwnerID(string) ([]entities.Account, error)
	SaveAccount(entities.Account) error
	DeleteAccount(entities.Account) error
}

type TransactionStorage interface {
	GetLastTransactions(string) ([]entities.Transaction, error)
	SaveTransaction(entities.Transaction) error
	DeleteTransaction(string) error
}

type LoanStorage interface {
	GetLoanByID(string) (entities.Loan, error)
	GetLoansByOwnerID(string) ([]entities.Loan, error)
	SaveLoan(entities.Loan) error
	DeleteLoan(string) error
}

type InstallmentStorage interface {
	GetInstallmentByID(string) (entities.Loan, error)
	GetInstallmentsByOwnerID(string) ([]entities.Loan, error)
	SaveInstallment(entities.Installment) error
	DeleteInstallment(string) error
}

type RoleStorage interface {
	GetRoles() ([]entities.Role, error)
}
