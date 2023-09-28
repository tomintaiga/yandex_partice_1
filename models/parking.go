package models

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	ID             string `gorm:"not null;unique"`
	Spot           string `gorm:"not null"`
	Date           time.Time
	CarPlateNumber string `gorm:"not null"`
	Status         string `gorm:"not null"`
}

type Employee struct {
	gorm.Model
	Login      string `gorm:"not null"`
	Balance    uint32
	MonthLimit uint32
	Bookings   []Booking `gorm:"many2many:employee_booking"`
}

type Manager struct {
	gorm.Model
	Login     string     `gorm:"not null"`
	Employees []Employee `gorm:"many2many:manager_employee"`
}

type ParkingSpot struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Parking struct {
	gorm.Model
	Spots    []ParkingSpot `gorm:"many2many:parking_spots"`
	Managers []Manager     `gorm:"many2many:parking_manager"`
	Scheme   string
}
