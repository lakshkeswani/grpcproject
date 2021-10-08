package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

func CreateDriver(Driver model.Driver) (string, error) {

	resources.Log.Info("service CreateVehicle called")
	err := genRepository.Save(&Driver)
	if err != nil {

		return "unable to save", err
	}
	return "Driver created", nil
}
func ReadDriver(DriverId int) (model.Driver, error) {
	var Driver model.Driver
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.FindById(DriverId, &Driver)
	if err != nil {
		resources.Log.Warn(" err at service ReadDriver called", err)
		return Driver, err
	}
	return Driver, err

}
func UpdateDriver(id int, Driver model.Driver) (model.Driver, error) {
	var driver model.Driver
	err := genRepository.FindById(id, &driver)
	if err != nil {
		resources.Log.Warn(" err at service UpdateDriver called", err)
		return driver, err
	}
	err = genRepository.Save(&Driver)
	if err != nil {
		resources.Log.Warn(" err at service UpdateDriver called", err)
		return driver, err
	}
	return Driver, err
}
func DeleteDriver(id interface{}) error {
	err := genRepository.Delete(id, &model.Driver{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteDriver called", err)
		return err
	}
	return nil

}
