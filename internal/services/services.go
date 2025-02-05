package services

import (
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

type IClientService interface {
	AddClient(client models.ClientHandler) (models.ClientData, error)
	GenerateToken(client models.ClientHandler) (models.ClientData, string, error)
}

type IContainerService interface {
	AddContainer(container models.ContainerHandler) (models.ContainerData, error)
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
