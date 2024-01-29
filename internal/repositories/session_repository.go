package repositories

import (
	"UserService/internal/database"
	"UserService/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type SessionRepository interface {
	Save(c context.Context, session models.Session) (interface{}, error)
	DeleteById(c context.Context, sessionId string) error
	FindById(c context.Context, sessionId string) (models.Session, error)
}

type SessionRepositoryImpl struct {
	database database.Database
}

func NewSessionRepositoryImpl(db database.Database) SessionRepository {
	return &SessionRepositoryImpl{
		database: db,
	}
}

func (s *SessionRepositoryImpl) Save(c context.Context, session models.Session) (interface{}, error) {
	sessions := s.database.Collection("sessions")
	result, err := sessions.InsertOne(c, session)
	return result, err
}

func (s *SessionRepositoryImpl) DeleteById(c context.Context, sessionId string) error {
	sessions := s.database.Collection("sessions")
	_, err := sessions.DeleteOne(c, bson.M{"_id": sessionId})
	return err
}

func (s *SessionRepositoryImpl) FindById(c context.Context, sessionId string) (models.Session, error) {
	sessions := s.database.Collection("sessions")
	var session models.Session
	err := sessions.FindOne(c, bson.M{"_id": sessionId}).Decode(&session)
	return session, err
}
