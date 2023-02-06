package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

// Creates a new repository of users using the users struct as base
func NewUserRepo(db *sql.DB) *users {
	return &users{db}
}

// Creates a user insert into database
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

// Returns all the users that contains the string content
func (repo users) Search(username string) ([]models.User, error) {
	username = fmt.Sprintf("%%%s%%", username) //%username%
	lines, err := repo.db.Query(
		"SELECT userId, username, email, createdAt FROM users WHERE username LIKE ?", username,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User
		if err = lines.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

// Fetches by Id from the database
func (repo users) SearchById(Id uint64) (models.User, error) {
	lines, err := repo.db.Query(
		"SELECT userId, username, email, createdAt FROM users WHERE userId = ?", Id,
	)
	if err != nil {
		// has to return empty user in case of error
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
