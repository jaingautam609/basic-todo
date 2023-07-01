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

	token := r.Header.Get("Authentication")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var taskAll []models.Task
	var userId int
	userId, err := dbHelper.GetSessionId(database.Todo, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskAll, err = dbHelper.AllTasks(database.Todo, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(taskAll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authentication")
	urlParts := strings.Split(r.URL.Path, "/")
	var id = urlParts[len(urlParts)-1]
	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	var userId int
	userId, err = dbHelper.GetSessionId(database.Todo, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	myTask, err := dbHelper.GetTaskById(database.Todo, num, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(myTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//w.WriteHeader(http.StatusOK)
}
func AddTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authentication")
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		log.Println(err)
		return
	}
	var userId int
	userId, err = dbHelper.GetSessionId(database.Todo, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	err = dbHelper.CreateTask(database.Todo, newTask.Title, newTask.Description, newTask.DueDate, newTask.Completed, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		//log.Println(err)
		return
	}

	taskAll, err := dbHelper.AllTasks(database.Todo, userId)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error to create account", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(taskAll)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authentication")

	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	err := dbHelper.UpdateTask(database.Todo, id, token)

	if err != nil {
		http.Error(w, "Error to update account", http.StatusBadRequest)
		return
	}
	var userId int
	userId, err = dbHelper.GetSessionId(database.Todo, token)
	if err != nil {
		http.Error(w, "Error to get session id", http.StatusBadRequest)
		return
	}
	var tasksAll []models.Task
	tasksAll, err = dbHelper.AllTasks(database.Todo, userId)
	if err != nil {
		//log.Println(err)
		http.Error(w, "Error to create account", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasksAll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Authentication")
	urlParts := strings.Split(r.URL.Path, "/")
	id := urlParts[len(urlParts)-1]
	num, err := strconv.Atoi(id)
	if err != nil {
		//log.Println(err)
		fmt.Println("Error:", err)
		return
	}
	err = dbHelper.DeleteTask(database.Todo, num, token)
	if err != nil {
		//log.Println(err)
		http.Error(w, "Error to delete task", http.StatusBadRequest)
		return
	}
	var userId int
	userId, err = dbHelper.GetSessionId(database.Todo, token)
	if err != nil {
		http.Error(w, "Error to get session id", http.StatusBadRequest)
		return
	}
	taskAll, err := dbHelper.AllTasks(database.Todo, userId) // after delete , we are show all the remaining tasks from here
	if err != nil {
		http.Error(w, "Error to create account", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(taskAll)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)
}
