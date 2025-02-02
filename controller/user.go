package controller

import (
	"akatech/dto"
	"akatech/helper"
	"akatech/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var payload dto.CreateUserRequest
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), nil))
		return
	}

	id, err := u.userService.NewUser(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ResponseAPI(false, http.StatusBadRequest, err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, helper.ResponseAPI(true, http.StatusOK, "success create new user", gin.H{
		"id":    id,
		"name":  payload.Name,
		"email": payload.Email,
	}))
}
