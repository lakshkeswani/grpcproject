package model

import "gorm.io/gorm"

type Vehicle struct {
	gorm.Model
	//	VehicleId int `json:"vehicleId",gorm:"primaryKey;not null;autoIncrement:true;index:idx_vehicle_id,unique"`
	VehicleType    string         `json:"vehicleType"`
	Capacity       int64          `json:"capacity"`
	RegistrationNo string         `gorm:" not null;autoIncrement:true"`
	BusRouteStops  []BusRouteStop `gorm:"foreignKey:VehicleId;references:ID"`
}
