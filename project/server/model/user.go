package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	//	UserId int `json:"userId",gorm:"index:idx_vehicle_id,primaryKey;autoIncrement:true"`
	FirstName string    `json:"firstname"gorm:"not null"`
	LastName  string    `json:"lastname",gorm:"not null"`
	Email     string    `gorm:"unique;not null",json:"email"`
	Password  string    `json:"password",gorm:"not null"`
	Bookings  []Booking `gorm:"foreignKey:UserId;references:ID"`
}
