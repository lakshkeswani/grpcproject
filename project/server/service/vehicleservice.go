package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
	"grpcproject/project/server/repository"
)

var genRepository repository.GenericRepository

func CreateVehicle(Vehicle model.Vehicle) (string, error) {
	resources.Log.Info("service CreateVehicle called")

	resources.Log.Info("service CreateVehicle called")
	err := genRepository.Save(&Vehicle)
	if err != nil {
		resources.Log.Warn(" err at service ReadVehicle called", err)

		return "unable to save", err
	}
	return "vehicle created", nil
}
func ReadVehicle(VehicleId int) (model.Vehicle, error) {
	resources.Log.Info("service CreateVehicle called")
	var Vehicle model.Vehicle
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.FindById(VehicleId, &Vehicle)
	if err != nil {
		resources.Log.Warn(" err at service ReadVehicle called", err)

		return Vehicle, err
	}
	return Vehicle, err

}
func UpdateVehicle(id int, Vehicle model.Vehicle) (model.Vehicle, error) {
	resources.Log.Info("service UpdateVehicle called")
	var vehicle model.Vehicle
	err := genRepository.FindById(id, &vehicle)
	if err != nil {
		resources.Log.Warn(" err at service CreateTrip called", err)
		return vehicle, err
	}
	err = genRepository.Save(&Vehicle)
	if err != nil {
		resources.Log.Warn(" err at service CreateTrip called", err)
		return vehicle, err
	}
	return Vehicle, err
}

func DeleteVehicle(id interface{}) error {
	resources.Log.Info("service DeleteVehicle called")

	err := genRepository.Delete(id, &model.Vehicle{})
	if err != nil {
		resources.Log.Fatal("err at service DeleteVehicle ", err)
		return err
	}
	return nil
}
