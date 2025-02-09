package repository

import (
	"fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/jmoiron/sqlx"
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
		logger.Log("Error", "Error inserting data into client table", err, fmt.Sprintf("email = %s", client.Email))
		return models.ClientData{}, err
	}

	logger.Log("Info", "Data successfully inserted into client table", err, fmt.Sprintf("clientID = %s", clientData.ClientID))

	return clientData, nil
}

func (c *ClientPostgres) GetClient(email string) (models.ClientData, error) {
	var clientData models.ClientData

	queries := fmt.Sprintf(`select clientID, login from client where email = $1`)
	err := c.db.Get(&clientData, queries, email)
	if err != nil {
		logger.Log("Error", "Error selecting data from client table", err, fmt.Sprintf("email = %s", email))
		return models.ClientData{}, err
	}

	clientData.Email = email

	logger.Log("Info", "Data successfully retrieved from client table", err, fmt.Sprintf("clientID = %s", clientData.ClientID))

	return clientData, nil
}
