package models

// User is ORM user
type User struct {
	ID    int    `่json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
