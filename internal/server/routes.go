package server

import (
	"UserService/internal/controllers"
	"UserService/internal/repositories"
	"UserService/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5173/"},
		AllowMethods:     []string{"POST", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Set-Cookie"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
