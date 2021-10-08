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

func (*Server) CreateStation(ctx context.Context, req *projectpb.CreateStationRequest) (*projectpb.CreateStationResponse, error) {
	var Station model.Station
	Station.StationName = req.GetStation().GetName()
	_, err := service.CreateStation(Station)
	if err != nil {
		return nil, err
	}

	return &projectpb.CreateStationResponse{Result: "Station created"}, nil

}

func (*Server) ReadStation(ctx context.Context, req *projectpb.ReadStationRequest) (*projectpb.ReadStationResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}

	if ok {
	}

	var Station model.Station
	Station, err = service.ReadStation(int(req.GetStationId()))
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}

	return &projectpb.ReadStationResponse{Station: &projectpb.Station{
		ID:   int64(Station.ID),
		Name: Station.StationName,
	}}, nil
}
func (*Server) UpdateStation(ctx context.Context, req *projectpb.UpdateStationRequest) (*projectpb.UpdateStationResponse, error) {
	var Station model.Station

	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}

	if err != nil {
		return nil, err
	}
	Station.ID = uint(req.GetStation().GetID())
	Station.StationName = req.GetStation().GetName()
	_, err = service.UpdateStation(int(req.GetStation().GetID()), Station)
	if err != nil {
		return nil, err
	}
	return &projectpb.UpdateStationResponse{Station: &projectpb.Station{
		ID:   req.GetStation().GetID(),
		Name: req.GetStation().GetName(),
	}}, nil

}
func (*Server) DeleteStation(ctx context.Context, req *projectpb.DeleteStationRequest) (*projectpb.DeleteStationResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}
	err = service.DeleteStation(req.GetStationId())
	if err != nil {
		return nil, err
	}
	return &projectpb.DeleteStationResponse{Result: "successfully deleted"}, nil

}
