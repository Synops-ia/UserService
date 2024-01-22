package services

import (
    "errors"

    "golang.org/x/crypto/bcrypt"
    "github.com/gin-contrib/sessions"

    "UserService/internal/models"
    "UserService/internal/repositories"
)

var (
    ErrUserNotFound = errors.New("error user not found")
    ErrUserAlreadyExists = errors.New("error user already exists")
    ErrHashingPassword = errors.New("error hashing password")
    ErrInvalidCredentials = errors.New("error invalid credentials")
    ErrCreatingSession = errors.New("error creating session")
)

type UserService interface {
    CreateUser(user models.User) (models.User, error)
    CreateSession(session sessions.Session, userToAuthenticate models.User) error
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

func (u *UserServiceImpl) CreateSession(session sessions.Session, userToAuthenticate models.User) error {
    userStored, exists := u.userRepository.FindByEmail(userToAuthenticate.Email)
    if !exists {
        return ErrUserNotFound
    }

    err := bcrypt.CompareHashAndPassword([]byte(userStored.Password), []byte(userToAuthenticate.Password))
    if err != nil {
        return ErrInvalidCredentials
    }

    session.Set("email", userToAuthenticate.Email)
    err = session.Save()
    if err != nil {
        return ErrCreatingSession
    }

    return nil 
}

