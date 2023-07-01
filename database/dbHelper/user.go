package dbHelper

import (
	"basic-todo/database"
	"basic-todo/models"
	"github.com/jmoiron/sqlx"
	"time"
)

func CreateTask(db *sqlx.DB, title, description string, dueDate time.Time, completed bool, userId int) error {
	SQL := `INSERT INTO tasks(title, description, due_date,completed,user_id) VALUES ($1,$2,$3,$4,$5)`
	_, err := db.Exec(SQL, title, description, dueDate, completed, userId)
	if err != nil {
		return err
	}
	return nil
}
func GetSessionId(db *sqlx.DB, token string) (int, error) {
	var userId int
	SQL := `select user_id from sessions where token=$1`
	err := db.QueryRowx(SQL, token).Scan(&userId)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func AllTasks(db sqlx.Ext, userId int) ([]models.Task, error) {
	var tasks []models.Task
	SQL := `SELECT id, title, description, completed, due_date, created_at FROM tasks WHERE user_id = $1 ORDER BY due_date DESC;`
	rows, err := db.Query(SQL, userId)
	if err != nil {
		return tasks, err
	}
	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.DueDate, &task.CreatedAt)
		if err != nil {

			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTaskById(db *sqlx.DB, id int, userId int) (models.Task, error) {
	var task models.Task
	SQL := `SELECT id, title, description, completed, due_date,created_at from tasks WHERE ID = $1 and user_id=$2 `
	err := db.QueryRowx(SQL, id, userId).Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.DueDate, &task.CreatedAt)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func DeleteTask(db *sqlx.DB, id int, token string) error {

	userId, err := GetSessionId(database.Todo, token)
	if err != nil {
		return err
	}
	SQL := `DELETE FROM tasks WHERE id = $1 and user_id=$2`
	_, err = db.Exec(SQL, id, userId)
	if err != nil {
		return err
	}
	return nil
}
func UpdateTask(db sqlx.Ext, id string, token string) error {

	userId, err := GetSessionId(database.Todo, token)
	if err != nil {
		return err
	}
	SQL := `UPDATE tasks SET Completed = true WHERE id = $1 and user_id=$2`
	_, err = db.Exec(SQL, id, userId)
	if err != nil {
		return err
	}
	return nil
}
