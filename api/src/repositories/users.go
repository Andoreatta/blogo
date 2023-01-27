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
	statement, err := repo.db.Prepare("INSERT INTO users (username, email, passwd) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastIdInserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIdInserted), nil
}
