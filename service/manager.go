package service

import (
	"fmt"

	"github.com/tomintaiga/yandex_partice_1/models"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

type ManagerService struct {
	secret_code string
	repo        repository.ManagerRepository
}

// NewManagerService will create and initialize Manager service
func NewManagerService(code string, repo repository.ManagerRepository) (*ManagerService, error) {
	return &ManagerService{
		secret_code: code,
		repo:        repo,
	}, nil
}

// Register add new manager to repository
func (srv *ManagerService) Register(secret_code string, login string) (models.Manager, error) {
	if secret_code != srv.secret_code {
		return models.Manager{}, fmt.Errorf("bad code")
	}

	return srv.repo.Register(login)
}
