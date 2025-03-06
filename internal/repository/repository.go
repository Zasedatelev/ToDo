package repository

import (
	"context"

	"github.com/Zasedatelev/ToDo.git/model"
)

type Repository interface {
	Get(ctx context.Context) ([]model.Task, error)
	Create(ctx context.Context, task model.Task)
	Update(ctx context.Context, task model.Task, id int32) error
	Delete(ctx context.Context, id int32) error
}
