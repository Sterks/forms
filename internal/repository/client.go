package repository

import (
	"context"
	"forms/internal/domain"
	"time"

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
		Firstname:  info.Firstname,
		Lastname:   info.Lastname,
		Patronomic: info.Patronomic,
		Position:   info.Position,
		Company:    info.Company,
		Phone:      info.Phone,
		Email:      info.Email,
		CreateAt:   time.Now(),
	})

	return res.InsertedID.(primitive.ObjectID), err
}
