package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var tasks []task

type task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Status      string    `json:"status"`
}

func init() {
	// read data from JSON file
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}
	// Unmarshal JSON data into tasks slice
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}
}

// list of all tasks as JSON
func getTasks(requestContext *gin.Context) {
	requestContext.IndentedJSON(http.StatusOK, tasks)
}

// get a specific task
func getTaskByID(requestContext *gin.Context) {
	id := requestContext.Param("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		requestContext.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid task id"})
		return
	}
	// loop over the list of tasks to find the id that matches the parameter
	for _, t := range tasks {
		if t.ID == taskID {
			requestContext.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	requestContext.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// add a new tasks
func addTask(requestContext *gin.Context) {
	var newTask task

	// BIND the JSON input to newTask
	if err := requestContext.ShouldBindJSON(&newTask); err != nil {
		requestContext.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	// generate a new ID
	newTask.ID = len(tasks) + 1

	// add the new task to the list
	tasks = append(tasks, newTask)
	requestContext.JSON(http.StatusCreated, newTask)
}

/* TODO-add getTasksByStatus */

func main() {
	router := gin.Default()

	router.GET("/tasks", getTasks)
	router.POST("/tasks", addTask)
	router.GET("/tasks/:id", getTaskByID)

	router.Run("localhost: 8020")
}
