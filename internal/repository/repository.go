package repository

import (
	"context"
	"forms/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client interface {
	Create(ctx context.Context, info domain.Client) (primitive.ObjectID, error)
}

type Repositories struct {
	Client Client
}

func NewRepository(db *mongo.Database) *Repositories {
	return &Repositories{
		Client: NewClientRepo(db),
	}
}
