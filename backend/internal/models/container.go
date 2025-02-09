package models

type ContainerHandler struct {
	ContainerIP string `json:"ipcontainer"`
}

type ContainerService struct {
	ContainerIP string `db:"ipcontainer"`
}

type ContainerData struct {
	ContainerIP string `db:"containerip" json:"ipcontainer"`
	ContainerID string `db:"containerid" json:"containerid"`
}
