package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	ProductID primitive.ObjectID `bson:"product_id" json:"product_id"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Status    string             `bson:"status" json:"status"`
}