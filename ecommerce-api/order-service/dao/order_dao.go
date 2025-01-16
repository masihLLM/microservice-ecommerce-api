package dao

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"ecommerce-api/order-service/models"
)

type OrderDAO struct {
	collection *mongo.Collection
}

func NewOrderDAO(client *mongo.Client, dbName string) *OrderDAO {
	collection := client.Database(dbName).Collection("orders")
	return &OrderDAO{collection: collection}
}

func (dao *OrderDAO) CreateOrder(ctx context.Context, order *models.Order) error {
	order.CreatedAt = time.Now()
	_, err := dao.collection.InsertOne(ctx, order)
	return err
}

func (dao *OrderDAO) FindOrderByID(ctx context.Context, id string) (*models.Order, error) {
	var order models.Order
	err := dao.collection.FindOne(ctx, bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (dao *OrderDAO) FindOrdersByUserID(ctx context.Context, userID string) ([]models.Order, error) {
	var orders []models.Order
	cursor, err := dao.collection.Find(ctx, bson.M{"userid": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order models.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (dao *OrderDAO) UpdateOrder(ctx context.Context, id string, order *models.Order) error {
	_, err := dao.collection.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": order})
	return err
}

func (dao *OrderDAO) DeleteOrder(ctx context.Context, id string) error {
	_, err := dao.collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}