package model

import (
	"gorm.io/gorm"
	"time"
)

type BusRouteStop struct {
	gorm.Model
	RouteID   int       `gorm:"primaryKey;not null;autoIncrement:false"`
	VehicleId int       `gorm:"primaryKey; not null;autoIncrement:false"`
	StationId int       `gorm:"primaryKey; not null;autoIncrement:false"`
	StartTime time.Time `json:"StartTime";gorm:"type:time"`
	EndTime   time.Time `json:"EndTime";gorm:"type:time"`
	StopNum   int
	Trips     []Trip `gorm:"foreignKey:BusRouteStopId;references:ID"`
}
