package repository

import (
	"fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/jmoiron/sqlx"
)

type ContainerPostgres struct {
	db *sqlx.DB
}

func NewContainerPostgres(db *sqlx.DB) *ContainerPostgres {
	return &ContainerPostgres{db: db}
}

func (c *ContainerPostgres) AddContainer(container models.ContainerService, clientID string) (models.ContainerData, error) {
	var containerData models.ContainerData

	tx, err := c.db.Begin()
	if err != nil {
		logger.Log("Error", "Error starting transaction", err)
		return models.ContainerData{}, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	queriesContainer := "INSERT INTO container (ipcontainer) VALUES ($1) RETURNING containerid, ipcontainer;"
	rowContainer := tx.QueryRow(queriesContainer, container.ContainerIP)

	err = rowContainer.Scan(&containerData.ContainerID, &containerData.ContainerIP)
	if err != nil {
		logger.Log("Error", "Error inserting data into container table", err, fmt.Sprintf("clientID = %s", clientID))
		return models.ContainerData{}, err
	}

	logger.Log("Info", "Data successfully inserted into container table", err, fmt.Sprintf("containerID = %s", containerData.ContainerID))

	queriesClientContainer := "INSERT INTO clientcontainer (containerid, clientid) VALUES ($1, $2);"
	_, err = tx.Exec(queriesClientContainer, containerData.ContainerID, clientID)
	if err != nil {
		logger.Log("Error", "Error inserting data into clientcontainer table", err,
			fmt.Sprintf("clientID = %s, containerID = %s", clientID, containerData.ContainerID))
		return models.ContainerData{}, err
	}
	logger.Log("Info", "Data successfully inserted into clientcontainer table", err,
		fmt.Sprintf("clientID = %s, containerID = %s", clientID, containerData.ContainerID))

	if err = tx.Commit(); err != nil {
		logger.Log("Error", "Error committing transaction", err)
		return models.ContainerData{}, err
	}

	return containerData, nil
}

func (c *ContainerPostgres) GetContainers() ([]models.ContainerData, error) {
	var containersList []models.ContainerData
	queries := "select ipcontainer, containerid from container;"
	err := c.db.Select(&containersList, queries)
	if err != nil {
		logger.Log("Error", "Error selecting list of devices", err)
		return []models.ContainerData{}, err
	}

	return containersList, nil
}
