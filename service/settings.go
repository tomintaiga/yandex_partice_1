package service

import (
	"time"

	"github.com/tomintaiga/yandex_partice_1/domain"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

type SettingsService struct {
	repo repository.SettingsRepository
}

// NewSettingsService create and initialize SettingsService
func NewSettingsService(repo repository.SettingsRepository) (*SettingsService, error) {
	return &SettingsService{
		repo: repo,
	}, nil
}

// SetGlobalBookingLimit set global bookink limit
func (srv *SettingsService) SetGlobalBookingLimit(limit uint32) error {
	return srv.repo.SetGlobalBookingLimit(limit)
}

// SetNotificationOptions update global notify options
func (srv *SettingsService) SetNotificationOptions(send_time time.Time, email string, template string) error {
	options := domain.NotifyOptions{
		SendTime:      send_time,
		ReceiverEmail: email,
		EmailTemplate: template,
	}

	return srv.repo.SetNotifyOptions(options)
}