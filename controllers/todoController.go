package controllers

import (
	"github.com/MehdiNaseryPak/todoList/initializers"
	"github.com/MehdiNaseryPak/todoList/models"
	"github.com/gin-gonic/gin"
)

func TodoAll(c *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)
	c.JSON(200, gin.H{
		"message": "success",
		"todos" : todos,
	})
}

func TodoShow(c *gin.Context){
	id := c.Param("id")
	var todo models.Todo
	initializers.DB.First(&todo,id)
	c.JSON(200, gin.H{
		"message": "success",
		"todo" : todo,
	})
}

func TodoCreate(c *gin.Context) {
	var body struct{
		Title string
		Description string
	}

	c.Bind(&body)

	todo := models.Todo{Title: body.Title,Description: body.Description}
	result := initializers.DB.Create(&todo)
	if result.Error != nil {
		c.Status(400)
		return 
	}
	c.JSON(200, gin.H{
		"message": "success",
		"todo" : todo,
	})

}

func TodoUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct{
		Title string
		Description string
	}
	c.Bind(&body)
	var todo models.Todo
	initializers.DB.First(&todo,id)

	initializers.DB.Model(&todo).Updates(models.Todo{Title: body.Title,Description: body.Description})
	c.JSON(200, gin.H{
		"message": "success",
		"todo" : todo,
	})
}

func TodoDelete(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Todo{},id)
	var todos []models.Todo
	initializers.DB.Find(&todos)
	c.JSON(200, gin.H{
		"message": "success",
		"todos" : todos,
	})
}