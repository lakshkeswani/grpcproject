package model

import "gorm.io/gorm"

type XUser struct {
	gorm.Model
	MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard struct {
	gorm.Model
	Number     string `json:"number"`
	UserNumber string `gorm:"size:191"`
}
