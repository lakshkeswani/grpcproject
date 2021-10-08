package repository

import (
	"grpcproject/project/server/db"
	"grpcproject/project/server/model"
)

func GetDriver(credentials model.Credentials) (int, error) {
	var driver model.Driver
	err := db.DB.Where("email = ? AND password = ?", credentials.Username, credentials.Password).First(&driver).Error
	if err != nil {
		return 0, err
	}
	/*if driver.Email !=credentials.Username {
		return 0, status.Errorf(
			codes.PermissionDenied, fmt.Sprintf("worong Drivername or password"))
	}*/

	return int(driver.ID), nil
}
