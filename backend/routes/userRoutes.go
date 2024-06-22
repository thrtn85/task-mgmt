package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/backend/controllers"
	"github.com/thrtn85/task-mgmt/backend/middleware"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.GET("/users", middleware.RequireAuth, controllers.GetUsers)
	router.GET("/users", middleware.RequireAuth, controllers.GetUsers)
}