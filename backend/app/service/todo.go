package service

import (
	"github.com/batuhannoz/todo/backend/app/handler"
	"github.com/batuhannoz/todo/backend/app/store/model"
	"github.com/gofiber/fiber/v2"
)

type TodoStore interface {
	AddTodo(todo *model.Todo) *model.Todo
}

type TodoService struct {
	todoStore TodoStore
}

func NewTodoService(store TodoStore) *TodoService {
	return &TodoService{
		todoStore: store,
	}
}

func (t *TodoService) AddTodo(ctx *fiber.Ctx, request handler.TodoRequest) error {
	return nil
}
