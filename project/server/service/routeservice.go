package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

func CreateRoute(Route model.Route) (string, error) {
	resources.Log.Info("service CreateRoute called")
	err := genRepository.Save(&Route)
	if err != nil {
		resources.Log.Warn(" err at service CreateRoute called", err)
		return "unable to save", err
	}
	return "route created", nil
}
func ReadRoute(RouteId int) (model.Route, error) {
	var Route model.Route
	resources.Log.Info("service ReadRoute called")
	err := genRepository.FindById(RouteId, &Route)
	if err != nil {
		resources.Log.Warn(" err at service ReadRoute called", err)
		return Route, err
	}
	return Route, err
}
func UpdateRoute(id int, Route model.Route) (model.Route, error) {
	var route model.Route
	err := genRepository.FindById(id, &route)
	if err != nil {
		resources.Log.Warn(" err at service UpdateRoute called", err)
		return route, err
	}
	err = genRepository.Save(&Route)
	if err != nil {
		resources.Log.Warn(" err at service UpdateRoute called", err)
		return route, err
	}
	return Route, err
}

func DeleteRoute(id interface{}) error {
	err := genRepository.Delete(id, &model.Route{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteRoute called", err)
		return err
	}
	return nil
}
