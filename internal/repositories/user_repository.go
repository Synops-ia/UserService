package repositories

import (
	"UserService/internal/database"
	"UserService/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepository interface {
	Save(c context.Context, user models.User) error
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

func (u *UserRepositoryImpl) Save(c context.Context, user models.User) error {
	users := u.database.Collection("users")
	_, err := users.InsertOne(c, user)
	return err
}

func (u *UserRepositoryImpl) FindByEmail(c context.Context, email string) (models.User, error) {
	users := u.database.Collection("users")
	var user models.User
	err := users.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (u *UserRepositoryImpl) insertMany(c context.Context, user map[string]models.User) error {
	usersCollection := u.database.Collection("users")
	var usersInterface []interface{}
	for _, value := range user {
		usersInterface = append(usersInterface, value)
	}
	_, err := usersCollection.InsertMany(c, usersInterface)
	return err
}
