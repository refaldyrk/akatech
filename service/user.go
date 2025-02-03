package service

import (
	"akatech/dto"
	"akatech/helper"
	"akatech/model"
	"akatech/repository"
	"context"
	"errors"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/xid"
	"time"
)

type UserService struct {
	userRepository *repository.UserRepository
	queue          amqp091.Queue
	channel        *amqp091.Channel
}

func NewUserService(userRepository *repository.UserRepository, queue amqp091.Queue, channel *amqp091.Channel) *UserService {
	return &UserService{userRepository, queue, channel}
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

	//Send Message
	err = u.channel.PublishWithContext(context.Background(), "", u.queue.Name, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(newUser.UserID),
	})

	if err != nil {
		return 0, nil
	}

	return id, nil
}
