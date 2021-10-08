package repository

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/db"
)

func GetAllBookingsByUser(id interface{}, value interface{}) error {

	query := "SELECT bookings.id, bookings.trip_id, bookings.user_id, bookings.pickup_time, bookings.drop_off_time, bookings.pickup_stop_id, bookings.drop_off_stop_id, bookings.fare, status FROM users inner join bookings on users.id=bookings.user_id where users.id=?"
	err := db.DB.Raw(query, id).Scan(value).Error
	if err != nil {
		resources.Log.Warn("error %v", err)
		return err
	}
	resources.Log.Info("GetAllBookingsByUser got data ")
	return nil
}
