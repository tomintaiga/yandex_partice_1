package store

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/domain"
	"gorm.io/gorm"
)

type EmployeeStore struct{}

func (s *EmployeeStore) GetEmployeeList() ([]domain.Employee, error) {
	return make([]domain.Employee, 0), fmt.Errorf("not implemented")
}

func (s *EmployeeStore) GetEmployee(login string) (domain.Employee, error) {
	l := log.With().Str("func", "EmployeeStore.GetEmployee").Str("login", login).Logger()

	emoloyee := domain.Employee{Login: login}
	err := getDB().Model(&domain.Employee{}).Preload("Bookings").First(&emoloyee).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.Error().Err(err).Msg("Can't get employee")
		return domain.Employee{}, err
	}

	if err != nil {
		l.Error().Err(err).Msg("Can't get manager")
		return domain.Employee{}, err
	}

	return emoloyee, nil
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
