package todotask

import (
	"github.com/Zasedatelev/ToDo.git/internal/repository"
	"github.com/Zasedatelev/ToDo.git/internal/service"
)

type serv struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) service.ToDoService {
	return &serv{
		repo: repo,
	}
}
