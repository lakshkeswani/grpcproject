package model

import (
	"gorm.io/gorm"
	"time"
)

type Trip struct {
	gorm.Model
	DriverId       int
	BusRouteStopId int
	Fare           int
	ScheduleDate   time.Time `gorm:"type:time"`
	Bookings       []Booking `gorm:"foreignKey:TripId;references:ID"`
}
