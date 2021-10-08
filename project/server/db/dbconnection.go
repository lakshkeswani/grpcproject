package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

var DB *gorm.DB
var err error

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(resources.DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		resources.Log.Panic("Cannot connect to DB")
	}
	//DB.Migrator().CreateConstraint(&User{}, "CreditCard")
	//DB.Migrator().CreateConstraint(&User{}, "fk_users_credit_card")
	//DB.Migrator().CreateIndex(&User{}, "fk_users_credit_cards")
	DB.Migrator().CreateConstraint(&model.Route{}, "BusRouteStop")
	DB.Migrator().CreateConstraint(&model.Vehicle{}, "BusRouteStop")
	DB.Migrator().CreateConstraint(&model.Station{}, "BusRouteStop")
	DB.Migrator().CreateConstraint(&model.Driver{}, "Trip")
	DB.Migrator().CreateConstraint(&model.BusRouteStop{}, "Trip")
	DB.Migrator().CreateConstraint(&model.User{}, "Booking")
	DB.Migrator().CreateConstraint(&model.Trip{}, "Booking")
	DB.Migrator().CreateConstraint(&model.Station{}, "Booking")
	resources.Log.Info(" Constraint Created  ")

	DB.AutoMigrate(&model.User{}, &model.Driver{}, &model.Vehicle{}, &model.Station{}, &model.Route{}, &model.BusRouteStop{}, &model.Trip{}, &model.Booking{})
	resources.Log.Info(" AutoMigrate successful ")

}
