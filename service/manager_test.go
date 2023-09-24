package service

import (
	"testing"

	"github.com/tomintaiga/yandex_partice_1/models"
)

type managerTestRepo struct{}

func (srv *managerTestRepo) Register(login string) (models.Manager, error) {
	return models.Manager{Login: login}, nil
}

func TestRegisterManager(t *testing.T) {
	secret := "abc"
	login := "Dima"
	srv, err := NewManagerService(secret, &managerTestRepo{})
	if err != nil {
		t.Errorf("Can't init manager service")
	}

	manager, err := srv.Register(secret, login)
	if err != nil {
		t.Errorf("Can't registry manager: %v", err)
	}

	if manager.Login != login {
		t.Errorf("Bad manager login: %v", manager.Login)
	}
}

func TestRegisterManagerError(t *testing.T) {
	secret := "abc"
	login := "Dima"
	srv, err := NewManagerService(secret, &managerTestRepo{})
	if err != nil {
		t.Errorf("Can't init manager service")
	}

	_, err = srv.Register(login, login)
	if err == nil {
		t.Errorf("Can register manager with invalid code")
	}
}
