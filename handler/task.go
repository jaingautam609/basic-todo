package handler

import (
	"basic-todo/database"
	"basic-todo/database/dbHelper"
	"basic-todo/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func AllTasks(w http.ResponseWriter, r *http.Request) {

	//token := r.Header.Get("Session-Id")
	token := r.Header.Get("Auth_id")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var taskall []models.Task
	taskall, err := dbHelper.AllTasks(database.Todo, token)
	if err != nil {
		http.Error(w, "Error to show all task", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskall)

}
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Auth_id")
	urlParts := strings.Split(r.URL.Path, "/")
	var id = urlParts[len(urlParts)-1]
	Num, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	mytask, err := dbHelper.GetTaskById(database.Todo, Num, token)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mytask)
}
func AddTask(w http.ResponseWriter, r *http.Request) {
	//log.Println("initializeing...")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Auth_id")
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		log.Println(err)
		return
	}
	errorCreated := dbHelper.CreateTask(database.Todo, newTask.Title, newTask.Description, newTask.DueDate, newTask.Completed, token)
	if errorCreated != nil {
		log.Println(errorCreated)
		return
	}
	taskall, err := dbHelper.AllTasks(database.Todo, token)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error to create account", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(taskall)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "successfull")
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
	token := r.Header.Get("Auth_id")
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	err := dbHelper.UpdateTask(database.Todo, id, token)
	if err != nil {
		http.Error(w, "Error to update account", http.StatusBadRequest)
		return
	}
	var tasksall []models.Task
	tasksall, err = dbHelper.AllTasks(database.Todo, token)
	if err != nil {
		//log.Println(err)
		http.Error(w, "Error to create account", http.StatusBadRequest)
	}
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
	token := r.Header.Get("Auth_id")
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	Num, err := strconv.Atoi(id)
	if err != nil {
		//log.Println(err)
		fmt.Println("Error:", err)
		return
	}
	err = dbHelper.DeleteTask(database.Todo, Num, token)
	if err != nil {
		//log.Println(err)
		http.Error(w, "Error to delete account", http.StatusBadRequest)
		return
	}
	// dbHelper package function is called here to delete any task by query.
	taskall, err := dbHelper.AllTasks(database.Todo, token) // after delete , we are show all the remaining tasks from here
	if err != nil {
		http.Error(w, "Error to create account", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(taskall)
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
