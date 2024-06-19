package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/backend/initializers"
	"github.com/thrtn85/task-mgmt/backend/models"
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

func GetTasksByStatus(c *gin.Context) {
	status := c.Param("status")
	if status == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "status parameter is required"})
		return
	}

	// Look up the requested task
	var tasks []models.Task
	if err := initializers.DB.Where("status = ?", status).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if len(tasks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found with the given status"})
		return
	}

	// Return the task
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func UpdateTask(c *gin.Context) {
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

	// Bind the JSON body to a map
	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse request body",
		})
		return
	}

	// Ensure ID is not updated
	delete(updateData, "ID")

	// Update the task with the provided fields
	if err := initializers.DB.Model(&task).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update task",
			"details": err.Error(),
		})
		return
	}

	// Return the updated task
	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func DeleteTask(c *gin.Context) {
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

	// Delete the task
	var task models.Task
	if err := initializers.DB.Delete(&task, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted",
	})
}
