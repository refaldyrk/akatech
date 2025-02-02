package repository

import (
	"akatech/model"
	"fmt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepository {
	return &UserRepository{DB}
}

func (u UserRepository) Create(user model.User) (uint, error) {
	result := u.DB.Create(&user)
	return user.ID, result.Error
}

func (u UserRepository) Find(field, value string) (model.User, error) {
	var user model.User
	err := u.DB.Where(fmt.Sprintf("%s = ?", field), value).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
