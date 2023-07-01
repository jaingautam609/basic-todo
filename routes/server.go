package routes

import (
	"basic-todo/handler"
	"net/http"
)

func ServerRoutes() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/create", handler.NewUser)
	http.HandleFunc("/tasks", handler.AllTasks)
	http.HandleFunc("/task/", handler.GetTaskById)
	http.HandleFunc("/task/add", handler.AddTask)
	http.HandleFunc("/task/update/", handler.UpdateTask)
	http.HandleFunc("/task/delete/", handler.DeleteTask)
}
