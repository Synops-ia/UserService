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
	ErrSessionNotFound    = errors.New("error session not found")
)

type SessionService interface {
	CreateSession(c context.Context, userToAuthenticate models.User) (string, error)
	DeleteSession(c context.Context, sessionId string) error
	LoginCheck(c context.Context, sessionId string) bool
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

func (s *SessionServiceImpl) CreateSession(c context.Context, userToAuthenticate models.User) (string, error) {
	userStored, err := s.userRepository.FindByEmail(c, userToAuthenticate.Email)
	if userStored == (models.User{}) {
		return "", ErrUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(userStored.Password), []byte(userToAuthenticate.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	session := models.Session{Id: uuid.New().String(), UserEmail: userStored.Email}
	sessionIdResult, err := s.sessionRepository.Save(c, session)
	if err != nil {
		return "", ErrCreatingSession
	}

	sessionId, ok := sessionIdResult.(string)
	if !ok {
		return "", ErrCreatingSession
	}

	return sessionId, nil
}

func (s *SessionServiceImpl) DeleteSession(c context.Context, sessionId string) error {
	deletedCount, err := s.sessionRepository.DeleteById(c, sessionId)
	if err != nil {
		return ErrDeletingSession
	}
	if deletedCount == 0 {
		return ErrSessionNotFound
	}

	return nil
}

func (s *SessionServiceImpl) LoginCheck(c context.Context, sessionId string) bool {
	session, err := s.sessionRepository.FindById(c, sessionId)
	if err != nil || session == (models.Session{}) {
		return false
	}
	return true
}
