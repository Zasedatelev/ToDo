package todotask

import (
	"context"
	"fmt"
	"log"

	"github.com/Zasedatelev/ToDo.git/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (r *ToDoRepository) Update(ctx context.Context, task model.Task, id int32) error {

	connString := "postgres://postgres:260616@localhost:5432/tasks?sslmode=disable"
	dbpool, err := pgxpool.New(ctx, connString)

	defer dbpool.Close()
	if err != nil {
		log.Fatalf("no connection to the database", err)
	}
	conn, err := dbpool.Acquire(ctx)
	if err != nil {
		log.Fatalf("Unable to acquire a database connection: %v\n", err)
	}

	defer conn.Release()

	sql := fmt.Sprintf("UPDATE tasks SET title=$1, description=$2, status=$3 WHERE id=$4")
	_, err = conn.Exec(ctx, sql, task.Title, task.Description, task.Status, id)

	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		return err
	}

	return nil

}
