package service

import (
	"context"
	"forms/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ClientInput struct {
	Firstname  string
	Lastname   string
	Patronomic string
	Position   string
	Company    string
	Phone      string
	Email      string
}

type Clients interface {
	Create(ctx context.Context, client ClientInput) (primitive.ObjectID, error)
}

type Services struct {
	Clients Clients
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	clientService := NewClientService(deps.Repos.Clients)
	return &Services{
		Clients: clientService,
	}
}
