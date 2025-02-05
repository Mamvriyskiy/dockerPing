package repository

import (
	// "fmt"
	"github.com/jmoiron/sqlx"
	// "github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

type ContainerPostgres struct {
	db *sqlx.DB
}

func NewContainerPostgres(db *sqlx.DB) *ContainerPostgres {
	return &ContainerPostgres{db: db}
}

func (c *ContainerPostgres) AddContainer(client models.ContainerService) (models.ContainerData, error) {
	var containerData models.ContainerData

	return containerData, nil
}
