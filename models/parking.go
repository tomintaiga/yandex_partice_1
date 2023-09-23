package models

import "time"

type Parking struct {
	ID        uint32
	Managers  []Manager
	Spots     []ParkingSpot
	Employees []Employee
	Scheme    string
}

type Manager struct {
	ID    uint32
	Login string `gorm:"not null"`
}

type Employee struct {
	ID         uint32
	Login      string `gorm:"not null"`
	Balance    uint32
	MonthLimit uint32
	Bookings   []Booking
}

type ParkingSpot struct {
	ID string `gorm:"not null"`
}

type BookingStatus string

const (
	STATUS_CANCELED BookingStatus = "canceled"
)

type Booking struct {
	ID             string `gorm:"not null"`
	Spot           ParkingSpot
	Date           time.Time
	CarPlateNumber string        `gorm:"not null"`
	Status         BookingStatus // See: https://stackoverflow.com/questions/68637265/how-can-i-add-enum-in-gorm
}
