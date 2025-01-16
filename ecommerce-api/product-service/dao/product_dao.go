package dao

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ecommerce-api/product-service/models"
)

var productCollection *mongo.Collection

func InitProductDAO(db *mongo.Database) {
	productCollection = db.Collection("products")
}

func CreateProduct(ctx context.Context, product *models.Product) error {
	product.ID = uuid.New().String()
	_, err := productCollection.InsertOne(ctx, product)
	return err
}

func FindProductByID(ctx context.Context, id string) (*models.Product, error) {
	var product models.Product
	err := productCollection.FindOne(ctx, bson.M{"id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func DeleteProduct(ctx context.Context, id string) error {
	_, err := productCollection.DeleteOne(ctx, bson.M{"id": id})
	return err
}