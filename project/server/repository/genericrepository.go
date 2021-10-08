package repository

import (
	"errors"
	"gorm.io/gorm"
	"grpcproject/project/resources"
	"grpcproject/project/server/db"
	"grpcproject/project/server/model"
)

type GenericRepository struct {
}
type Genericfunction interface {
	Save(value interface{}) error
	FindById(value interface{}) error
	Delete(value interface{}) error
}

func (r GenericRepository) FindById(id interface{}, value interface{}) error {
	resources.Log.Info(" FindById Called  ")

	err := db.DB.First(value, id).Error
	if err != nil {
		resources.Log.Warn("error %v", err)

		errors.Is(err, gorm.ErrRecordNotFound)
		return err
	}
	db.DB.First(value, id)
	return nil
}

func (r GenericRepository) Delete(id interface{}, value model.IModel) error {
	resources.Log.Info(" Delete Called  ")

	err := db.DB.Delete(value, id).Error
	if err != nil {
		resources.Log.Warn("error %v", err)
		return err
	}
	db.DB.Delete(value, id)
	return nil
}

func (r GenericRepository) Save(value interface{}) error {
	resources.Log.Info(" Save Called  ")

	err := db.DB.Save(value).Error
	if err != nil {
		resources.Log.Warn("error %v", err)

		return err
	}
	db.DB.Save(value)
	return nil
}
