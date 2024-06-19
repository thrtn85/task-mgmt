package initializers

import (
	"github.com/thrtn85/task-mgmt/backend/models"
)

func SyncDatabase() {
	// Migrate the schema
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Task{})
}
