package routes

import (
	"basic-todo/handler"
	"net/http"
)

func ServerRoutes() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/new", handler.NewUser)
	http.Handle("/tasks", handler.AllTasks())
	http.HandleFunc("/tasks/", handler.GetTaskById)
	http.HandleFunc("/add-task", handler.AddTask)
	http.HandleFunc("/update/", handler.UpdateTask)
	http.HandleFunc("/delete/", handler.DeleteTask)
	http.HandleFunc("/tasks/ordered", handler.OrderedTasks)
	http.HandleFunc("/tasks/ordered/due", handler.OrderedTasksDue)
	http.HandleFunc("/tasks/ordered/completed", handler.OrderedTasksCompleted)

}
