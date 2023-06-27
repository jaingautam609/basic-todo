package main

import (
	"basic-todo/database"
	"basic-todo/routes"
	"github.com/sirupsen/logrus"
	"net/http"
)

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

	routes.ServerRoutes()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logrus.Panicf("Failed to start server with error: %+v", err)
	}
}
