package controller

import (
	"context"
	"errors"
	"github.com/golang/protobuf/ptypes"
	"gorm.io/gorm"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
	"grpcproject/project/server/service"
)

func (*Server) CreateTrip(ctx context.Context, req *projectpb.CreateTripRequest) (*projectpb.CreateTripResponse, error) {

	var Trip model.Trip
	Trip.ID = uint(req.GetTrip().GetTripId())
	Trip.BusRouteStopId = int(req.GetTrip().GetBusRouteStopId())
	scheduledate, err := ptypes.Timestamp(req.GetTrip().GetScheduleDate())
	if err != nil {
	}
	Trip.ScheduleDate = scheduledate
	Trip.Fare = int(req.GetTrip().GetFare())
	Trip.DriverId = int(req.GetTrip().GetDriverId())
	_, err = service.CreateTrip(Trip)
	if err != nil {
		return nil, err
	}
	return &projectpb.CreateTripResponse{Result: "Trip created"}, nil

}

func (*Server) ReadTrip(ctx context.Context, req *projectpb.ReadTripRequest) (*projectpb.ReadTripResponse, error) {
	resources.Log.Info(" Delete Called  ")

	//ok,err:=tokenwork(ctx, 0)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if ok {
	//}

	var Trip model.Trip
	Trip, err := service.ReadTrip(int(req.GetTripId()))
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	trip, err := tripToTrippb(Trip)
	return &projectpb.ReadTripResponse{Trip: trip}, nil
}
func tripToTrippb(Trip model.Trip) (*projectpb.Trip, error) {
	ScheduleDate, err := ptypes.TimestampProto(Trip.ScheduleDate)
	if err != nil {
		return nil, err
	}
	return &projectpb.Trip{
		TripId:       int64(Trip.ID),
		ScheduleDate: ScheduleDate,
		DriverId:     int64(Trip.DriverId),
		Fare:         int64(Trip.Fare),
	}, nil
}
func (*Server) UpdateTrip(ctx context.Context, req *projectpb.UpdateTripRequest) (*projectpb.UpdateTripResponse, error) {
	var Trip model.Trip

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
	Trip.ID = uint(req.GetTrip().GetTripId())
	Trip.BusRouteStopId = int(req.GetTrip().GetBusRouteStopId())
	scheduledate, err := ptypes.Timestamp(req.GetTrip().GetScheduleDate())
	if err != nil {
	}
	Trip.ScheduleDate = scheduledate
	Trip.Fare = int(req.GetTrip().GetFare())
	_, err = service.UpdateTrip(int(req.GetTrip().GetTripId()), Trip)
	if err != nil {
		return nil, err
	}
	trip, err := tripToTrippb(Trip)
	return &projectpb.UpdateTripResponse{Trip: trip}, nil

}
func (*Server) DeleteTrip(ctx context.Context, req *projectpb.DeleteTripRequest) (*projectpb.DeleteTripResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}
	err = service.DeleteTrip(req.GetTripId())

	if err != nil {
		return nil, err
	}
	return &projectpb.DeleteTripResponse{Result: "sucessfully deleted"}, nil
}
