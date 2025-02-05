package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) AddClient(client models.ClientService) (models.ClientData, error) {
	fmt.Println("PS: ", client)
	return models.ClientData{}, nil
}
