package models

type Session struct {
	Id        string `json:"_id" bson:"_id"`
	UserEmail string `json:"user_email"`
}
