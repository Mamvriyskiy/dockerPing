package services

import (
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
)

type IClientService interface {
	AddClient(models.ClientHandler) (models.ClientData, error)
}


type Services struct {
	IClientService
}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
		IClientService: NewClientService(repo.ClientPostgresReposipory),
	}
}
