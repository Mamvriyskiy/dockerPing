package services

import (
	"github.com/Mamvriyskiy/dockerPing/backend/internal/models"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/repository"
)

type HistoryService struct {
	repo repository.HistoryPostgresRepository
}

func NewHistoryService(repo repository.HistoryPostgresRepository) *HistoryService {
	return &HistoryService{repo: repo}
}

func (s *HistoryService) AddContainersStatus(containers []models.HistoryHandler) error {
	containersServ := make([]models.HistoryService, len(containers))
	for i, cnt := range containers {
		containersServ[i] = models.HistoryService{
			ContainerID: cnt.ContainerID,
			Status:      cnt.Status,
			TimePing:    cnt.TimePing,
		}
	}

	return s.repo.AddContainersStatus(containersServ)
}
