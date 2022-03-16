package repository

import (
	"context"
	"forms/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepo struct {
	db *mongo.Collection
}

func NewClientRepo(db *mongo.Database) *ClientRepo {
	return &ClientRepo{
		db: db.Collection(clientsCollection),
	}
}

func (c *ClientRepo) Create(ctx context.Context, info domain.Client) (primitive.ObjectID, error) {
	res, err := c.db.InsertOne(ctx, domain.Client{
		// TODO Добавить время
	})

	return res.InsertedID.(primitive.ObjectID), err
}
