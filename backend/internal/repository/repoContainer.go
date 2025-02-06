package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
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
		logger.Log("Error", "Transaction", "Failed to start transaction:", err)
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
		logger.Log("Error", "Scan", "Error inserting into container table:", err)
		return models.ContainerData{}, err
	}

	logger.Log("Info", "Scan", "Inserted into container table:", err, containerData.ContainerID)

	queriesClientContainer := "INSERT INTO clientcontainer (containerid, clientid) VALUES ($1, $2);"
	_, err = tx.Exec(queriesClientContainer, containerData.ContainerID, clientID)
	if err != nil {
		logger.Log("Error", "Exec", "Error inserting into clientcontainer table:", err)
		return models.ContainerData{}, err
	}
	logger.Log("Info", "Exec", "Successfully inserted into clientcontainer table", err)

	if err = tx.Commit(); err != nil {
		logger.Log("Error", "Transaction", "Failed to commit transaction:", err)
		return models.ContainerData{}, err
	}

	return containerData, nil
}

func (c *ContainerPostgres) GetContainers() ([]models.ContainerData, error) {
	var containersList []models.ContainerData
	queries := "select ipcontainer from container;"
	err := c.db.Select(&containersList, queries)
	if err != nil {
		logger.Log("Error", "Select", "Error get list devices:", err)
		return []models.ContainerData{}, err
	}

	return containersList, nil
}
