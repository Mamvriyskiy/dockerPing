package services

import (
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
)


type Services struct {

}

func NewServicesPsql(repo *repository.Repository) *Services {
	return &Services{
	}
}
