package models

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUser(t *testing.T) {
	username := "Rick"
	password := "ChangeMe"

	t.Run("New user success", func(t *testing.T) {
		user, err := CreateUser(username, password)

		if err != nil {
			t.Fatal(err)
		}
		if user.Username != username {
			t.Fatalf("want %s got %s", username, user.Username)
		}
		err = bcrypt.CompareHashAndPassword(user.hashedPassword, []byte(password))
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("New user, empty username error", func(t *testing.T) {
		want := InvalidUsername
		_, err := CreateUser("", password)
		if err != want {
			t.Errorf("want %s got %s", want, err)
		}
	})

	t.Run("New user, empty password error", func(t *testing.T) {
		want := InvalidPassword
		_, err := CreateUser(username, "")
		if err != want {
			t.Errorf("want %s got %s", want, err)
		}
	})

	t.Run("User, mismatch hash and password", func(t *testing.T) {
		user, _ := CreateUser(username, password)
		testCases := []struct {
			password string
			want     error
		}{
			{password, nil},
			{"wrong password", WrongPassword},
			{"", WrongPassword},
		}
		for _, tc := range testCases {
			err := user.CheckPassword(tc.password)
			if err != tc.want {
				t.Fatalf("want %s got %s", tc.want, err)
			}
		}

	})
}
