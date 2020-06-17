package models

// User is ORM user
type User struct {
	ID         string `json:"id" mapstructure:"CustID"`
	FirstName  string `json:"first_name" mapstructure:"Custfirstname"`
	LastName   string `json:"last_name"  mapstructure:"Custlastname"`
	Username   string `json:"user_name" mapstructure:"Custusername"`
	Password   string `json:"password" mapstructure:"Custpassword"`
	StatusPDPA string `json:"status_PDPA" mapstructure:"statusPDPA"`
}

type Car struct {
	
}
