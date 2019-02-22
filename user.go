package main

type User struct {
	username string
}

func CreateUser(username string) User {
	return User{username}
}
