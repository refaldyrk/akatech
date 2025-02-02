package service

import (
	"akatech/dto"
	"akatech/helper"
	"akatech/model"
	"akatech/repository"
	"errors"
	"github.com/rs/xid"
	"time"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (u *UserService) NewUser(req dto.CreateUserRequest) (uint, error) {
	_, err := u.userRepository.Find("email", req.Email)
	if err == nil {
		return 0, errors.New("email has registered")
	}

	hashedpPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return 0, err
	}
	newUser := model.User{
		UserID:    xid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  hashedpPassword,
		CreatedAt: time.Now().UTC().Unix(),
	}
	id, err := u.userRepository.Create(newUser)
	if err != nil {
		return 0, err
	}

	return id, nil
}
