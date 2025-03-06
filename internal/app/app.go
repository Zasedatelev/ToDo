package app

import (
	"context"
	"log"

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

	if a.ServiceProvider == nil {
		log.Println("ServiceProvider is nil")
	}
	if a.ServiceProvider.todoService == nil {
		log.Println("todoService is nil")
	}

	a.Server.Get("/tasks", a.ServiceProvider.todoService.Get)     //!!!! invalid memory address or nil pointer dereference
	a.Server.Post("/tasks", a.ServiceProvider.todoService.Create) //!!!!! invalid memory address or nil pointer dereference

	if err != nil {
		return nil, err
	}

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
func (a *App) initServiceProvider(_ context.Context) error {
	a.ServiceProvider = NewServiceProvider()
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

	// log.Printf("Server is running on %s", a.ServiceProvider.HTTPConfig().Port())

	err := a.Server.Listen("8080")

	if err != nil {
		return err
	}

	return nil
}
