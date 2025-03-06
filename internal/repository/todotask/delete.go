package todotask

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (r *ToDoRepository) Delete(ctx context.Context, id int32) error {
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

	_, err = conn.Exec(ctx, "DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Fatalf("Unable to DELETE: %v\n", err)
		return err

	}
	log.Printf("Field with the %d indicator has been deleted true", id)
	return nil

}
