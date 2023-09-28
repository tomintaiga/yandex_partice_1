package store

import (
	"sync"

	"github.com/tomintaiga/yandex_partice_1/models"
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
	var err error
	db, err = gorm.Open(sqlite.Open(db_filename), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(
		&models.Booking{},
		&models.Employee{},
		&models.Manager{},
		&models.ParkingSpot{},
		&models.Parking{},
	)

	if err != nil {
		panic(err)
	}
}

// getDB singleton for DB
func getDB() *gorm.DB {
	once.Do(initialize)
	return db.Debug()
}
