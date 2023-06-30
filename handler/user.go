package handler

import (
	"basic-todo/database"
	"basic-todo/database/authentication"
	"basic-todo/models"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		//http.Error(w, "Error to create account", http.StatusBadRequest)
		//return
	}
	uId, err := authentication.Login(database.Todo, User.Username, User.Password)
	if err != nil {
		//w.WriteHeader(http.StatusUnauthorizedd)
		http.Error(w, "Unauthorized person", http.StatusUnauthorized)
		log.Println(w, http.StatusUnauthorized)
		return
	}
	token := uuid.New().String()
	err = authentication.CreateSession(database.Todo, token, uId)
	if err != nil {
		log.Println("sessions not created")
		return
	}
	w.Header().Set("Auth_id", token)
	fmt.Fprintf(w, "successfull")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		//http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	token := r.Header.Get("Auth_id")
	err := authentication.Logout(database.Todo, token)
	if err != nil {
		//http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprintf(w, "successfull")
}
func NewUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		//w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var User models.User
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	err = authentication.Create(database.Todo, User.Username, User.Password)
	if err != nil {
		http.Error(w, "Error to create account", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "successfull")
}
