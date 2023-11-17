package handler

import (
	"final_project_1/dto"
	"final_project_1/pkg/errs"
	"final_project_1/pkg/helpers"
	"final_project_1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) todoHandler {
	return todoHandler{
		todoService: todoService,
	}
}

// GetAllTodos godoc
// @Tags todos
// @Description Get All Todos
// @ID get-all-todos
// @Produce json
// @Success 200 {object} dto.GetAllTodoResponse
// @Router /todos [get]
func (th *todoHandler) GetAllTodos(ctx *gin.Context) {

	result, err := th.todoService.GetAllTodos()

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// CreateTodos godoc
// @Tags todos
// @Description Create a Todo
// @ID create-todo
// @Produce json
// @Param request body dto.CreateTodoRequest true "Todo Request Body"
// @Success 201 {object} dto.CreateTodoResponse
// @Router /todos [post]
func (th *todoHandler) CreateTodo(ctx *gin.Context) {

	var todoRequest dto.CreateTodoRequest

	if err := ctx.ShouldBindJSON(&todoRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := th.todoService.CreateTodo(todoRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, result)

}

// GetTodoById godoc
// @Summary Get a Todo by ID
// @Tags todos
// @Description Get a Todo by ID
// @ID get-todo-by-id
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} dto.GetTodoById
// @Router /todos/{id} [get]
func (th *todoHandler) GetTodoById(ctx *gin.Context) {

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	result, err := th.todoService.GetTodoById(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// UpdateTodo godoc
// @Summary Update a Todo
// @Tags todos
// @Description Update a Todo by ID
// @ID update-todo
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param request body dto.UpdateRequest true "Todo Update Request Body"
// @Success 200 {object} dto.UpdateResponse
// @Router /todos/{id} [put]
func (th *todoHandler) UpdateTodo(ctx *gin.Context) {
	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	var updateRequest dto.UpdateRequest

	if err := ctx.ShouldBindJSON(&updateRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")
		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := th.todoService.UpdateTodo(updateRequest, id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, result)

}

// DeleteTodo godoc
// @Summary Delete a Todo
// @Tags todos
// @Description Delete a Todo by ID
// @ID delete-todo
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200
// @Router /todos/{id} [delete]
func (th *todoHandler) DeleteTodo(ctx *gin.Context) {

	id, err := helpers.GetParamId(ctx, "id")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	err = th.todoService.DeleteTodo(id)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo has been successfully deleted",
	})
}
