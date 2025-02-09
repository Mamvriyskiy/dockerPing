package repository

import (
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/jmoiron/sqlx"
)

type ClientPostgresReposipory interface {
	AddClient(client models.ClientService) (models.ClientData, error)
	GetClient(email string) (models.ClientData, error)
}

type ContainerPostgresReposipory interface {
	AddContainer(container models.ContainerService, clientID string) (models.ContainerData, error)
	GetContainers() ([]models.ContainerData, error)
}
type HistoryPostgresRepository interface {
	AddContainersStatus(container []models.HistoryService) error
}

type Repository struct {
	ClientPostgresReposipory
	ContainerPostgresReposipory
	HistoryPostgresRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ClientPostgresReposipory:    NewClientPostgres(db),
		ContainerPostgresReposipory: NewContainerPostgres(db),
		HistoryPostgresRepository:   NewHistoryPostgres(db),
	}
}
