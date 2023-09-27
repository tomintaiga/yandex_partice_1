package repository

import "github.com/tomintaiga/yandex_partice_1/domain"

type EmployeeRepository interface {
	// GetEmployeeList get all employees from store
	// TODO: Add pagination
	GetEmployeeList() ([]domain.Employee, error)

	// GetEployee will return eployee by it's login
	GetEmployee(login string) (domain.Employee, error)

	// Register add new employee to store. After this, employee object will have valid ID
	// Employee loging must be unique, or error must be returned
	Register(login string) (domain.Employee, error)

	// Delete will remove employee from stopre
	// If emploee not found, error must be returned
	Delete(login string) error

	// UpdateEmployee will update employee data
	// If emploee not found, error must be returned
	UpdateEmployee(employee domain.Employee) error
}
