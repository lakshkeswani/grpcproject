package model

import "gorm.io/gorm"

type Driver struct {
	gorm.Model
	Name      string `json:"Name"`
	ContactNo string `json:"ContactNo"`
	LicenceID string `json:"LicenceId"`
	Email     string `gorm:"unique;not null",json:"email"`
	Password  string `json:"password",gorm:"not null"`
	Trips     []Trip `gorm:"foreignKey:DriverId;references:ID"`
}
