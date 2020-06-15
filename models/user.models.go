package models

// User is ORM user
type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"user_name"`
	Password   string `json:"password"`
	StatusPDPA string `json:"status_PDPA"`
}
