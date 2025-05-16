package main

import (
	"github.com/MehdiNaseryPak/todoList/initializers"
	"github.com/MehdiNaseryPak/todoList/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectionDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.Todo{})
}