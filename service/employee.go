package service

import (
	"github.com/tomintaiga/yandex_partice_1/models"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

const (
	EMPLOYEE_REPOSITORY_NAME = "employee_repository"
)

type EmployeeServiceConfig map[string]interface{}

type EmployeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(cfg EmployeeServiceConfig) (*EmployeeService, error) {
	return &EmployeeService{
		repo: cfg[EMPLOYEE_REPOSITORY_NAME].(repository.EmployeeRepository),
	}, nil
}

// GetEmployeeList retrieve employee list from repository
func (srv *EmployeeService) GetEmployeeList() ([]models.Employee, error) {
	return srv.repo.GetEmployeeList()
}

// GetEmployeeInfo retrieve employee info rom repository
func (srv *EmployeeService) GetEmployeeInfo(login string) (models.Employee, error) {
	return srv.repo.GetEployee(login)
}

// Register add new employee to repository
func (srv *EmployeeService) Register(login string) (models.Employee, error) {
	employee := models.Employee{Login: login}
	err := srv.repo.Register(&employee)

	return employee, err
}

// Delete remove employee from repository
func (srv *EmployeeService) Delete(login string) error {
	employee := models.Employee{Login: login}
	err := srv.repo.Delete(&employee)

	return err
}
