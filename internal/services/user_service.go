package services

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"UserService/internal/models"
	"UserService/internal/repositories"
)

var (
	ErrUserAlreadyExists = errors.New("error user already exists")
	ErrHashingPassword   = errors.New("error hashing password")
	ErrSavingUser        = errors.New("error saving user")
)

type UserService interface {
	CreateUser(c context.Context, user models.User) error
}

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserServiceImpl(userRepository repositories.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (u *UserServiceImpl) CreateUser(c context.Context, user models.User) error {
	userStored, err := u.userRepository.FindByEmail(c, user.Email)
	if userStored != (models.User{}) {
		return ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return ErrHashingPassword
	}
	user.Password = string(hashedPassword)
	err = u.userRepository.Save(c, user)
	if err != nil {
		return ErrSavingUser
	}

	return nil
}
