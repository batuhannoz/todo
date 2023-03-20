package store

import (
	"github.com/batuhannoz/todo/backend/app/store/model"
	"gorm.io/gorm"
)

type TodoStore struct {
	connection *gorm.DB
}

func NewTodoStore(connection *gorm.DB) *TodoStore {
	return &TodoStore{
		connection: connection,
	}
}

func (t *TodoStore) AddTodo(todo *model.Todo) *model.Todo {
	t.connection.Create(&todo)
	return todo
}
