package main

import (
	"basic-todo/database"
	"basic-todo/database/dbHelper"
	"basic-todo/models"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
)

func allTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var taskall []models.Task = dbHelper.AllTasks(database.Todo)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(taskall)
	}
}
func AddTask(w http.ResponseWriter, r *http.Request) {
	log.Println("initializeing...")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		return
	}
	errorCreated := dbHelper.CreateTask(database.Todo, newTask.ID, newTask.Title, newTask.Description, newTask.Completed)
	if errorCreated != nil {
		return
	}
	taskall := dbHelper.AllTasks(database.Todo)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(taskall)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func getTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	urlParts := strings.Split(r.URL.Path, "/")
	var id = urlParts[len(urlParts)-1]
	mytask, err := dbHelper.GetTaskById(database.Todo, id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mytask)
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	dbHelper.UpdateTask(database.Todo, id)
	var tasksall []models.Task = dbHelper.AllTasks(database.Todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasksall)
}
func main() {
	if err := database.ConnectAndMigrate(
		"localhost",
		"5433",
		"todo",
		"local",
		"local",
		database.SSLModeDisable); err != nil {
		logrus.Panicf("Failed to initialize and migrate database with error: %+v", err)
	}
	http.Handle("/tasks", allTasks())
	http.HandleFunc("/tasks/", getTaskById)
	http.HandleFunc("/add-task", AddTask)
	http.HandleFunc("/update/", updateTask)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
