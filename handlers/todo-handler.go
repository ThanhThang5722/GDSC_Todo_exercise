package handlers

import (
	"net/http"

	"github.com/ThanhThang5722/GDSC_Todo_exercise/models"
	"github.com/gin-gonic/gin"
)

func GetAllTodos(ctx *gin.Context) {
	resulst, err := models.GetAllTodos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": "Fail to Get Data",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Data": resulst,
	})
}

func UpdateTodo(ctx *gin.Context) {
	var todo models.Todo
	err := ctx.ShouldBindBodyWithJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": "Fail to Update Data",
		})
		return
	}
	err = todo.Update()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": "Fail to Update Data",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": "update successfuly",
	})
}

func DeleteTodo(ctx *gin.Context) {
	todoID := ctx.Query("id")
	err := models.DeleteTodo(todoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status": "Fail to Delete Todo",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": "Delete Succesfully",
	})
}

func CreateTodo(ctx *gin.Context) {
	var newTodo models.NewTodo
	err := ctx.ShouldBindBodyWithJSON(&newTodo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": "Not valid",
		})
		return
	}
	err = newTodo.SaveNewTodo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": "Not valid",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Status": "Create new Todo Succesfully",
	})
}

func GetTodoByID(ctx *gin.Context) {
	id := ctx.Query("id")
	todo, err := models.GetTodoByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status": "Couldn't get the data",
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"Data": todo,
	})
}
