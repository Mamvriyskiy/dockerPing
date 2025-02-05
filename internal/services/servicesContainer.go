package services

import (
	"github.com/Mamvriyskiy/dockerPing/internal/models"
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
)

type ContainerService struct {
	repo repository.ContainerPostgresReposipory
}

func NewContainerService(repo repository.ContainerPostgresReposipory) *ContainerService {
	return &ContainerService{repo: repo}
}

func (s *ContainerService) AddContainer(container models.ContainerHandler) (models.ContainerData, error) {
	containerService := models.ContainerService{
		Container: container.Container,
	}

	return s.repo.AddContainer(containerService)
}
