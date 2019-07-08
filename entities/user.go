package entities

// User entity
type User struct {
	ID       int    `json:"id" fury:"primary_key"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
