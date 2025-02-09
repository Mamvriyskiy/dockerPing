package models

import (
	"time"
)

type HistoryHandler struct {
	ContainerID string    `json:"containerid"`
	Status      string    `json:"status"`
	TimePing    time.Time `json:"timeping"`
}

type HistoryService struct {
	ContainerID string    `db:"containerid"`
	Status      string    `db:"status"`
	TimePing    time.Time `db:"timeping"`
}

type HistoryData struct {
	ContainerID string    `db:"containerid" json:"containerid"`
	Status      string    `db:"status" json:"status"`
	TimePing    time.Time `db:"timeping" json:"timeping"`
}
