package store

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/tomintaiga/yandex_partice_1/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

const (
	db_filename = "parking.db"
)

func initialize() {
	l := log.With().Str("func", "db.initialize").Logger()

	var err error
	db, err = gorm.Open(sqlite.Open(db_filename), &gorm.Config{})
	if err != nil {
		l.Error().Err(err).Msg("Can't open database")
		panic(err)
	}

	err = db.AutoMigrate(
		&domain.Booking{},
		&domain.Employee{},
		&domain.Manager{},
		&domain.ParkingSpot{},
		&domain.Parking{},
	)

	if err != nil {
		l.Error().Err(err).Msg("Can't migrate database")
		panic(err)
	}
}

// getDB singleton for DB
func getDB() *gorm.DB {
	once.Do(initialize)
	return db.Debug()
}
