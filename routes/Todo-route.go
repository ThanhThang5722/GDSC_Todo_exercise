package routes

import (
	"github.com/ThanhThang5722/GDSC_Todo_exercise/handlers"
	"github.com/gin-gonic/gin"
)

func TodoRouter(r *gin.RouterGroup) {
	TodoRouter := r.Group("/todo")
	{
		TodoRouter.GET("/all", handlers.GetAllTodos)
		TodoRouter.GET("/", handlers.GetTodoByID)
		TodoRouter.POST("/", handlers.CreateTodo)
		TodoRouter.PUT("/", handlers.UpdateTodo)
		TodoRouter.DELETE("/", handlers.DeleteTodo)
	}
}
