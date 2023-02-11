package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Input sanitization
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}
	if err := user.format(step); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(step string) error {
	switch {
	case user.Username == "":
		{
			return errors.New("the username field is empty")
		}
	case user.Email == "":
		{
			return errors.New("the e-mail field is empty")
		}
	case checkmail.ValidateFormat(user.Email) != nil:
		{
			return errors.New("the e-mail has an invalid format")
		}
	case step == "registration" && user.Password == "":
		{
			return errors.New("the password field is empty")
		}
	}
	return nil
}

func (user *User) format(step string) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Username = strings.TrimSpace(user.Email)

	switch step {
	case "registration":
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return nil
}
