package repository

import "github.com/tomintaiga/yandex_partice_1/models"

type ManagerRepository interface {
	// Register add new manager to store. After this, manager object will have valid ID
	// Manager login must be unique, or error must be created
	Register(manager *models.Manager) error
}
