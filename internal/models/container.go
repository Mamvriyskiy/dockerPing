package models

type Container struct {
	ContainerID string `db:"containerid"`
	ContainerIP string `db:"ipcontainer" json:"ipcontainer"`
}

type ContainerHandler struct {
	Container
}

type ContainerService struct {
	Container
}

type ContainerData struct {
	Container
}
