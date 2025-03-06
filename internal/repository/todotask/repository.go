package todotask

import (
	"github.com/Zasedatelev/ToDo.git/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ToDoRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.Repository {
	return &ToDoRepository{
		db: db,
	}
}
