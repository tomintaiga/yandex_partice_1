package store

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/domain"
	"gorm.io/gorm"
)

type ManagerStore struct{}

func (s *ManagerStore) Register(login string) (*domain.Manager, error) {
	l := log.With().Str("func", "ManagerStore.Register").Str("login", login).Logger()

	manager := domain.Manager{Login: login}
	tx := getDB().Find(&manager)

	if tx.Error != nil {
		l.Error().Err(tx.Error).Msg("Can't check if manager exist")
		return &domain.Manager{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		l.Debug().Msg("Manager not found")

		r := getDB().Create(&manager)
		if r.Error != nil {
			l.Error().Err(r.Error).Msg("Can't create")
		}

		l.Debug().Int("id", int(manager.ID)).Msg("Created")
		return &manager, nil
	}

	l.Error().Msg("Already created")
	return &domain.Manager{}, fmt.Errorf("manager is already created")
}

func (s *ManagerStore) Get(login string) (*domain.Manager, error) {
	l := log.With().Str("func", "ManagerStore.Get").Str("login", login).Logger()

	manager := domain.Manager{Login: login}
	err := getDB().Model(&domain.Manager{}).Preload("Employees").First(&manager).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.Error().Err(err).Msg("Can't get manager")
		return &domain.Manager{}, err
	}

	if err != nil {
		l.Error().Err(err).Msg("Can't get manager")
		return &domain.Manager{}, err
	}

	return &manager, nil
}

func (s *ManagerStore) Update(manager *domain.Manager) (*domain.Manager, error) {
	l := log.With().Str("func", "ManagerStore.Update").Str("login", manager.Login).Logger()

	tx := getDB().Save(manager)
	if tx.RowsAffected == 0 {
		l.Error().Msg("Can't update manager")
		if tx.Error != nil {
			l.Error().Err(tx.Error).Send()
		}

		return &domain.Manager{}, fmt.Errorf("can't update manager")
	}

	return manager, nil
}

func (s *ManagerStore) AddEmployee(manager *domain.Manager, employee *domain.Employee) (*domain.Manager, error) {
	l := log.With().Str("func", "ManagerStore.AddEmployee").Str("manager", manager.Login).Str("employee", employee.Login).Logger()
	err := getDB().Model(manager).Association("Employees").Append(employee)
	if err != nil {
		l.Error().Err(err).Msg("Can't add employee")
		return manager, err
	}

	l.Debug().Msg("Employee added")
	manager.Employees = append(manager.Employees, *employee)
	return manager, nil
}
