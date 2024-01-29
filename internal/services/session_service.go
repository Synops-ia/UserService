package services

import (
	"UserService/internal/models"
	"UserService/internal/repositories"
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound       = errors.New("error user not found")
	ErrInvalidCredentials = errors.New("error invalid credentials")
	ErrCreatingSession    = errors.New("error creating session")
	ErrDeletingSession    = errors.New("error deleting session")
)

type SessionService interface {
	CreateSession(c context.Context, userToAuthenticate models.User) (interface{}, error)
	DeleteSession(c context.Context, sessionId string) error
}

type SessionServiceImpl struct {
	sessionRepository repositories.SessionRepository
	userRepository    repositories.UserRepository
}

func NewSessionServiceImpl(sessionRepository repositories.SessionRepository, userRepository repositories.UserRepository) SessionService {
	return &SessionServiceImpl{
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}

func (s *SessionServiceImpl) CreateSession(c context.Context, userToAuthenticate models.User) (interface{}, error) {
	userStored, err := s.userRepository.FindByEmail(c, userToAuthenticate.Email)
	if userStored == (models.User{}) {
		return nil, ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(userStored.Password), []byte(userToAuthenticate.Password))
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	session := models.Session{Id: uuid.New().String(), UserEmail: userStored.Email}
	sessionId, err := s.sessionRepository.Save(c, session)
	if err != nil {
		return nil, ErrCreatingSession
	}

	return sessionId, nil
}

func (s *SessionServiceImpl) DeleteSession(c context.Context, sessionId string) error {
	err := s.sessionRepository.DeleteById(c, sessionId)
	if err != nil {
		return ErrDeletingSession
	}

	return nil
}
