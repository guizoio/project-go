package repository

import (
	"awesomeProject1/src/domain/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u UserRepository) List() ([]entities.User, error) {
	var user []entities.User
	result := u.db.Find(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (u UserRepository) Create(data entities.User) error {
	if err := u.db.Create(&data); err != nil {
		return err.Error
	}
	return nil
}
