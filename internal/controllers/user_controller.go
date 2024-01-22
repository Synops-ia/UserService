package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"

    "UserService/internal/services"
    "UserService/internal/models"
)

type UserController interface {
    CreateUser(c *gin.Context)
    CreateSession(c *gin.Context)
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

    newUser, err := u.userService.CreateUser(newUser)
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
    err := u.userService.CreateSession(session, userToAuthenticate)
    if err != nil {
        c.JSON(errorToCode(err), err.Error())
        return
    }

    c.JSON(http.StatusOK, session.Get("email"))
}

func errorToCode(err error) int {
    switch err {
    case services.ErrUserNotFound:
        return http.StatusNotFound
    case services.ErrUserAlreadyExists:
        return http.StatusConflict
    case services.ErrHashingPassword:
        return http.StatusInternalServerError
    case services.ErrInvalidCredentials:
        return http.StatusUnauthorized
    case services.ErrCreatingSession:
        return http.StatusInternalServerError
    default:
        return http.StatusInternalServerError
    }
}

