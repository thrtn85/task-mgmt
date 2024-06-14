package config

import (
	"encoding/json"
	"os"

	"github.com/thrtn85/task-mgmt/models"
)

var Tasks []models.Task

func Init() {
	// read initial data from JSON file
	loadTasksFromJSON()
}

// Function to load tasks from JSON file
func loadTasksFromJSON() {
	// open and read data from JSON file
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		panic(err)
	}
	// Unmarshal JSON data into tasks slice
	err = json.Unmarshal(data, &Tasks)
	if err != nil {
		panic(err)
	}
}

// Function to save tasks to JSON file
func SaveTasksToJSON() error {
	// Serialize tasks slice to JSON
	data, err := json.Marshal(Tasks)
	if err != nil {
		return err
	}
	// Open tasks.json file in append mode
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write JSON data to tasks.json file
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}