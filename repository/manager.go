package repository

import "github.com/tomintaiga/yandex_partice_1/domain"

type ManagerRepository interface {
	// Register add new manager to store. After this, manager object will have valid ID
	// Manager login must be unique, or error must be returned
	Register(login string) (*domain.Manager, error)

	// Get manager by it's login
	Get(login string) (*domain.Manager, error)

	// Update manager date
	Update(manager *domain.Manager) (*domain.Manager, error)
}
