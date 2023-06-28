package models

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	DueDate     time.Time `json:"due_date"`
}
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UID      int    `json:"uid"`
}
