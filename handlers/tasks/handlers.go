package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/DanielK_v/taskGrader/models"
)

func GetAllTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Failed to fetch tasks"},
		)
		return
	}

	context.JSON(http.StatusOK, tasks)
}

func GetTaskById(context *gin.Context) {
	id, inputError := strconv.ParseUint(string(context.Param("id")), 10, 64)
	if inputError != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid task ID"},
		)
		return
	}

	task, err := models.GetTaskById(id)

	if err == sql.ErrNoRows {
		context.JSON(
			http.StatusNotFound,
			gin.H{"error": "Task not found"},
		)
		return
	}

	if task == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	context.JSON(http.StatusOK, task)
}

func AddTask(context *gin.Context) {
	var task models.Task
	if err := context.ShouldBindJSON(&task); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	newTask, err := models.AddTask(task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}
	

	context.JSON(http.StatusCreated, gin.H{"task": newTask})
}

func DeleteTask(context *gin.Context) {
	id, inputError := strconv.ParseUint(string(context.Param("id")), 10, 64)
	if inputError != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Invalid task ID"},
		)
		return
	}

	err := models.DeleteTask(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong! Could not delete task"})
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}
