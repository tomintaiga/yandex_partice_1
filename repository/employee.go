package repository

import "github.com/tomintaiga/yandex_partice_1/models"

type EmployeeRepository interface {
	// GetEmployeeList get all employees from store
	// TODO: Add pagination
	GetEmployeeList() ([]models.Employee, error)

	// GetEployee will return eployee by it's login
	GetEployee(login string) (models.Employee, error)

	// Register add new employee to store. After this, employee object will have valid ID
	// Employee loging must be unique, or error must be created
	Register(employee *models.Employee) error

	// Delete will remove employee from stopre
	// If emploee not found, error must be created
	Delete(employee *models.Employee) error
}
