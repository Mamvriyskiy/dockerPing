package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

type ClientPostgresReposipory interface {
	AddClient(client models.ClientService) (models.ClientData, error) 
}

type Repository struct {
	ClientPostgresReposipory
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ClientPostgresReposipory: NewClientPostgres(db),
	}
}
