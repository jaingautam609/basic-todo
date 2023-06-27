package dbHelper

import (
	"basic-todo/database"
	"basic-todo/models"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func CreateTask(db sqlx.Ext, title, description string, due_date time.Time, completed bool, token string) error {
	user_id, err := GetsessionId(database.Todo, token)
	if err != nil {
		return err
	}
	SQL := `INSERT INTO tasks(title, description, due_date,completed,userid) VALUES ($1,$2,$3,$4,$5)`
	_, err = db.Query(SQL, title, description, due_date, completed, user_id)
	log.Println("Added task.")
	if err != nil {
		return err
	}
	return nil
}
func GetsessionId(db sqlx.Ext, token string) (int, error) {
	SQL := `select user_id from sessions where token=$1`
	rows, err := db.Query(SQL, token)
	if err != nil {
		return -1, err
	}
	var user_id int
	for rows.Next() {

		err = rows.Scan(&user_id)
		if err != nil {
			return -2, err
		}

	}

	return user_id, nil
}
func AllTasks(db sqlx.Ext, token string) []models.Task {
	user_id, err := GetsessionId(database.Todo, token)
	SQL := `SELECT id, title, description, completed, due_date,created_at from todo where userid=$1`
	var tasks []models.Task
	rows, err := db.Query(SQL, user_id)
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

///*********************

func OrderedTasks(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT id, title, description, completed, due_date,created_at  FROM todo ORDER BY created_at ASC;`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func OrderedTasksDue(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT id, title, description, completed, due_date,created_at from todo order by due_date desc`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func OrderedTasksCompleted(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT id, title, description, completed, due_date,created_at from todo where completed = true`
	var tasks []models.Task
	rows, err := db.Query(SQL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.DueDate)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

/////////*************************
func GetTaskById(db sqlx.Ext, id string, token string) (models.Task, error) {
	user_id, err := GetsessionId(database.Todo, token)
	SQL := `SELECT id, title, description, completed, due_date,created_at from todo WHERE ID = $1 and userid=$2 `
	rows, err := db.Query(SQL, id, user_id)
	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)
	return task, nil
}
func DeleteTask(db sqlx.Ext, id string, token string) {
	user_id, err := GetsessionId(database.Todo, token)
	SQL := `DELETE FROM Tasks WHERE id = $1 and userid=$2`
	_, err = db.Query(SQL, id, user_id)
	if err != nil {
		return
	}
}
func UpdateTask(db sqlx.Ext, id string, token string) {
	user_id, err := GetsessionId(database.Todo, token)
	SQL := `UPDATE todo SET Completed = true WHERE id = $1 and userid=$2`
	_, err = db.Query(SQL, id, user_id)
	if err != nil {
		return
	}
}
