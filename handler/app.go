package handler

import (
	"final_project_1/database"
	"final_project_1/docs"
	"final_project_1/repository/todo_repository/todo_pg"
	"final_project_1/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {

	database.HandleDatabaseConnection()
	db := database.GetDatabaseInstance()

	todoRepo := todo_pg.NewTodoPG(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := NewTodoHandler(todoService)

	port := "4000"

	r := gin.Default()

	docs.SwaggerInfo.Title = "Final Project 1"
	docs.SwaggerInfo.Description = "Final Project 1 Hactiv8 by Kelompok 4"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:4000/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	todoRoute := r.Group("/todos")
	{

		todoRoute.GET("", todoHandler.GetAllTodos)
		todoRoute.POST("", todoHandler.CreateTodo)
		todoRoute.GET("/:id", todoHandler.GetTodoById)
		todoRoute.PUT("/:id", todoHandler.UpdateTodo)
		todoRoute.DELETE("/:id", todoHandler.DeleteTodo)
	}
	r.Run("127.0.0.1:" + port)
}
