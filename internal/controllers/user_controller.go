package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"

    "UserService/internal/services"
    "UserService/internal/models"
)

type UserController interface {
    CreateUser(c *gin.Context)
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
    c.JSON(http.StatusCreated, newUser)
}

func errorToCode(err error) int {
    switch err {
    case services.ErrUserAlreadyExists:
        return http.StatusConflict
    case services.ErrHashingPassword:
        return http.StatusInternalServerError
    default:
        return http.StatusInternalServerError
    }
}

