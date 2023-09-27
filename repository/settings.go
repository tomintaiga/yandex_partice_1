package repository

import "github.com/tomintaiga/yandex_partice_1/domain"

type SettingsRepository interface {
	// SetGlobalBookingLimit set global booking limit for employee
	SetGlobalBookingLimit(limit uint32) error

	// SetNotifyOptions set global notification options
	SetNotifyOptions(domain.NotifyOptions) error
}
