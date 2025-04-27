package main

import (
	taskHendler "github.com/DanielK_v/taskGrader/handlers/tasks"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/tasks", taskHendler.GetAllTasks)
	router.GET("/tasks/:id", taskHendler.GetTaskById)
	router.POST("/tasks", taskHendler.AddTask)
	router.DELETE("/tasks/:id", taskHendler.DeleteTask)

	router.Run()

}
