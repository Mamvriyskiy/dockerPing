package services

import (
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/repository"
)

type IClientService interface {
	AddClient(client models.ClientHandler) (models.ClientData, error)
	GenerateToken(client models.ClientHandler) (models.ClientData, string, error)
}

type IContainerService interface {
	AddContainer(container models.ContainerHandler, clientID string) (models.ContainerData, error)
	GetContainers() ([]models.ContainerData, error)
}

type IHistoryService interface {
	AddContainersStatus(containers []models.HistoryHandler) error
}

type Services struct {
	IClientService
	IContainerService
	IHistoryService
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IClientService:    NewClientService(repo.ClientPostgresReposipory),
		IContainerService: NewContainerService(repo.ContainerPostgresReposipory),
		IHistoryService:   NewHistoryService(repo.HistoryPostgresRepository),
	}
}
