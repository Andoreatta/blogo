package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewUserRepo(db)
	userFromDB, err := repo.SearchByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckHash(userFromDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.CreateToken(userFromDB.ID)
	fmt.Println(token)
}
