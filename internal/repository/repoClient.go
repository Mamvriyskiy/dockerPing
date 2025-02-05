package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

type ClientPostgres struct {
	db *sqlx.DB
}

func NewClientPostgres(db *sqlx.DB) *ClientPostgres {
	return &ClientPostgres{db: db}
}

func (c *ClientPostgres) AddClient(client models.ClientService) (models.ClientData, error) {
	var clientData models.ClientData

	queries := fmt.Sprintf(`insert into client (password, login, email) values ($1, $2, $3) returning clientID, email, login`)

	row := c.db.QueryRow(queries, client.Password, client.Login, client.Email)
	
	err := row.Scan(&clientData.ClientID, &clientData.Email, &clientData.Login)
	if err != nil {
		logger.Log("Error", "Scan", "Error insert into client table:", err)
		return models.ClientData{}, err
	}

	return clientData, nil
}
