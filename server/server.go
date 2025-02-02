package server

import (
	"akatech/config"
	"akatech/controller"
	"akatech/model"
	"akatech/repository"
	"akatech/service"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func App(ctx context.Context) *gin.Engine {
	//Initial Config
	configBase := config.NewConfig()

	//Migrate
	err := configBase.DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err.Error())
	}

	//Repository
	userRepository := repository.NewUserRepository(configBase.DB)

	//Service
	userService := service.NewUserService(userRepository)

	//Controller
	userController := controller.NewUserController(userService)

	app := gin.Default()

	app.Use(gin.Recovery())
	app.Use(gin.Logger())

	// cors	config
	cfg := cors.DefaultConfig()
	cfg.AllowOrigins = []string{"https://www.carikerjaai.com", "http://localhost:5173", "https://admin.carikerjaai.com"}
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	cfg.AllowHeaders = []string{"*"}

	app.Use(cors.New(cfg))

	//Webhook

	apiRoutesV1 := app.Group("/api/v1")

	// User
	apiRoutesV1.POST("/user", userController.CreateUser)
	return app
}
