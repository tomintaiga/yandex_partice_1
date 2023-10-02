package domain

import (
	"time"

	"gorm.io/gorm"
)

const (
	STATUS_CANCELED string = "canceled"
	STATUS_BOOKED   string = "booked"
)

type Booking struct {
	gorm.Model
	// ID             string `gorm:"primarykey;not null;unique"`
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
	Bookings   []Booking `gorm:"foreignKey:ID"`
}

type Manager struct {
	gorm.Model
	Login     string     `gorm:"not null"`
	Employees []Employee `gorm:"foreignKey:ID"`
}

type ParkingSpot struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Parking struct {
	gorm.Model
	Spots    []ParkingSpot `gorm:"foreignKey:ID"`
	Managers []Manager     `gorm:"foreignKey:ID"`
	Scheme   string
}

type NotifyOptions struct {
	SendTime      time.Time
	ReceiverEmail string
	EmailTemplate string
}
