package services

import (
    "errors"
    bcrypt "golang.org/x/crypto/bcrypt"

    "UserService/internal/models"
    "UserService/internal/repositories"
)

var (
    ErrUserAlreadyExists = errors.New("error user already exists")
    ErrHashingPassword = errors.New("error hashing password")
)

type UserService interface {
    CreateUser(user models.User) (models.User, error)
}

type UserServiceImpl struct {
    userRepository repositories.UserRepository 
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{
        userRepository: userRepository,
    }
}

func (u *UserServiceImpl) CreateUser(user models.User) (models.User, error) {
    if _, exists := u.userRepository.FindByEmail(user.Email); exists {
        return models.User{}, ErrUserAlreadyExists
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, ErrHashingPassword
	}
	user.Password = string(hashedPassword)

    return u.userRepository.Save(user), nil
}

