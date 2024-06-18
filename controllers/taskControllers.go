package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/initializers"
	"github.com/thrtn85/task-mgmt/models"
)

func GetTasks(c *gin.Context) {
	var task []models.Task
	// Fetch all task from the database
	initializers.DB.Find(&task)
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var body struct {
		Title       string
		Description string
		DueDate     time.Time
		Status      string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Create the task
	task := models.Task{Title: body.Title, Description: body.Description, DueDate: body.DueDate, Status: body.Status}
	result := initializers.DB.Create(&task)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create task",
		})
		return
	}
	// Respond
	c.JSON(http.StatusOK, gin.H{
		"message": task,
	})
}

func GetTaskByID(c *gin.Context) {
	// Extract ID from URL parameters
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID parameter is required",
		})
		return
	}

	// Convert ID to uint
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID format",
		})
		return
	}

	// Look up the requested task
	var task models.Task
	if err := initializers.DB.First(&task, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	// Return the task
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

/*
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
*/
