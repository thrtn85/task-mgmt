package initializers

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	dbSourceName := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(sqlite.Open(dbSourceName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
