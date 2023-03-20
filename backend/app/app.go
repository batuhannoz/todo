package app

import (
	"github.com/batuhannoz/todo/backend/app/handler"
	"github.com/batuhannoz/todo/backend/app/service"
	"github.com/batuhannoz/todo/backend/app/store"
	"github.com/batuhannoz/todo/backend/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

type App struct {
	fiber *fiber.App
	db    *gorm.DB
}

func (a *App) Initialize(config *config.Config) error {
	var err error
	a.db, err = store.NewMySQL(config.MySQL)
	if err != nil {
		return err
	}

	todoStore := store.NewTodoStore(a.db)

	todoService := service.NewTodoService(todoStore)

	appHandler := handler.NewAppHandler(todoService)

	a.fiber = fiber.New()
	a.fiber.Use(cors.New())
	a.registerRoutes(appHandler)
	return nil
}

func (a *App) registerRoutes(handler *handler.AppHandler) {
	a.fiber.Post("/todo", handler.AddTodo)
}

func (a *App) Run(host string) error {
	return a.fiber.Listen(host)
}
