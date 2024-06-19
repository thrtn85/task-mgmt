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

/*


func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid task id"})
		return
	}
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, t := range config.Tasks {
		if t.ID == taskID {
			config.Tasks[i] = updatedTask
			config.Tasks[i].ID = taskID
			if err := config.SaveTasksToJSON(); err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "failed to save tasks"})
				return
			}
			c.IndentedJSON(http.StatusOK, config.Tasks[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
*/

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

