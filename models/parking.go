package models

import "time"

type Parking struct {
	Spots  []string
	Scheme string
}

type Manager struct {
	Login string
}

type Employee struct {
	Login      string
	Balance    uint32
	MonthLimit uint32
	Bookings   []string
}

type BookingStatus string

const (
	STATUS_CANCELED BookingStatus = "canceled"
	STATUS_BOOKED   BookingStatus = "booked"
)

type Booking struct {
	ID             string
	Spot           string
	Date           time.Time
	CarPlateNumber string
	Status         BookingStatus
}

type NotifyOptions struct {
	SendTime      time.Time
	ReceiverEmail string
	EmailTemplate string
}
