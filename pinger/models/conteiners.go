package models

import (
	"time"
)

type Container struct {
	ContainerID string    `json:"containerid"`
	ContainerIP string    `json:"ipcontainer"`
	Status      string    `json:"status"`
	TimePing    time.Time `json:"timeping"`
}
