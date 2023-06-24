package dbHelper

import (
	"basic-todo/models"
	"github.com/jmoiron/sqlx"
	"log"
)

func CreateTask(db sqlx.Ext, id, title, description string, completed bool) error {
	SQL := `INSERT INTO todo(ID, Title, Description, Completed) VALUES ($1,$2,$3,$4)`
	_, err := db.Query(SQL, id, title, description, completed)
	log.Println("Added task.")
	if err != nil {
		return err
	}
	return nil
}
func AllTasks(db sqlx.Ext) []models.Task {
	SQL := `SELECT id, title, description, completed from todo`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return tasks
	}
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
		if err != nil {
			return nil
		}
		tasks = append(tasks, task)
	}
	return tasks
}
func GetTaskById(db sqlx.Ext, id string) (models.Task, error) {
	SQL := `SELECT ID, Title, Description, Completed from todo WHERE ID = $1`
	rows, err := db.Query(SQL, id)
	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	return task, nil
}
func UpdateTask(db sqlx.Ext, id string) {
	SQL := `UPDATE todo SET Completed = true WHERE ID = $1`
	_, err := db.Query(SQL, id)
	if err != nil {
		return
	}
}
