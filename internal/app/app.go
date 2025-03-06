package app

import (
	"context"

	"github.com/Zasedatelev/ToDo.git/internal/config"
	"github.com/gofiber/fiber"
)

type App struct {
	ServiceProvider *serviceProvider
	Server          *fiber.App
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)

	if err != nil {
		return nil, err
	}

	a.Server.Get("/tasks", a.ServiceProvider.todoService.Get)
	a.Server.Post("/tasks", a.ServiceProvider.todoService.Create)
	a.Server.Delete("/tasks/:id", a.ServiceProvider.todoService.Delete)
	a.Server.Put("/tasks/:id", a.ServiceProvider.todoService.Update)

	return a, nil
}

func (a *App) Run() error {
	return a.runServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initServiceProvider(ctx context.Context) error {
	a.ServiceProvider = NewServiceProvider()
	a.ServiceProvider.pgConfig = a.ServiceProvider.PGConfig()
	a.ServiceProvider.pgPool = a.ServiceProvider.PgPool(ctx)
	a.ServiceProvider.httpConfig = a.ServiceProvider.HTTPConfig()
	a.ServiceProvider.todoRepository = a.ServiceProvider.ToDoRepository(ctx)
	a.ServiceProvider.todoService = a.ServiceProvider.ToDoService(ctx)

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load("/Users/olegzasedatelev/ToDoSkillRock/local.env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {

	a.Server = fiber.New()

	return nil
}

func (a *App) runServer() error {

	err := a.Server.Listen("8080")

	if err != nil {
		return err
	}

	return nil
}
