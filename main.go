package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thrtn85/task-mgmt/config"
	"github.com/thrtn85/task-mgmt/routes"
)

func main() {
	config.Init()

	router := gin.Default()
	routes.TaskRoutes(router)

	router.Run("localhost: 8080")
}
