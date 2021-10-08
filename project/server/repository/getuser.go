package repository

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/db"
	"grpcproject/project/server/model"
)

func GetUser(credentials model.Credentials) (int, error) {
	var user model.User
	err := db.DB.Where("email= ? AND password= ?", credentials.Username, credentials.Password).First(&user).Error
	if err != nil {
		resources.Log.Warn("user can't log in  %v", err)
	}
	return int(user.ID), nil
}
