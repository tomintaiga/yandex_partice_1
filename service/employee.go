package service

import (
	"github.com/tomintaiga/yandex_partice_1/domain"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepository
}

// NewEmployeeService will create and initialize EmployeeService
func NewEmployeeService(repo repository.EmployeeRepository) (*EmployeeService, error) {
	return &EmployeeService{
		repo: repo,
	}, nil
}

// GetEmployeeList retrieve employee list from repository
func (srv *EmployeeService) GetEmployeeList() ([]domain.Employee, error) {
	return srv.repo.GetEmployeeList()
}

// GetEmployeeInfo retrieve employee info rom repository
func (srv *EmployeeService) GetEmployeeInfo(login string) (domain.Employee, error) {
	return srv.repo.GetEmployee(login)
}

// Register add new employee to repository
func (srv *EmployeeService) Register(login string) (domain.Employee, error) {
	return srv.repo.Register(login)
}

// Delete remove employee from repository
func (srv *EmployeeService) Delete(login string) error {
	return srv.repo.Delete(login)
}

// SetBookingLimit update eployee booking limit
func (srv *EmployeeService) SetBookingLimit(login string, limit uint32) error {
	employee, err := srv.repo.GetEmployee(login)
	if err != nil {
		return err
	}

	employee.MonthLimit = limit
	return srv.repo.UpdateEmployee(employee)
}

// GetEmployeeBalance get balance for selected GetEmployeeBalance
func (srv *EmployeeService) GetEmployeeBalance(login string) (uint32, error) {
	employee, err := srv.repo.GetEmployee(login)
	if err != nil {
		return 0, err
	}

	return employee.Balance, nil
}

// AddBooking will add booking ID to employee booking history
func (srv *EmployeeService) AddBooking(login string, booking_id string) error {
	employee, err := srv.repo.GetEmployee(login)
	if err != nil {
		return err
	}

	employee.Bookings = append(employee.Bookings, booking_id)
	return srv.repo.UpdateEmployee(employee)
}
