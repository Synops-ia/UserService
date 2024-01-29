package services

import (
	"context"
	"errors"

	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"

	"UserService/internal/models"
	"UserService/internal/repositories"
)

var (
	ErrUserNotFound       = errors.New("error user not found")
	ErrUserAlreadyExists  = errors.New("error user already exists")
	ErrHashingPassword    = errors.New("error hashing password")
	ErrSavingUser         = errors.New("error saving user")
	ErrInvalidCredentials = errors.New("error invalid credentials")
	ErrCreatingSession    = errors.New("error creating session")
	ErrDeletingSession    = errors.New("error deleting session")
)

type UserService interface {
	CreateUser(c context.Context, user models.User) error
	CreateSession(c context.Context, session sessions.Session, userToAuthenticate models.User) error
	DeleteSession(session sessions.Session) error
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

func (u *UserServiceImpl) CreateSession(c context.Context, session sessions.Session, userToAuthenticate models.User) error {
	userStored, err := u.userRepository.FindByEmail(c, userToAuthenticate.Email)
	if userStored == (models.User{}) {
		return ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(userStored.Password), []byte(userToAuthenticate.Password))
	if err != nil {
		return ErrInvalidCredentials
	}

	session.Set("email", userToAuthenticate.Email)
	session.Options(sessions.Options{
		HttpOnly: false,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
	})
	err = session.Save()
	if err != nil {
		return ErrCreatingSession
	}

	return nil
}

func (u *UserServiceImpl) DeleteSession(session sessions.Session) error {
	session.Set("email", "")
	session.Clear()
	session.Options(sessions.Options{Path: "/", MaxAge: -1})
	err := session.Save()
	if err != nil {
		return ErrDeletingSession
	}

	return nil
}
