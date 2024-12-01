package main

import (
	"github.com/ThanhThang5722/GDSC_Todo_exercise/pkg/database"
	"github.com/ThanhThang5722/GDSC_Todo_exercise/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.GetDbInstance()
	defer db.Close()

	myRouter := gin.Default()

	apiGroup := myRouter.Group("/api")

	routes.TodoRouter(apiGroup)

	myRouter.Run()
}
