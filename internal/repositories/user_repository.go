package repositories

import (
	"UserService/internal/database"
	"UserService/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	Save(c context.Context, user models.User) (interface{}, error)
	FindByEmail(c context.Context, email string) (models.User, error)
}

type UserRepositoryImpl struct {
	database database.Database
}

func NewUserRepositoryImpl(db database.Database) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		database: db,
	}
}

func (u *UserRepositoryImpl) Save(c context.Context, user models.User) (interface{}, error) {
	users := u.database.Collection("users")
	insertedId, err := users.InsertOne(c, user)
	return insertedId, err
}

func (u *UserRepositoryImpl) FindByEmail(c context.Context, email string) (models.User, error) {
	users := u.database.Collection("users")
	var user models.User
	err := users.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}
