package repository

import (
	"context"
	"forms/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Clients interface {
	Create(ctx context.Context, info domain.Client) (primitive.ObjectID, error)
}

type Repositories struct {
	Clients Clients
}

func NewRepository(db *mongo.Database) *Repositories {
	return &Repositories{
		Clients: NewClientRepo(db),
	}
}
