package controllers

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"

	"UserService/internal/models"
	"UserService/internal/services"
)

type UserController interface {
	CreateUser(c *gin.Context)
	CreateSession(c *gin.Context)
	DeleteSession(c *gin.Context)
}

type UserControllerImpl struct {
	userService services.UserService
}

func NewUserControllerImpl(userService services.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		userService: userService,
	}
}

func (u *UserControllerImpl) CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := u.userService.CreateUser(c, newUser)
	if err != nil {
		c.JSON(errorToCode(err), err.Error())
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, newUser)
}

func (u *UserControllerImpl) CreateSession(c *gin.Context) {
	var userToAuthenticate models.User

	if err := c.ShouldBindJSON(&userToAuthenticate); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	session := sessions.Default(c)

	err := u.userService.CreateSession(c, session, userToAuthenticate)
	if err != nil {
		c.JSON(errorToCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, session.Get("email"))
}

func (u *UserControllerImpl) DeleteSession(c *gin.Context) {
	println("cookies: " + c.GetHeader("Cookie"))
	session := sessions.Default(c)
	err := u.userService.DeleteSession(session)
	if err != nil {
		c.JSON(errorToCode(err), err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func errorToCode(err error) int {
	switch {
	case errors.Is(err, services.ErrUserNotFound):
		return http.StatusNotFound
	case errors.Is(err, services.ErrUserAlreadyExists):
		return http.StatusConflict
	case errors.Is(err, services.ErrHashingPassword):
		return http.StatusInternalServerError
	case errors.Is(err, services.ErrInvalidCredentials):
		return http.StatusUnauthorized
	case errors.Is(err, services.ErrCreatingSession):
		return http.StatusInternalServerError
	case errors.Is(err, services.ErrDeletingSession):
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
