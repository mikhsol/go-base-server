package main

import "testing"

func TestCreateUser(t *testing.T) {
	username := "Rick"
	user := CreateUser(username)

	if user.username != username {
		t.Errorf("want %s got %s", username, user.username)
	}
}
