package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
	"grpcproject/project/server/repository"
)

func CreateBooking(Booking model.Booking) (string, error) {
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.Save(&Booking)
	if err != nil {
		resources.Log.Warn(" err at service CreateBooking called", err)

		return "unable to save", err
	}
	return "booking created", nil
}
func ReadBooking(BookingId int) (model.Booking, error) {
	var Booking model.Booking
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.FindById(BookingId, &Booking)
	if err != nil {
		resources.Log.Warn(" err at service ReadBooking called", err)
		return Booking, err
	}
	return Booking, err

}
func UpdateBooking(id int, Booking model.Booking) (model.Booking, error) {
	var booking model.Booking
	err := genRepository.FindById(id, &booking)
	if err != nil {
		resources.Log.Warn(" err at service UpdateBooking called", err)
		return booking, err
	}
	err = genRepository.Save(&Booking)
	if err != nil {
		resources.Log.Warn(" err at service UpdateBooking called", err)
		return booking, err
	}
	return Booking, err
}
func DeleteBooking(id interface{}) error {
	err := genRepository.Delete(id, &model.Booking{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteBooking called", err)
		return err
	}
	return nil

}
func GetAllBookingsByUser(user_id int, value interface{}) {
	err := repository.GetAllBookingsByUser(user_id, value)
	if err != nil {
		resources.Log.Warn(" err at service GetAllBookingsByUser called", err)

	}
}
