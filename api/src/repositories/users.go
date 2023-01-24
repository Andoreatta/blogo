package repositories

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// Creates a new repository of users using the users struct as base
func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

// Creates insert into database
func (repo users) Create(user models.User) (uint64, error) {
	return 0, nil
}
