package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/thrtn85/backend/task-mgmt/middleware"
	"github.com/thrtn85/task-mgmt/backend/initializers"
	"github.com/thrtn85/task-mgmt/backend/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	// Use the CORS middleware
	router.Use(middleware.CORSConfig())

	router.OPTIONS("/*path", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // Debugging: Replace with specific origin
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Header("Access-Control-Max-Age", "86400")
		c.Status(http.StatusNoContent)
	})

	// Serve static files from the "frontend" directory
	router.Static("/app", "../frontend")

	// Serve HTML files from the "frontend" directory
	router.GET("/html/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.File("../frontend/" + name + ".html")
	})

	routes.UserRoutes(router)

	router.GET("/dashboard", middleware.RequireAuth, func(c *gin.Context) {
		// Access user from context if needed
		user, _ := c.Get("user")
		c.JSON(200, gin.H{
			"message": "You are authorized to access dashboard",
			"user":    user, // Example: send user details if needed
		})
	})

	routes.TaskRoutes(router)

	// Start the server on port 5500
	if err := router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
