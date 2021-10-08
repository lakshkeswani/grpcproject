package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
	"grpcproject/project/server/repository"
	"grpcproject/project/server/service"
)

func (*Server) CreateDriver(ctx context.Context, req *projectpb.CreateDriverRequest) (*projectpb.CreateDriverResponse, error) {
	resources.Log.Info("service CreateDriver called")

	var Driver model.Driver
	Driver.Email = req.GetDriver().GetEmail()
	Driver.Password = req.GetDriver().GetPassword()

	Driver.Name = req.GetDriver().GetName()
	Driver.Password = req.GetDriver().GetPassword()
	Driver.ContactNo = req.GetDriver().GetContactNo()
	_, err := service.CreateDriver(Driver)

	if err != nil {
		return nil, err
	}

	return &projectpb.CreateDriverResponse{Result: "Driver created"}, nil

}

func (*Server) ReadDriver(ctx context.Context, req *projectpb.ReadDriverRequest) (*projectpb.ReadDriverResponse, error) {
	resources.Log.Info("service ReadDriver called")
	ok, err := tokenwork(ctx, int(req.GetDriverId()))
	if err != nil {
		return nil, err
	}
	if ok {

	}
	var Driver model.Driver
	Driver, err = service.ReadDriver(int(req.GetDriverId()))
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}

	return &projectpb.ReadDriverResponse{Driver: &projectpb.Driver{
		Id:        int64(Driver.ID),
		Email:     Driver.Email,
		Password:  Driver.Password,
		Name:      Driver.Name,
		ContactNo: Driver.ContactNo,
	}}, nil
}
func (*Server) UpdateDriver(ctx context.Context, req *projectpb.UpdateDriverRequest) (*projectpb.UpdateDriverResponse, error) {
	var Driver model.Driver

	resources.Log.Info("service UpdateDriver called")
	ok, err := tokenwork(ctx, int(req.GetDriver().GetId()))
	if err != nil {
		return nil, err
	}
	if ok {

	}

	Driver.ID = uint(req.GetDriver().GetId())
	Driver.Email = req.GetDriver().GetEmail()
	Driver.Password = req.GetDriver().GetPassword()
	Driver.Name = req.GetDriver().GetName()
	Driver.ContactNo = req.GetDriver().GetContactNo()
	Driver.Password = req.GetDriver().GetPassword()
	Driver, err = service.UpdateDriver(int(req.GetDriver().GetId()), Driver)
	if err != nil {
		return nil, err
	}
	return &projectpb.UpdateDriverResponse{Driver: &projectpb.Driver{
		Id:        req.GetDriver().GetId(),
		Email:     req.GetDriver().GetEmail(),
		Password:  req.GetDriver().GetPassword(),
		Name:      req.GetDriver().GetName(),
		ContactNo: req.GetDriver().GetContactNo(),
	}}, nil

}
func (*Server) DeleteDriver(ctx context.Context, req *projectpb.DeleteDriverRequest) (*projectpb.DeleteDriverResponse, error) {
	resources.Log.Info("service DeleteDriver called")
	ok, err := tokenwork(ctx, int(req.GetDriverId()))
	if err != nil {
		return nil, err
	}
	if ok {
	}
	err = service.DeleteDriver(req.GetDriverId())
	if err != nil {
		return nil, err
	}
	return &projectpb.DeleteDriverResponse{Result: "sucessfully deleted"}, nil

}

func (*Server) LoginDriver(ctx context.Context, req *projectpb.LoginRequest) (*projectpb.LoginResponse, error) {

	credentials := model.Credentials{
		Username: req.GetUserCredentials().GetUsername(),
		Password: req.GetUserCredentials().GetPassword(),
	}
	id, err := repository.GetDriver(credentials)
	if err != nil {
		return nil, err
	}
	token, err := GenrateToken(credentials, id)

	if err != nil {
		return nil, err
	}
	return &projectpb.LoginResponse{
		Jwt: token,
	}, nil
}
