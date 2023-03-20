package app

import (
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

	// TODO create store with db || todoStore := store.NewTodoStore(a.db)

	// TODO create service with store || todoService := service.NewTodoService(todoStore)

	// TODO create handler with service || appHandler := handler.NewHandler(todoService)

	a.fiber = fiber.New()
	a.fiber.Use(cors.New())
	// a.registerRoutes(appHandler)
	return nil
}

/*
func (a *App) registerRoutes(handler *handler.Handler) {
	a.fiber.Post("/todo", handler.AddTodo)
}
*/

func (a *App) Run(host string) error {
	return a.fiber.Listen(host)
}
