package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
)

type ClientPostgresReposipory interface {
	AddClient(client models.ClientService) (models.ClientData, error) 
	GetClient(email string) (models.ClientData, error) 
}

type ContainerPostgresReposipory interface {
	AddContainer(container models.ContainerService, clientID string) (models.ContainerData, error)
	GetContainers() ([]models.ContainerData, error)
}

type Repository struct {
	ClientPostgresReposipory
	ContainerPostgresReposipory
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ClientPostgresReposipory: NewClientPostgres(db),
		ContainerPostgresReposipory: NewContainerPostgres(db),
	}
}
