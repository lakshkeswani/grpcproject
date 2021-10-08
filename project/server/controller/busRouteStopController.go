package controller

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/db"
	"grpcproject/project/server/model"
	"grpcproject/project/server/service"
)

func (*Server) CreateBusRouteStop(ctx context.Context, req *projectpb.CreateBusRouteStopRequest) (*projectpb.CreateBusRouteStopResponse, error) {

	var BusRouteStop model.BusRouteStop
	BusRouteStop.ID = uint(req.GetBusRouteStop().GetBusRouteStopId())
	BusRouteStop.StationId = int(req.GetBusRouteStop().GetStationId())
	BusRouteStop.RouteID = int(req.GetBusRouteStop().GetRouteID())
	BusRouteStop.VehicleId = int(req.GetBusRouteStop().GetVehicleId())
	starttime, err := ptypes.Timestamp(req.GetBusRouteStop().GetStartTime())
	if err != nil {
		return nil, err
	}
	BusRouteStop.StartTime = starttime
	endtime, err := ptypes.Timestamp(req.GetBusRouteStop().GetEndTime())
	BusRouteStop.EndTime = endtime
	_, err = service.UpdateBusRouteStop(int(req.GetBusRouteStop().GetBusRouteStopId()), BusRouteStop)
	return &projectpb.CreateBusRouteStopResponse{Result: "BusRouteStop created"}, nil

}

func (*Server) ReadBusRouteStop(ctx context.Context, req *projectpb.ReadBusRouteStopRequest) (*projectpb.ReadBusRouteStopResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}

	if ok {
	}

	var BusRouteStop model.BusRouteStop
	BusRouteStop, err = service.ReadBusRouteStop(int(req.GetBusRouteStopId()))
	starttime, err := ptypes.TimestampProto(BusRouteStop.StartTime)
	if err != nil {
		return nil, err
	}
	endtime, err := ptypes.TimestampProto(BusRouteStop.EndTime)
	if err != nil {
		return nil, err
	}
	return &projectpb.ReadBusRouteStopResponse{BusRouteStop: &projectpb.BusRouteStop{
		BusRouteStopId: int64(BusRouteStop.ID),
		RouteID:        int64(BusRouteStop.RouteID),
		VehicleId:      int64(BusRouteStop.VehicleId),
		StationId:      int64(BusRouteStop.StationId),
		StartTime:      starttime,
		EndTime:        endtime,
	}}, nil
}
func (*Server) UpdateBusRouteStop(ctx context.Context, req *projectpb.UpdateBusRouteStopRequest) (*projectpb.UpdateBusRouteStopResponse, error) {
	var BusRouteStop model.BusRouteStop

	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}

	db.DB.First(&BusRouteStop, req.GetBusRouteStop().GetBusRouteStopId())
	BusRouteStop.ID = uint(req.GetBusRouteStop().GetBusRouteStopId())
	BusRouteStop.StationId = int(req.GetBusRouteStop().GetStationId())
	BusRouteStop.RouteID = int(req.GetBusRouteStop().GetRouteID())
	BusRouteStop.VehicleId = int(req.GetBusRouteStop().GetVehicleId())
	starttime, err := ptypes.Timestamp(req.GetBusRouteStop().GetStartTime())
	BusRouteStop.StartTime = starttime
	endtime, err := ptypes.Timestamp(req.GetBusRouteStop().GetEndTime())
	BusRouteStop.EndTime = endtime
	service.UpdateBusRouteStop(int(req.GetBusRouteStop().GetBusRouteStopId()), BusRouteStop)
	if err != nil {
		return nil, err
	}
	starttime1, err := ptypes.TimestampProto(BusRouteStop.StartTime)
	endtime1, err := ptypes.TimestampProto(BusRouteStop.EndTime)

	return &projectpb.UpdateBusRouteStopResponse{BusRouteStop: &projectpb.BusRouteStop{
		BusRouteStopId: int64(BusRouteStop.ID),
		RouteID:        int64(BusRouteStop.RouteID),
		VehicleId:      int64(BusRouteStop.VehicleId),
		StationId:      int64(BusRouteStop.StationId),
		StartTime:      starttime1,
		EndTime:        endtime1,
	}}, nil

}
func (*Server) DeleteBusRouteStop(ctx context.Context, req *projectpb.DeleteBusRouteStopRequest) (*projectpb.DeleteBusRouteStopResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, 0)
	if err != nil {
		return nil, err
	}
	if ok {
	}
	err = service.DeleteBusRouteStop(req.GetBusRouteStopId())
	if err != nil {
		return nil, err
	}
	return &projectpb.DeleteBusRouteStopResponse{Result: "sucessfully deleted"}, nil
}
