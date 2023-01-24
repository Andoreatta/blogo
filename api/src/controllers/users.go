package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating User."))
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewUserRepo(db)
	repo.Create(user)
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
