package store

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/domain"
	"github.com/tomintaiga/yandex_partice_1/models"
)

type ManagerStore struct{}

func (s *ManagerStore) Register(login string) (domain.Manager, error) {
	l := log.With().Str("func", "ManagerStore.Register").Str("login", login).Logger()

	manager := models.Manager{Login: login}
	tx := getDB().Find(&manager)

	if tx.Error != nil {
		l.Error().Err(tx.Error).Msg("Can't check if manager exist")
		return domain.Manager{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		l.Debug().Msg("Manager not found")

		r := getDB().Create(&manager)
		if r.Error != nil {
			l.Error().Err(r.Error).Msg("Can't create")
		}

		l.Debug().Int("id", int(manager.ID)).Msg("Created")
		return domain.Manager{Login: manager.Login}, nil
	}

	l.Error().Msg("Already created")
	return domain.Manager{}, fmt.Errorf("manager is already created")
}
