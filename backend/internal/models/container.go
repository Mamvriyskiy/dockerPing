package models

type ContainerHandler struct {
	ContainerIP string `json:"ipcontainer"`
}

type ContainerService struct {
	ContainerIP string
}

type ContainerData struct {
	ContainerIP string `db:"ipcontainer" json:"ipcontainer"`
	ContainerID string `db:"containerid" json:"containerid"`
}
