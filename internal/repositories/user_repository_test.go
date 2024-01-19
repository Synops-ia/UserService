package repositories

import (
    "testing"
    "github.com/stretchr/testify/assert"
    
    "UserService/internal/models"
)

func TestUserRepository_FindByEmailAlreadyExists(t *testing.T) {
    expectedUser := models.User{Email: "email", Password: "password"}

    userRepo := NewUserRepositoryImpl();
    userRepo.users[expectedUser.Email] = expectedUser

    actualUser, exists := userRepo.FindByEmail(expectedUser.Email)

    assert.True(t, exists)
    assert.Equal(t, expectedUser, actualUser)
}

func TestUserRepository_FindByEmailInexistent(t *testing.T) {
    email := "email"

    userRepo := NewUserRepositoryImpl();

    actualUser, exists := userRepo.FindByEmail(email)

    assert.False(t, exists)
    assert.Equal(t, models.User{}, actualUser)
}

