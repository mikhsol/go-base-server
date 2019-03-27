package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var InvalidUsername = errors.New("invalid username")
var InvalidPassword = errors.New("invalid password")
var WrongPassword = errors.New("wrong password")

type User struct {
	Username       string
	hashedPassword []byte
}

func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword(u.hashedPassword, []byte(password))
	if err != nil {
		// TODO: Think about proper logging of bcrypt errors
		return WrongPassword
	}
	return nil
}

func validateUsername(username string) (string, error) {
	if username == "" {
		return "", InvalidUsername
	}
	return username, nil
}

func validatePassword(password string) (string, error) {
	if password == "" {
		return "", InvalidPassword
	}
	return password, nil
}

// CreateUser return new instance of user based with associated username and
// password hash based on bcrypt algorithm. Return error if username or password
// was invalid. OS specific errors can be returned by bcrypt
// GenerateFromPassword function.
func CreateUser(username, password string) (*User, error) {
	username, err := validateUsername(username)
	if err != nil {
		return nil, err
	}

	password, err = validatePassword(password)
	if err != nil {
		return nil, err
	}

	hashedPwd, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := User{username, hashedPwd}
	return &user, nil
}
