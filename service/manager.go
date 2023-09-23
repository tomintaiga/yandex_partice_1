package service

import (
	"fmt"

	"github.com/tomintaiga/yandex_partice_1/models"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

const (
	SECRET_CODE_NAME        = "secret_code"
	MANAGER_REPOSITORY_NAME = "manager_repository"
)

type ManagerServiceConfig map[string]interface{}

type ManagerService struct {
	secret_code string
	repo        repository.ManagerRepository
}

// NewManagerService will create and initialize Manager service
func NewManagerService(cfg ManagerServiceConfig) (*ManagerService, error) {
	return &ManagerService{
		secret_code: cfg[SECRET_CODE_NAME].(string),
		repo:        cfg[MANAGER_REPOSITORY_NAME].(repository.ManagerRepository),
	}, nil
}

// Register add new manager to repository
func (srv *ManagerService) Register(secret_code string, login string) (models.Manager, error) {
	if secret_code != srv.secret_code {
		return models.Manager{}, fmt.Errorf("bad code")
	}

	manager := models.Manager{Login: login}
	err := srv.repo.Register(&manager)

	return manager, err
}
