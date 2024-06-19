package main

import (
	"github.com/gin-gonic/gin"
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

	routes.TaskRoutes(router)
	routes.UserRoutes(router)

	router.Run()
}
