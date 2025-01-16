package dao

import (
	"context"
	"errors"
	"github.com/yourusername/ecommerce-api/auth-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitUserCollection(collection *mongo.Collection) {
	userCollection = collection
}

func CreateUser(ctx context.Context, user models.User) error {
	_, err := userCollection.InsertOne(ctx, user)
	return err
}

func FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.User{}, errors.New("user not found")
		}
		return models.User{}, err
	}
	return user, nil
}