package model

import "gorm.io/gorm"

type Route struct {
	gorm.Model

	RouteName     string         `json:"RouteName"`
	BusRouteStops []BusRouteStop `gorm:"foreignKey:RouteID;references:ID"`
}
