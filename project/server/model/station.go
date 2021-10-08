package model

import "gorm.io/gorm"

type Station struct {
	gorm.Model
	StationName   string         `json:"StationName"`
	BusRouteStops []BusRouteStop `gorm:"foreignKey:StationId;references:ID"`
	Bookings      []Booking      `gorm:"foreignKey:PickupStopId;references:ID"`
}
