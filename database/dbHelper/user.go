package dbHelper

import (
	"basic-todo/database"
	"basic-todo/models"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func CreateTask(db sqlx.Ext, title, description string, dueDate time.Time, completed bool, token string) error {
	userId, err := GetSessionId(database.Todo, token)
	if err != nil {
		return err
	}
	SQL := `INSERT INTO tasks(title, description, dueDate,completed,userid) VALUES ($1,$2,$3,$4,$5)`
	_, err = db.Query(SQL, title, description, dueDate, completed, userId)
	log.Println("Added task.")
	if err != nil {
		return err
	}
	return nil
}
func GetSessionId(db sqlx.Ext, token string) (int, error) {
	SQL := `select user_id from sessions where token=$1`
	rows, err := db.Query(SQL, token)
	if err != nil {
		return 0, err
	}
	var userId int
	for rows.Next() {

		err = rows.Scan(&userId)
		if err != nil {
			//log.Println("Error found here")
			return 0, err
		}

	}

	return userId, nil
}
func AllTasks(db sqlx.Ext, token string) ([]models.Task, error) {
	var tasks []models.Task
	userId, err := GetSessionId(database.Todo, token)
	if err != nil {
		return tasks, err
	}
	SQL := `SELECT id, title, description, completed, dueDate,createdAt from tasks where userid=$1`
	rows, err := db.Query(SQL, userId)
	//TODO send http error status
	if err != nil {
		//fmt.Fprintf("hello")
		log.Println(err)
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
	//TODO send http status
	return tasks, nil
}

///*********************

func OrderedTasks(db sqlx.Ext) ([]models.Task, error) {
	SQL := `SELECT id, title, description, completed, dueDate,createdAt  FROM todo ORDER BY created_at ASC;`
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
	SQL := `SELECT id, title, description, completed, dueDate,createdAt from todo order by due_date desc`
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
	SQL := `SELECT id, title, description, completed, dueDate,createdAt from todo where completed = true`
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
func GetTaskById(db sqlx.Ext, id int, token string) (models.Task, error) {
	userId, err := GetSessionId(database.Todo, token)
	SQL := `SELECT id, title, description, completed, dueDate,createdAt from tasks WHERE ID = $1 and userid=$2 `
	rows, err := db.Query(SQL, id, userId)
	if err != nil {
		return models.Task{}, err
	}
	var task models.Task
	//var tasks []models.Task
	for rows.Next() {

		err = rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.DueDate, &task.CreatedAt)
		if err != nil {
		}
	}
	return task, nil
}

func DeleteTask(db sqlx.Ext, id int, token string) error {
	userId, err := GetSessionId(database.Todo, token)
	SQL := `DELETE FROM Tasks WHERE id = $1 and userid=$2`
	_, err = db.Query(SQL, id, userId)
	if err != nil {
		return err
	}
	return nil
}
func UpdateTask(db sqlx.Ext, id string, token string) error {
	userId, err := GetSessionId(database.Todo, token)
	SQL := `UPDATE todo SET Completed = true WHERE id = $1 and userid=$2`
	_, err = db.Query(SQL, id, userId)
	if err != nil {
		return err
	}
	return nil
}
