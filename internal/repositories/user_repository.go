package repositories

import (
	"UserService/internal/models"
	"fmt"
)

type UserRepository interface {
    Save(user models.User) models.User
    FindByEmail(email string) (models.User, bool)
}

type UserRepositoryImpl struct {
    users map[string]models.User
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
    return &UserRepositoryImpl{
        users: make(map[string]models.User),
    }
}

func NewUserRepositoryImplWithUsers(users map[string]models.User) *UserRepositoryImpl {
    usersMapFromInput := make(map[string]models.User)
    for key, value := range users {
        usersMapFromInput[key] = value
    }

    return &UserRepositoryImpl{
        users: usersMapFromInput,
    }
}

func (u *UserRepositoryImpl) Save(user models.User) models.User {
    u.users[user.Email] = user
    return user
}

func (u *UserRepositoryImpl) FindByEmail(email string) (models.User, bool) {
    user, exists := u.users[email]
    fmt.Println("User: ", user)
    return user, exists
}

