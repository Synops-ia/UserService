package server

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"os"
	"time"

	"UserService/internal/controllers"
	"UserService/internal/repositories"
	"UserService/internal/services"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))

	store.Options(sessions.Options{
		MaxAge:   60 * 15, // 15 minutes
		HttpOnly: false,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		Domain:   "localhost",
		Path:     "/",
	})
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "POST", "GET", "DELETE", "PATCH"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Type", "Set-Cookie"},
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
