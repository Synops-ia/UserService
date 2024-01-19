package repositories

import (
    "testing"
    "github.com/stretchr/testify/assert"
    
    "UserService/internal/models"
)

func TestUserRepository_FindByEmailAlreadyExists(t *testing.T) {
    email := "email"
    password := "password"

    userRepo := NewUserRepositoryImpl();
    userRepo.users["email"] = models.User{Email: email, Password: password}

    user, exists := userRepo.FindByEmail(email)

    assert.True(t, exists)
    assert.NotNil(t, user)
}

