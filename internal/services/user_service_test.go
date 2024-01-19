package services

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "UserService/internal/models"
    "UserService/internal/repositories"
)

func TestUserService_CreateUserAlreadyExists(t *testing.T) {
    expectedUser := models.User{Email: "email", Password: "password"}
    expectedUsers := map[string]models.User{expectedUser.Email: expectedUser}

    userRepo := repositories.NewUserRepositoryImplWithUsers(expectedUsers)
    userService := NewUserServiceImpl(userRepo)
    actualUser, err := userService.CreateUser(expectedUser)

    assert.Equal(t, ErrUserAlreadyExists, err)
    assert.Equal(t, models.User{}, actualUser)
}

func TestUserService_CreateUser(t *testing.T) {
    expectedUser := models.User{Email: "email", Password: "password"}

    userRepo := repositories.NewUserRepositoryImpl()
    userService := NewUserServiceImpl(userRepo)
    actualUser, err := userService.CreateUser(expectedUser)

    assert.Nil(t, err)
    assert.Equal(t, expectedUser.Email, actualUser.Email)
    assert.NotEqual(t, "", actualUser.Password)
}

