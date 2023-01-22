package repository

import (
	"auth-rest-api/configs"
	"auth-rest-api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	c *mongo.Collection
}

type UserRepository interface {
	Save(c context.Context, user models.User) (string, error)
	GetAll(c context.Context) ([]*models.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepository{
		c: (*mongo.Collection)(configs.ConnectDB().Database("go-db").Collection("users")),
	}
}

// GetAll implements UserRepository
func (r *userRepository) GetAll(c context.Context) ([]*models.User, error) {
	cursor, err := r.c.Find(c, bson.M{})

	userList := make([]*models.User, 0)
	cursor.All(c, &userList)

	return userList, err
}

// Save implements UserRepository
func (r *userRepository) Save(c context.Context, user models.User) (string, error) {
	_, err := r.c.InsertOne(c, user)
	return "success to create an account", err
}
