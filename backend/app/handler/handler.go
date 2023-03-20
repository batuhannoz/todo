package handler

import "github.com/gofiber/fiber/v2"

type TodoService interface {
	AddTodo(ctx *fiber.Ctx, request TodoRequest) error
}

type AppHandler struct {
	todoService TodoService
}

func NewAppHandler(todoService TodoService) *AppHandler {
	return &AppHandler{
		todoService: todoService,
	}
}

func (a *AppHandler) AddTodo(ctx *fiber.Ctx) error {
	return nil
}
