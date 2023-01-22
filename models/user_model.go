package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Firstname string             `json:"firstname"`
	Lastname  string             `json:"lastname"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
}
