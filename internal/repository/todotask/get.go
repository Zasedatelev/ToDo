package todotask

import (
	"context"
	"log"

	"github.com/Zasedatelev/ToDo.git/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (r *ToDoRepository) Get(ctx context.Context) ([]model.Task, error) {
	connString := "postgres://postgres:260616@localhost:5432/tasks?sslmode=disable"
	dbpool, err := pgxpool.New(ctx, connString)
	task := model.Task{}
	defer dbpool.Close()

	if err != nil {
		log.Fatalf("no connection to the database", err)
	}

	rows, err := dbpool.Query(ctx, "select * from tasks")
	if err != nil {
		return nil, err
	}

	tasks := []model.Task{}

	for rows.Next() {
		err = rows.Scan(task.TaskId, &task.Description, &task.Status)
		if err != nil {
			log.Fatalf("scan failed: %v\n", err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
