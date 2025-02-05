package services

import (
	"github.com/Mamvriyskiy/dockerPing/internal/models"
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
)

type ClientService struct {
	repo repository.ClientPostgresReposipory
}

func NewClientService(repo repository.ClientPostgresReposipory) *ClientService {
	return &ClientService{repo: repo}
}

func (s *ClientService) AddClient(client models.ClientHandler) (models.ClientData, error) {
	clientServ := models.ClientService{
		Client: client.Client,
		Password: client.Password,
	}

	return s.repo.AddClient(clientServ)
}