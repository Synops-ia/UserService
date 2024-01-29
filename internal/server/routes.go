package server

import (
	"UserService/internal/controllers"
	"UserService/internal/repositories"
	"UserService/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	userRepository := repositories.NewUserRepositoryImpl(s.db)
	userService := services.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserControllerImpl(userService)

	sessionRepository := repositories.NewSessionRepositoryImpl(s.db)
	sessionService := services.NewSessionServiceImpl(sessionRepository, userRepository)
	sessionController := controllers.NewSessionControllerImpl(sessionService)

	api := r.Group("api/v1")

	api.POST("/users", userController.CreateUser)
	api.POST("/sessions", sessionController.CreateSession)
	api.DELETE("/sessions", sessionController.DeleteSession)
	api.GET("/me", sessionController.LoginCheck)

	return r
}
