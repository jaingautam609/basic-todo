package handler

import (
	"basic-todo/database"
	"basic-todo/database/dbHelper"
	"basic-todo/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func AllTasks() http.HandlerFunc {
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
func GetTaskById(w http.ResponseWriter, r *http.Request) {
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
func OrderedTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	taskall, err := dbHelper.OrderedTasks(database.Todo)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskall)
}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
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

// Delete task by ID function is here and calls another function from package dbHelper .
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	log.Println("initializeing...")
	if r.Method != http.MethodPost { // checking is the requested methed by client is put method or not
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	dbHelper.DeleteTask(database.Todo, id)      // dbHelper package function is called here to delete any task by query.
	taskall := dbHelper.AllTasks(database.Todo) // after delete , we are show all the remaining tasks from here
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(taskall)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func OrderedTasksDue(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	taskall, err := dbHelper.OrderedTasksDue(database.Todo)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskall)
}
func OrderedTasksCompleted(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	taskall, err := dbHelper.OrderedTasksCompleted(database.Todo)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskall)
}
