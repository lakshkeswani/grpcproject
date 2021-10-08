package controller

import (
	"context"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/db"
	"grpcproject/project/server/model"
	"grpcproject/project/server/service"
)

func (*Server) CreateVehicle(ctx context.Context, req *projectpb.CreateVehicleRequest) (*projectpb.CreateVehicleResponse, error) {
	resources.Log.Info(" CreateVehicle Called  ")

	var Vehicle model.Vehicle
	Vehicle.Capacity = req.GetVehicle().GetCapacity()
	Vehicle.VehicleType = req.GetVehicle().GetVehicleType()
	Vehicle.RegistrationNo = req.GetVehicle().GetRegistrationNo()
	result, err := service.CreateVehicle(Vehicle)
	if err != nil {
		resources.Log.Warn("CreateVehicle %v", err)

		return nil, err
	}
	return &projectpb.CreateVehicleResponse{Result: result}, err

}

func (*Server) ReadVehicle(ctx context.Context, req *projectpb.ReadVehicleRequest) (*projectpb.ReadVehicleResponse, error) {
	resources.Log.Info(" ReadVehicle Called  ")

	//ok,err:=tokenwork(ctx, 0)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if ok {
	//}

	var Vehicle model.Vehicle
	Vehicle, err := service.ReadVehicle(int(req.GetVehicleId()))
	if err != nil {
		resources.Log.Warn("ReadVehicle %v", err)

		return nil, err
	}
	return &projectpb.ReadVehicleResponse{Vehicle: &projectpb.Vehicle{
		Id:             int64(Vehicle.ID),
		VehicleType:    Vehicle.VehicleType,
		Capacity:       Vehicle.Capacity,
		RegistrationNo: Vehicle.RegistrationNo,
	}}, nil
}
func (*Server) UpdateVehicle(ctx context.Context, req *projectpb.UpdateVehicleRequest) (*projectpb.UpdateVehicleResponse, error) {
	var Vehicle model.Vehicle
	resources.Log.Info(" CreateVehicle Called  ")

	//ok,err:=tokenwork(ctx, 0)
	//if err != nil {
	//	return nil, err
	//}
	//if ok {
	//}
	Vehicle.ID = uint(req.GetVehicle().GetId())
	Vehicle.Capacity = req.GetVehicle().GetCapacity()
	Vehicle.VehicleType = req.GetVehicle().GetVehicleType()
	Vehicle.RegistrationNo = req.GetVehicle().GetRegistrationNo()
	_, err := service.UpdateVehicle(int(req.GetVehicle().GetId()), Vehicle)
	if err != nil {
		resources.Log.Warn("UpdateVehicle error %v", err)

		return nil, err
	}

	return &projectpb.UpdateVehicleResponse{Vehicle: &projectpb.Vehicle{
		Id:             req.GetVehicle().GetId(),
		VehicleType:    req.GetVehicle().GetVehicleType(),
		RegistrationNo: req.GetVehicle().GetRegistrationNo(),
		Capacity:       req.GetVehicle().GetCapacity(),
	}}, nil

}
func (*Server) DeleteVehicle(ctx context.Context, req *projectpb.DeleteVehicleRequest) (*projectpb.DeleteVehicleResponse, error) {
	var Vehicle model.Vehicle
	resources.Log.Info(" DeleteVehicle Called  ")

	_, err := tokenwork(ctx, 0)
	if err != nil {
		resources.Log.Warn("DeleteVehicle error %v", err)

		return nil, err
	}

	err = service.DeleteVehicle(req.GetVehicleId())

	if err != nil {
		resources.Log.Warn("error %v", err)

		return nil, err
	}
	db.DB.Delete(&Vehicle, req.GetVehicleId())
	return &projectpb.DeleteVehicleResponse{Result: "sucessfully deleted"}, nil

}
