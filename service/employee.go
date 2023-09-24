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

// NewEmployeeService will create and initialize EmployeeService
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

// SetBookingLimit update eployee booking limit
func (srv *EmployeeService) SetBookingLimit(login string, limit uint32) error {
	employee, err := srv.repo.GetEployee(login)
	if err != nil {
		return err
	}

	employee.MonthLimit = limit
	return srv.repo.UpdateEmployee(&employee)
}

// GetEmployeeBalance get balance for selected GetEmployeeBalance
func (srv *EmployeeService) GetEmployeeBalance(login string) (uint32, error) {
	employee, err := srv.repo.GetEployee(login)
	if err != nil {
		return 0, err
	}

	return employee.Balance, nil
}

// AddBooking will add booking ID to employee booking history
func (srv *EmployeeService) AddBooking(login string, booking_id string) error {
	employee, err := srv.repo.GetEployee(login)
	if err != nil {
		return err
	}

	employee.Bookings = append(employee.Bookings, booking_id)
	return srv.repo.UpdateEmployee(&employee)
}
