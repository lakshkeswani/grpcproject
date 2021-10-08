package controller

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"grpcproject/project/projectpb"
	"grpcproject/project/resources"
	"grpcproject/project/server/db"
	"grpcproject/project/server/model"
	"grpcproject/project/server/repository"
	"grpcproject/project/server/service"
)

func (*Server) CreateUser(ctx context.Context, req *projectpb.CreateUserRequest) (*projectpb.CreateUserResponse, error) {

	var user model.User
	user.Email = req.GetUser().GetEmail()
	user.Password = req.GetUser().GetPassword()
	user.FirstName = req.GetUser().GetFirstname()
	user.LastName = req.GetUser().GetLastname()
	user.Password = req.GetUser().GetPassword()
	_, err := service.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return &projectpb.CreateUserResponse{Result: "user created"}, nil

}
func (*Server) ReadUser(ctx context.Context, req *projectpb.ReadUserRequest) (*projectpb.ReadUserResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	if ok {

	}

	var user model.User
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	db.DB.First(&user, req.GetUserId())

	return &projectpb.ReadUserResponse{User: &projectpb.User{
		Id:        int64(user.ID),
		Email:     user.Email,
		Password:  user.Password,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
	}}, nil
}
func (*Server) UpdateUser(ctx context.Context, req *projectpb.UpdateUserRequest) (*projectpb.UpdateUserResponse, error) {
	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, int(req.GetUser().GetId()))
	if err != nil {
		return nil, err
	}
	if ok {

	}

	var user model.User
	err = db.DB.First(&user, req.GetUser().GetId()).Error
	if err != nil {
		errors.Is(err, gorm.ErrRecordNotFound)
		return nil, err
	}
	db.DB.First(&user, req.GetUser().GetId())
	user.ID = uint(req.GetUser().GetId())
	user.Email = req.GetUser().GetEmail()
	user.Password = req.GetUser().GetPassword()
	user.FirstName = req.GetUser().GetFirstname()
	user.LastName = req.GetUser().GetLastname()
	user.Password = req.GetUser().GetPassword()
	err = db.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	db.DB.Save(&user)
	return &projectpb.UpdateUserResponse{User: &projectpb.User{
		Id:        req.GetUser().GetId(),
		Email:     req.GetUser().GetEmail(),
		Password:  req.GetUser().GetPassword(),
		Firstname: req.GetUser().GetFirstname(),
		Lastname:  req.GetUser().GetLastname(),
	}}, nil

}
func (*Server) DeleteUser(ctx context.Context, req *projectpb.DeleteUserRequest) (*projectpb.DeleteUserResponse, error) {

	resources.Log.Info("service CreateVehicle called")
	ok, err := tokenwork(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	if ok {

	}
	var user model.User
	db.DB.Delete(&user, req.GetUserId())
	return &projectpb.DeleteUserResponse{Result: "sucessfully deleted"}, nil

}

func (*Server) LoginUser(ctx context.Context, req *projectpb.LoginRequest) (*projectpb.LoginResponse, error) {

	credentials := model.Credentials{
		Username: req.GetUserCredentials().GetUsername(),
		Password: req.GetUserCredentials().GetPassword(),
	}
	id, err := repository.GetUser(credentials)
	if err != nil {
		return nil, err
	}
	token, err := GenrateToken(credentials, id)

	if err != nil {
		return nil, err
	}
	resources.Localsession.AddUser(token)
	return &projectpb.LoginResponse{
		Jwt: token,
	}, nil
}

func (*Server) Userlogout(ctx context.Context, in *projectpb.UserLogoutRequest) (*projectpb.UserLogoutResponse, error) {
	_, err := tokenwork(ctx, int(in.GetUserId()))
	if err != nil {

	}

	token, err := extactToken(ctx)
	if err != nil {
		return nil, err
	}
	ok := resources.Localsession.RemoveUser(token)
	if !ok {
		return &projectpb.UserLogoutResponse{Result: "logout failed user already loged out"}, nil

	}
	return &projectpb.UserLogoutResponse{Result: "logout sucessfully"}, nil
}
