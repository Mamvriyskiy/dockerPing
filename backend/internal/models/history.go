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
	ContainerID string
	Status      string
	TimePing    time.Time
}

type HistoryData struct {
	ContainerIP string `db:"ipcontainer" json:"ipcontainer"`
	ContainerID string `db:"containerid" json:"containerid"`
	Status      string `db:"statusping" json:"status"`
	TimePing    string `db:"timeping" json:"timeping"`
}
