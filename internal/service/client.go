package service

import (
	"context"
	"forms/internal/domain"
	"forms/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientService struct {
	repo repository.Clients
}

func NewClientService(repo repository.Clients) *ClientService {
	return &ClientService{
		repo: repo,
	}
}

func (c *ClientService) Create(ctx context.Context, client ClientInput) (primitive.ObjectID, error) {

	var dtoClient domain.Client
	dtoClient.Firstname = client.Firstname
	dtoClient.Lastname = client.Lastname
	dtoClient.Patronomic = client.Patronomic
	dtoClient.Company = client.Company
	dtoClient.Email = client.Email
	dtoClient.Phone = client.Phone
	dtoClient.Position = client.Position

	id, err := c.repo.Create(ctx, dtoClient)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return id, err
}
