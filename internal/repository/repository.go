package repository

import (
	"context"

	"github.com/Zasedatelev/ToDo.git/model"
)

type Repository interface {
	Get(ctx context.Context) ([]model.Task, error)
	Create(ctx context.Context, task model.Task)
	// Update(id int) error
	// Delete(id int) error
}
