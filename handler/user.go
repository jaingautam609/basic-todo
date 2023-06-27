package handler

import (
	"basic-todo/authen"
	"basic-todo/database"
	"basic-todo/models"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		return
	}
	UID, err := authen.Login(database.Todo, User.Username, User.Password)
	if err != nil && UID <= 0 {
		log.Println(w, http.StatusUnauthorized)
		return
	}
	token := uuid.New().String()
	w.Header().Set("Auth_id", token)
	authen.CreateSession(database.Todo, token, UID)
	fmt.Fprintf(w, "successfull")
}
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Auth_id")
	authen.Logout(database.Todo, token)
	fmt.Fprintf(w, "successfull")
}
func NewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		return
	}
	err = authen.Create(database.Todo, User.Username, User.Password)
	if err != nil {
		return
	}

	fmt.Fprintf(w, "successfull")
}
