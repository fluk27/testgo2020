package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ManagerSQL is all fuction manager MySQL
type ManagerSQL struct {
}

//ConnectSQL is function connect to MySQL
func (ManagerSQL) ConnectSQL() {
	db, err := sql.Open("mysql", "root:keep1234@/costume")
	if err != nil {
		log.Fatalln("err from sql:", err)
	}
	fmt.Println("connected")
	defer db.Close()
}
