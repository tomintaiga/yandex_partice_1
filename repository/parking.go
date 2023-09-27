package repository

import (
	"time"

	"github.com/tomintaiga/yandex_partice_1/domain"
)

type ParkingRepository interface {
	// SetScheme will set schmete to selected parking
	SetScheme(parking_id uint32, scheme []byte) error

	// GetScheme get parking scheme
	GetScheme(parking_id uint32) ([]byte, error)

	// GetParkingById return parking object by it's id
	// If parking with such ID not found, error must be returned
	GetParkingById(id uint32) (*domain.Parking, error)

	// SetParkingSpots set list of parking spots for parking
	SetParkingSpots(parking_id uint32, spots []string) error

	// GetParkingSpotsForDate return list of available parking slots for selected date
	GetParkingSpotsForDate(parking_id uint32, date time.Time) ([]string, error)

	// BookSlot will reserve selected parking slot for selected day
	BookSlot(parking_id uint32, spot string, car_plate string, date time.Time) (domain.Booking, error)

	// GetBookingById will find booking with provided id
	GetBookingById(id string) (domain.Booking, error)

	// CancelBooking mark booking with provided ID as canceled and return it
	CancelBooking(id string) (domain.Booking, error)
}
