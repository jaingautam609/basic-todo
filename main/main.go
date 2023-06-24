package main

import (
	"basic-todo/database"
	"basic-todo/database/dbHelper"
	"basic-todo/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func allTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var taskall []models.Task = dbHelper.AllTasks(database.Todo)
		json.NewEncoder(w).Encode(taskall)
	}
}
func main() {
	if err := database.ConnectAndMigrate(
		"localhost",
		"5434",
		"todo",
		"local",
		"local",
		database.SSLModeDisable); err != nil {
		logrus.Panicf("Failed to initialize and migrate database with error: %+v", err)
	}
	http.Handle("/tasks", allTasks())
	log.Fatal(http.ListenAndServe(":3001", nil))
}
