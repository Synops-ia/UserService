package server

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"UserService/internal/controllers"
	"UserService/internal/repositories"
	"UserService/internal/services"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
	store.Options(sessions.Options{
		MaxAge: 60 * 15, // 15 minutes
	})
	r.Use(sessions.Sessions("session", store))

	userRepository := repositories.NewUserRepositoryImpl(s.db)
	userService := services.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserControllerImpl(userService)

	api := r.Group("api/v1")

	api.POST("/users", userController.CreateUser)
	api.POST("/sessions", userController.CreateSession)
	api.DELETE("/sessions", userController.DeleteSession)

	return r
}
