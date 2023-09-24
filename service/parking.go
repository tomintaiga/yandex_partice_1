package service

import (
	"fmt"
	"time"

	"github.com/tomintaiga/yandex_partice_1/models"
	"github.com/tomintaiga/yandex_partice_1/repository"
)

type ParkingService struct {
	repo repository.ParkingRepository
}

// NewParkingService create and initialize ParkingService
func NewParkingService(repo repository.ParkingRepository) (*ParkingService, error) {
	return &ParkingService{
		repo: repo,
	}, nil
}

// SetScheme set new parking scheme
func (srv *ParkingService) SetScheme(parking_id uint32, scheme []byte) error {
	return srv.repo.SetScheme(parking_id, scheme)
}

func (srv *ParkingService) GetScheme(parking_id uint32) ([]byte, error) {
	return srv.repo.GetScheme(parking_id)
}

// SetParkingSpots set new parking spots
func (srv *ParkingService) SetParkingSpots(parking_id uint32, spots []string) error {
	return srv.repo.SetParkingSpots(parking_id, spots)
}

// GetParkingSpots return parking spots for selected parking
func (srv *ParkingService) GetParkingSpots(parking_id uint32) ([]string, error) {
	parking, err := srv.repo.GetParkingById(parking_id)
	if err != nil {
		return []string{}, err
	}

	return parking.Spots, nil
}

// DeleteParkingSpot remove parking spot from parking
func (srv *ParkingService) DeleteParkingSpot(parking_id uint32, spot string) error {
	parking, err := srv.repo.GetParkingById(parking_id)
	if err != nil {
		return err
	}

	target_spots := make([]string, 0)
	for _, cur := range parking.Spots {
		if cur != spot {
			target_spots = append(target_spots, cur)
		}
	}

	return srv.repo.SetParkingSpots(parking_id, target_spots)
}

// GetAvailableParkingSlots return available parking slots for selected date
func (srv *ParkingService) GetAvailableParkingSlots(parking_id uint32, date time.Time) ([]string, error) {
	return srv.repo.GetParkingSpotsForDate(parking_id, date)
}

// BookSpot book selected spot for selected date and car
func (srv *ParkingService) BookSpot(parking_id uint32, spot string, car_plate string, date time.Time) (models.Booking, error) {

	// Check if we can book spot
	spots, err := srv.repo.GetParkingSpotsForDate(parking_id, date)
	if err != nil {
		return models.Booking{}, err
	}

	for _, cur_spot := range spots {
		if cur_spot == spot {
			return srv.repo.BookSlot(parking_id, spot, car_plate, date)
		}
	}

	return models.Booking{}, fmt.Errorf("slot occupied")
}

// GetBookings will extract booking info for every booking id provided
// If booking with ID not found, error will be returned
// TODO: Need optimization
func (srv *ParkingService) GetBookings(id []string) ([]models.Booking, error) {
	result := make([]models.Booking, 0)

	for _, cur := range id {
		booking, err := srv.repo.GetBookingById(cur)
		if err != nil {
			return []models.Booking{}, nil
		}

		result = append(result, booking)
	}

	return result, nil
}

func (srv *ParkingService) CancelBooking(id string) (models.Booking, error) {
	return srv.repo.CancelBooking(id)
}
