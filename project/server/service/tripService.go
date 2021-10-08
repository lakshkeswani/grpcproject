package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

func CreateTrip(Trip model.Trip) (string, error) {
	resources.Log.Info("service CreateVehicle called")
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.Save(&Trip)
	if err != nil {
		resources.Log.Warn(" err at service CreateTrip called", err)

		return "unable to save", err
	}
	return "trip created", nil
}
func ReadTrip(TripId int) (model.Trip, error) {
	resources.Log.Info("service ReadTrip called")

	var Trip model.Trip
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.FindById(TripId, &Trip)
	if err != nil {
		resources.Log.Warn(" err at service ReadTrip called", err)
		return Trip, err
	}
	return Trip, err
}
func UpdateTrip(id int, Trip model.Trip) (model.Trip, error) {
	resources.Log.Info("service UpdateTrip called")

	var trip model.Trip
	err := genRepository.FindById(id, &trip)
	if err != nil {
		resources.Log.Warn(" err at service UpdateTrip called", err)

		return trip, err
	}
	err = genRepository.Save(&Trip)
	if err != nil {
		resources.Log.Warn(" err at service UpdateTrip called", err)
		return trip, err
	}
	return Trip, err
}

func DeleteTrip(id interface{}) error {
	resources.Log.Info("service DeleteTrip called")

	err := genRepository.Delete(id, &model.Trip{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteTrip called", err)

		return err
	}
	return nil
}
