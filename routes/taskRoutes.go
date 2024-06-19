package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/controllers"
	"github.com/thrtn85/task-mgmt/middleware"
)

func TaskRoutes(router *gin.Engine) {
	router.GET("/tasks", middleware.RequireAuth, controllers.GetTasks)
	router.POST("/tasks", middleware.RequireAuth, controllers.CreateTask)
	router.GET("/tasks/:id", middleware.RequireAuth, controllers.GetTaskByID)
	router.GET("/tasks/status/:status", middleware.RequireAuth, controllers.GetTasksByStatus)
	router.DELETE("/tasks/:id", middleware.RequireAuth, controllers.DeleteTask)
	/* router.PUT("/tasks/update/:id", middleware.RequireAuth, controllers.UpdateTask)
	*/
}
