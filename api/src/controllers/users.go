package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}
	// 	w.Write([]byte("Creating User."))

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewUserRepo(db)
	userId, err := repo.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("Id inserted: %d", userId)))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User."))
}
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching User."))
}
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Fetching Users."))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User."))
}
