package server

import (
    "os"
    "net/http"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"

    "UserService/internal/repositories"
    "UserService/internal/services"
    "UserService/internal/controllers"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

    store := cookie.NewStore([]byte(os.Getenv("SESSION_KEY")))
    store.Options(sessions.Options{
        MaxAge: 60 * 15, // 15 minutes 
    })
    r.Use(sessions.Sessions("session", store))
    
    userRepository := repositories.NewUserRepositoryImpl()
    userService := services.NewUserServiceImpl(userRepository)
    userController := controllers.NewUserControllerImpl(userService)

    api := r.Group("api/v1")

    api.POST("/users", userController.CreateUser)
    api.POST("/sessions", userController.CreateSession)

	return r
}

