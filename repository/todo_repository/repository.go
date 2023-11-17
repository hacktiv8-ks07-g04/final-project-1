package todo_repository

import (
	"final_project_1/entity"
	"final_project_1/pkg/errs"
)

type TodoRepository interface {
	GetAllTodos() ([]entity.Todo, errs.MessageErr)
	CreateTodo(todo *entity.Todo) errs.MessageErr
	GetTodoById(todo *entity.Todo) errs.MessageErr
	UpdateTodo(todo *entity.Todo) errs.MessageErr
	DeleteTodo(id int) errs.MessageErr
}
