package repository

import (
	// "fmt"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/jmoiron/sqlx"
)

type HistoryPostgres struct {
	db *sqlx.DB
}

func NewHistoryPostgres(db *sqlx.DB) *HistoryPostgres {
	return &HistoryPostgres{db: db}
}

func (c *HistoryPostgres) AddContainersStatus(containers []models.HistoryService) error {
	for _, cnt := range containers {
		queriesAddHistory := "INSERT INTO historycontainer (containerid, statusping, timeping) VALUES ($1, $2, $3);"
		_, err := c.db.Exec(queriesAddHistory, cnt.ContainerID, cnt.Status, cnt.TimePing)
		if err != nil {
			logger.Log("Error", "Error inserting into historycontainer table:", err)
			return err
		}
	}

	return nil
}
