package services

import (
	"github.com/Mamvriyskiy/dockerPing/backend/internal/repository"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
)

type IClientService interface {
	AddClient(client models.ClientHandler) (models.ClientData, error)
	GenerateToken(client models.ClientHandler) (models.ClientData, string, error)
}

type IContainerService interface {
	AddContainer(container models.ContainerHandler, clientID string) (models.ContainerData, error)
	GetContainers() ([]models.ContainerData, error)
}

type Services struct {
	IClientService
	IContainerService
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IClientService: NewClientService(repo.ClientPostgresReposipory),
		IContainerService: NewContainerService(repo.ContainerPostgresReposipory),
	}
}
