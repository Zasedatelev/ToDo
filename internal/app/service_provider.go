package app

import (
	"context"
	"log"

	"github.com/Zasedatelev/ToDo.git/internal/config"
	"github.com/Zasedatelev/ToDo.git/internal/repository"
	todorepository "github.com/Zasedatelev/ToDo.git/internal/repository/todotask"
	"github.com/Zasedatelev/ToDo.git/internal/service"
	todoservice "github.com/Zasedatelev/ToDo.git/internal/service/todotask"
	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceProvider struct {
	pgConfig       config.PGConfig
	pgPool         *pgxpool.Pool
	httpConfig     config.HTTPConfig
	todoRepository repository.Repository
	todoService    service.ToDoService
}

func NewServiceProvider() *serviceProvider {

	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig != nil {
		ctf, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = ctf
	}

	return s.pgConfig

}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig != nil {
		http, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = http
	}

	return s.httpConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.New(ctx, "postgres://postgres:260616@localhost:5432/tasks?sslmode=disable")
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.pgPool = pool
	}

	return s.pgPool

}

func (s *serviceProvider) ToDoRepository(ctx context.Context) repository.Repository {
	if s.todoRepository == nil {
		s.todoRepository = todorepository.NewRepository(s.PgPool(ctx))
	}

	return s.todoRepository
}

func (s *serviceProvider) ToDoService(ctx context.Context) service.ToDoService {
	if s.todoService == nil {
		s.todoService = todoservice.NewService(s.ToDoRepository(ctx))
	}

	return s.todoService
}
