package todotask

import (
	"context"
	"log"

	"github.com/Zasedatelev/ToDo.git/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (r *ToDoRepository) Create(ctx context.Context, task model.Task) {

	connString := "postgres://postgres:260616@localhost:5432/tasks?sslmode=disable"
	dbpool, err := pgxpool.New(ctx, connString)
	defer dbpool.Close()

	if err != nil {
		log.Fatalf("No connection to the database", err)
	}

	sql := "INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)"

	_, err = dbpool.Exec(ctx, sql, task.Title, task.Description, task.Status)

	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
	}

}
