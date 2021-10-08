package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
	"grpcproject/project/server/service"
)

func (*Server) CreateRoute(ctx context.Context, req *projectpb.CreateRouteRequest) (*projectpb.CreateRouteResponse, error) {

	var Route model.Route
	Route.RouteName = req.GetRoute().GetRouteName()
	_, err := service.CreateRoute(Route)
	if err != nil {
		return nil, err
	}
	return &projectpb.CreateRouteResponse{Result: "Route created"}, nil

}

func (*Server) ReadRoute(ctx context.Context, req *projectpb.ReadRouteRequest) (*projectpb.ReadRouteResponse, error) {

	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}
	var Route model.Route
	Route, err = service.ReadRoute(int(req.GetRouteId()))
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	return &projectpb.ReadRouteResponse{Route: &projectpb.Route{
		RouteId:   int64(Route.ID),
		RouteName: Route.RouteName,
	}}, nil
}
func (*Server) UpdateRoute(ctx context.Context, req *projectpb.UpdateRouteRequest) (*projectpb.UpdateRouteResponse, error) {
	var Route model.Route

	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}

	Route.ID = uint(req.GetRoute().GetRouteId())
	Route.RouteName = req.GetRoute().GetRouteName()
	_, err = service.UpdateRoute(int(req.GetRoute().GetRouteId()), Route)
	if err != nil {
		return nil, err
	}
	return &projectpb.UpdateRouteResponse{Route: &projectpb.Route{
		RouteId:   req.GetRoute().GetRouteId(),
		RouteName: req.GetRoute().GetRouteName(),
	}}, nil

}
func (*Server) DeleteRoute(ctx context.Context, req *projectpb.DeleteRouteRequest) (*projectpb.DeleteRouteResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}
	err = service.DeleteRoute(req.GetRouteId())
	if err != nil {
		return nil, err
	}
	return &projectpb.DeleteRouteResponse{Result: "successfully deleted"}, nil

}
