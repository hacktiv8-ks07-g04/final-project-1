package service

import (
	"final_project_1/dto"
	"final_project_1/entity"
	"final_project_1/pkg/errs"
	"final_project_1/repository/todo_repository"

	"github.com/asaskevich/govalidator"
)

type todoService struct {
	todoRepository todo_repository.TodoRepository
}

type TodoService interface {
	GetAllTodos() (*[]dto.GetAllTodoResponse, errs.MessageErr)
	CreateTodo(payload dto.CreateTodoRequest) (*dto.CreateTodoResponse, errs.MessageErr)
	GetTodoById(id int) (*dto.GetTodoById, errs.MessageErr)
	UpdateTodo(payload dto.UpdateRequest, id int) (*dto.UpdateResponse, errs.MessageErr)
	DeleteTodo(id int) errs.MessageErr
}

func NewTodoService(todoRepository todo_repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepository,
	}
}

func (ts *todoService) GetAllTodos() (*[]dto.GetAllTodoResponse, errs.MessageErr) {

	todos, err := ts.todoRepository.GetAllTodos()

	if err != nil {
		return nil, err
	}

	var response []dto.GetAllTodoResponse

	for _, todo := range todos {
		todosResponse := todo.TodoToTodoResponses()
		response = append(response, todosResponse)
	}

	return &response, nil

}

func (ts *todoService) CreateTodo(payload dto.CreateTodoRequest) (*dto.CreateTodoResponse, errs.MessageErr) {
	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	todos := &entity.Todo{
		Title: payload.Title,
	}

	err := ts.todoRepository.CreateTodo(todos)

	if err != nil {
		return nil, err
	}

	response := dto.CreateTodoResponse{
		Id:        todos.Id,
		Title:     todos.Title,
		Completed: todos.Completed,
		CreatedAt: todos.CreatedAt,
		UpdatedAt: todos.UpdatedAt,
	}

	return &response, nil

}

func (ts *todoService) GetTodoById(id int) (*dto.GetTodoById, errs.MessageErr) {

	todo := &entity.Todo{
		Id: id,
	}

	err := ts.todoRepository.GetTodoById(todo)

	if err != nil {
		return nil, err
	}

	response := dto.GetTodoById{
		Id:        todo.Id,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &response, nil

}

func (ts *todoService) UpdateTodo(payload dto.UpdateRequest, id int) (*dto.UpdateResponse, errs.MessageErr) {
	_, errv := govalidator.ValidateStruct(payload)

	if errv != nil {
		return nil, errs.NewBadRequest(errv.Error())
	}

	todo := &entity.Todo{
		Id:    id,
		Title: payload.Title,
	}

	err := ts.todoRepository.UpdateTodo(todo)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateResponse{
		Id:        todo.Id,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return &response, nil
}

func (ts *todoService) DeleteTodo(id int) errs.MessageErr {
	err := ts.todoRepository.DeleteTodo(id)

	if err != nil {
		return err
	}

	return nil
}
