package dbHelper

import (
	"basic-todo/models"
	"github.com/jmoiron/sqlx"
)

func AllTasks(db sqlx.Ext) []models.Task {
	SQL := `SELECT ID, Title, Description, Completed from Tasks`
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
