package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/controllers"
)

func TaskRoutes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetTasks)
	router.POST("/tasks", controllers.AddTask)
	router.GET("/tasks/:id", controllers.GetTaskByID)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.PUT("/tasks/update/:id", controllers.UpdateTask)
	router.GET("/tasks/status/:status", controllers.GetTasksByStatus)
}
