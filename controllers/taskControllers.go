package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/config"
	"github.com/thrtn85/task-mgmt/models"
)

func GetTasks(requestContext *gin.Context) {
	requestContext.IndentedJSON(http.StatusOK, config.Tasks)
}

func GetTaskByID(requestContext *gin.Context) {
	id := requestContext.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid task id"})
		return
	}
	for _, t := range config.Tasks {
		if t.ID == taskID {
			requestContext.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	requestContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func GetTasksByStatus(requestContext *gin.Context) {
	status := requestContext.Param("status")
	if status == "" {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"message": "status parameter is required"})
		return
	}
	var tasksWithStatus []models.Task
	for _, t := range config.Tasks {
		if t.Status == status {
			tasksWithStatus = append(tasksWithStatus, t)
		}
	}
	if len(tasksWithStatus) > 0 {
		requestContext.IndentedJSON(http.StatusOK, tasksWithStatus)
	} else {
		requestContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "no tasks found with the given status"})
	}
}

func AddTask(requestContext *gin.Context) {
	var newTask models.Task
	if err := requestContext.ShouldBindJSON(&newTask); err != nil {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	maxID := 0
	for _, t := range config.Tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	newTask.ID = maxID + 1
	config.Tasks = append(config.Tasks, newTask)
	if err := config.SaveTasksToJSON(); err != nil {
		requestContext.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to save tasks"})
		return
	}
	requestContext.JSON(http.StatusCreated, newTask)
}

func UpdateTask(requestContext *gin.Context) {
	id := requestContext.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid task id"})
		return
	}
	var updatedTask models.Task
	if err := requestContext.ShouldBindJSON(&updatedTask); err != nil {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, t := range config.Tasks {
		if t.ID == taskID {
			config.Tasks[i] = updatedTask
			config.Tasks[i].ID = taskID
			if err := config.SaveTasksToJSON(); err != nil {
				requestContext.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to save tasks"})
				return
			}
			requestContext.IndentedJSON(http.StatusOK, config.Tasks[i])
			return
		}
	}
	requestContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func DeleteTask(requestContext *gin.Context) {
	id := requestContext.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid task id"})
		return
	}
	for i, t := range config.Tasks {
		if t.ID == taskID {
			config.Tasks = append(config.Tasks[:i], config.Tasks[i+1:]...)
			if err := config.SaveTasksToJSON(); err != nil {
				requestContext.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to save tasks"})
				return
			}
			requestContext.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted"})
			return
		}
	}
	requestContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
