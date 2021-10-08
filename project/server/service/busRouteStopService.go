package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

func CreateBusRouteStop(BusRouteStop model.BusRouteStop) (string, error) {

	resources.Log.Info("service CreateBusRouteStop called")
	err := genRepository.Save(&BusRouteStop)
	if err != nil {

		return "unable to save", err
	}
	return "busRouteStop created", nil
}
func ReadBusRouteStop(BusRouteStopId int) (model.BusRouteStop, error) {
	var BusRouteStop model.BusRouteStop
	resources.Log.Info("service ReadBusRouteStop called")
	err := genRepository.FindById(BusRouteStopId, &BusRouteStop)
	if err != nil {
		resources.Log.Warn(" err at service ReadBusRouteStop called", err)
		return BusRouteStop, err
	}
	return BusRouteStop, err

}
func UpdateBusRouteStop(id int, BusRouteStop model.BusRouteStop) (model.BusRouteStop, error) {
	var busRouteStop model.BusRouteStop
	err := genRepository.FindById(id, &busRouteStop)
	if err != nil {
		resources.Log.Warn(" err at service UpdateBusRouteStop called", err)
		return busRouteStop, err
	}
	err = genRepository.Save(&BusRouteStop)
	if err != nil {
		resources.Log.Warn(" err at service UpdateBusRouteStop called", err)
		return busRouteStop, err
	}
	return BusRouteStop, err
}

func DeleteBusRouteStop(id interface{}) error {
	err := genRepository.Delete(id, &model.BusRouteStop{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteBusRouteStop called", err)
		return err
	}
	return nil
}
