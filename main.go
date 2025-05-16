package main

import (
	"github.com/MehdiNaseryPak/todoList/controllers"
	"github.com/MehdiNaseryPak/todoList/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}
func main() {
	router := gin.Default()
	router.GET("/todos", controllers.TodoAll)
	router.GET("/todos/:id", controllers.TodoShow)
	router.POST("/todos", controllers.TodoCreate)
	router.PUT("/todos/:id", controllers.TodoUpdate)
	router.DELETE("/todos/:id", controllers.TodoDelete)
	router.Run()
}
