package repository

import "github.com/tomintaiga/yandex_partice_1/models"

type EmployeeRepository interface {
	// GetEmployeeList get all employees from store
	// TODO: Add pagination
	GetEmployeeList() ([]models.Employee, error)

	// GetEployee will return eployee by it's login
	GetEmployee(login string) (models.Employee, error)

	// Register add new employee to store. After this, employee object will have valid ID
	// Employee loging must be unique, or error must be returned
	Register(login string) (models.Employee, error)

	// Delete will remove employee from stopre
	// If emploee not found, error must be returned
	Delete(login string) error

	// UpdateEmployee will update employee data
	// If emploee not found, error must be returned
	UpdateEmployee(employee models.Employee) error
}
