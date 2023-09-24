package service

import (
	"time"

	"github.com/tomintaiga/yandex_partice_1/models"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

const (
	SETTINGS_REPOSITORY_NAME = "settings_repository"
)

type SettingsServiceConfig map[string]interface{}

type SettingsService struct {
	repo repository.SettingsRepository
}

// NewSettingsService create and initialize SettingsService
func NewSettingsService(cfg ParkingServiceConfig) (*SettingsService, error) {
	return &SettingsService{
		repo: cfg[SETTINGS_REPOSITORY_NAME].(repository.SettingsRepository),
	}, nil
}

// SetGlobalBookingLimit set global bookink limit
func (srv *SettingsService) SetGlobalBookingLimit(limit uint32) error {
	return srv.repo.SetGlobalBookingLimit(limit)
}

// SetNotificationOptions update global notify options
func (srv *SettingsService) SetNotificationOptions(send_time time.Time, email string, template string) error {
	options := models.NotifyOptions{
		SendTime:      send_time,
		ReceiverEmail: email,
		EmailTemplate: template,
	}

	return srv.repo.SetNotifyOptions(options)
}
