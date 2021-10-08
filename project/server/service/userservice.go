package service

import (
	"grpcproject/project/resources"
	"grpcproject/project/server/model"
)

func CreateUser(User model.User) (string, error) {
	resources.Log.Info("service CreateUser called")

	resources.Log.Info("service CreateVehicle called")
	err := genRepository.Save(&User)
	if err != nil {
		resources.Log.Warn(" err at service CreateUser called", err)

		return "unable to save", err
	}
	return "user created", nil
}
func ReadUser(UserId int) (model.User, error) {
	resources.Log.Info("service ReadUser called")

	var User model.User
	resources.Log.Info("service CreateVehicle called")
	err := genRepository.FindById(UserId, &User)
	if err != nil {
		resources.Log.Warn(" err at service ReadUser called", err)
		return User, err
	}
	resources.Log.Info("returing response to controller  service ReadUser called")

	return User, err

}
func UpdateUser(id int, User model.User) (model.User, error) {
	resources.Log.Info("service UpdateUser called")

	var user model.User
	err := genRepository.FindById(id, &user)
	if err != nil {
		resources.Log.Warn(" err at service UpdateUser called", err)

		return user, err
	}
	err = genRepository.Save(&User)
	if err != nil {
		resources.Log.Warn(" err at service UpdateUser called", err)
		return user, err
	}
	resources.Log.Info("returing response to controller  service UpdateUser called")

	return User, err
}

func DeleteUser(id interface{}) error {
	resources.Log.Info("service DeleteUser called")
	err := genRepository.Delete(id, &model.User{})
	if err != nil {
		resources.Log.Warn(" err at service DeleteUser called", err)
		return err
	}
	return nil
}
