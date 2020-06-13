package models

// User is ORM user
type User struct {
	ID       int    `json:"id"`
	FirstName     string `json:"first_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
