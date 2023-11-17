package todo_pg

import (
	"errors"
	"final_project_1/entity"
	"final_project_1/pkg/errs"
	"final_project_1/repository/todo_repository"

	"gorm.io/gorm"
)

type todoPG struct {
	db *gorm.DB
}

func NewTodoPG(db *gorm.DB) todo_repository.TodoRepository {
	return &todoPG{
		db: db,
	}
}

func (t *todoPG) GetAllTodos() ([]entity.Todo, errs.MessageErr) {
	var todos []entity.Todo
	result := t.db.Find(&todos).Error

	if result != nil {
		return nil, errs.NewInternalServerError("something Went Wrong")
	}

	return todos, nil
}

func (t *todoPG) CreateTodo(todo *entity.Todo) errs.MessageErr {

	err := t.db.Create(&todo).Error

	if err != nil {
		return errs.NewInternalServerError("something Went Wrong")
	}

	return nil
}

func (t *todoPG) GetTodoById(todo *entity.Todo) errs.MessageErr {
	err := t.db.First(&todo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.NewNotFoundError("Not found")
		}
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (t *todoPG) UpdateTodo(todo *entity.Todo) errs.MessageErr {
	title := todo.Title

	result := t.db.First(&todo, todo.Id)
	if result.Error != nil {
		return errs.NewNotFoundError("not found")
	}

	result = t.db.Model(&todo).Update("title", title)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}

func (t *todoPG) DeleteTodo(id int) errs.MessageErr {
	result := t.db.First(&entity.Todo{}, id)
	if result.Error != nil {
		return errs.NewNotFoundError("not found")
	}

	result = t.db.Delete(&entity.Todo{}, id)

	if result.Error != nil {
		return errs.NewInternalServerError("Internal Server Error")
	}

	return nil
}
