package services

import (
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/repository"
)

type ContainerService struct {
	repo repository.ContainerPostgresReposipory
}

func NewContainerService(repo repository.ContainerPostgresReposipory) *ContainerService {
	return &ContainerService{repo: repo}
}

func (s *ContainerService) AddContainer(container models.ContainerHandler, clientID string) (models.ContainerData, error) {
	containerService := models.ContainerService{
		Container: container.Container,
	}

	return s.repo.AddContainer(containerService, clientID)
}

func (s *ContainerService) GetContainers() ([]models.ContainerData, error) {
	return s.repo.GetContainers()
}
