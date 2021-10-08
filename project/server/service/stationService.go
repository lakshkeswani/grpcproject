package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

func CreateStation(Station model.Station) (string, error) {
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.Save(&Station)
	if err != nil {
		resources.Log.Warn(" err at service CreateStation called", err)
		return "unable to save", err
	}
	return "Station created", nil
}
func ReadStation(StationId int) (model.Station, error) {
	resources.Log.Info("service ReadStation called")

	var Station model.Station
	err := genRepository.FindById(StationId, &Station)
	if err != nil {
		resources.Log.Warn(" err at service ReadStation called", err)
		return Station, err
	}
	return Station, err

}
func UpdateStation(id int, Station model.Station) (model.Station, error) {
	resources.Log.Info("service UpdateStation called")
	var station model.Station
	err := genRepository.FindById(id, &station)
	if err != nil {
		resources.Log.Warn(" err at service UpdateStation called", err)
		return station, err
	}
	err = genRepository.Save(&Station)
	if err != nil {
		resources.Log.Warn(" err at service UpdateStation called", err)
		return station, err
	}
	return Station, err
}
func DeleteStation(id interface{}) error {
	resources.Log.Info("service DeleteStation called")

	err := genRepository.Delete(id, &model.Station{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteStation called", err)
		return err
	}
	return nil

}
