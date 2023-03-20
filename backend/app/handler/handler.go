package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type TodoService interface {
	AddTodo(ctx *fiber.Ctx, request TodoRequest) error
}

type AppHandler struct {
	todoService TodoService
	validator   validator.Validate
}

func NewAppHandler(todoService TodoService) *AppHandler {
	return &AppHandler{
		todoService: todoService,
		validator:   validator.Validate{},
	}
}

func (a *AppHandler) AddTodo(ctx *fiber.Ctx) error {
	var req TodoRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	err := a.validator.Struct(req)
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	if err = a.todoService.AddTodo(ctx, req); err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusOK)
}
