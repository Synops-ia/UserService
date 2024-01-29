package server

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"os"
	"time"

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
		MaxAge:   60 * 15, // 15 minutes
		HttpOnly: false,
		Secure:   false,
		Domain:   "localhost",
		Path:     "/",
	})
	r.Use(sessions.Sessions("session", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5173/"},
		AllowMethods:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Set-Cookie"},
		ExposeHeaders:    []string{"Content-Type", "Set-Cookie"},
		MaxAge:           12 * time.Hour,
	}))

	userRepository := repositories.NewUserRepositoryImpl(s.db)
	userService := services.NewUserServiceImpl(userRepository)
	userController := controllers.NewUserControllerImpl(userService)

	api := r.Group("api/v1")

	api.POST("/users", userController.CreateUser)
	api.POST("/sessions", userController.CreateSession)
	api.DELETE("/sessions", userController.DeleteSession)

	return r
}
