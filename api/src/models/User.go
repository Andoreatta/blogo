package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Input sanitization
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

func (user *User) validate() error {
	switch {
	case user.Username == "":
		{
			return errors.New("the username field is empty")
		}
	case user.Email == "":
		{
			return errors.New("the e-mail field is empty")
		}
	case user.Password == "":
		{
			return errors.New("the password field is empty")
		}
	}
	return nil
}

func (user *User) format() {
	user.Username = strings.TrimSpace(user.Username)
	user.Username = strings.TrimSpace(user.Email)
}
