package main

import (
	"log"

	taskHendler "github.com/DanielK_v/taskGrader/handlers/tasks"
	"github.com/DanielK_v/taskGrader/services/database"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer database.Db.Close()

	router := gin.Default()

	router.GET("/tasks", taskHendler.GetAllTasks)
	router.GET("/tasks/:id", taskHendler.GetTaskById)
	router.POST("/tasks", taskHendler.AddTask)
	router.DELETE("/tasks/:id", taskHendler.DeleteTask)

	router.Run()

}
