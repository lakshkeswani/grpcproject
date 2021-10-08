package model

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	TripId        int       `gorm:"not null"`
	UserId        int       `gorm:"not null"`
	PickupTime    time.Time `json:"StartTime";gorm:"type:time"`
	DropOffTime   time.Time `json:"EndTime";gorm:"type:time"`
	PickupStopId  int       `gorm:"not null"`
	DropOffStopId int       `gorm:"not null"`
	Fare          int       `gorm:"not null"`
	Status        string    `gorm:"not null"`
}
