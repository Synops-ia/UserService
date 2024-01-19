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

func (u *UserRepositoryImpl) Save(user models.User) models.User {
    u.users[user.Email] = user
    return user
}

func (u *UserRepositoryImpl) FindByEmail(email string) (models.User, bool) {
    user, exists := u.users[email]
    fmt.Println("User: ", user)
    return user, exists
}

