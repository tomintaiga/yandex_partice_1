package store

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/domain"
)

type EmployeeStore struct{}

func (s *EmployeeStore) GetEmployeeList() ([]domain.Employee, error) {
	return make([]domain.Employee, 0), fmt.Errorf("not implemented")
}

func (s *EmployeeStore) GetEmployee(login string) (domain.Employee, error) {
	l := log.With().Str("func", "EmployeeStore.GetEmployee").Str("login", login).Logger()

	employee := domain.Employee{Login: login}
	tx := getDB().Find(&employee)

	if tx.Error != nil {
		l.Error().Err(tx.Error).Msg("Can't check if manager exist")
		return domain.Employee{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		l.Debug().Msg("Employee not found")
		return domain.Employee{}, fmt.Errorf("not found")
	}

	return employee, nil
}

func (s *EmployeeStore) Register(login string) (*domain.Employee, error) {

	employee := domain.Employee{Login: login}
	tx := getDB().Create(&employee)

	if tx.RowsAffected == 0 {
		return &domain.Employee{}, fmt.Errorf("can't create employee")
	}

	if tx.Error != nil {
		return &domain.Employee{}, tx.Error
	}

	return &employee, nil
}

func (s *EmployeeStore) Delete(login string) error {
	return fmt.Errorf("not implemented")
}

func (s *EmployeeStore) UpdateEmployee(employee domain.Employee) error {
	return fmt.Errorf("not implemented")
}
