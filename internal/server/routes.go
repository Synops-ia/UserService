package server

import (
	"net/http"
	"github.com/gin-gonic/gin"

    "UserService/internal/repositories"
    "UserService/internal/services"
    "UserService/internal/controllers"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
    
    userRepository := repositories.NewUserRepositoryImpl()
    userService := services.NewUserServiceImpl(userRepository)
    userController := controllers.NewUserControllerImpl(userService)

    api := r.Group("api/v1")

    api.POST("/users", userController.CreateUser)

	return r
}

