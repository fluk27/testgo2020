package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type ManagerSql struct {
}

func (ManagerSql) ConnectSql() {
	db, err := sql.Open("mysql", "root:keep1234@/costume")
	if err != nil {
		log.Fatalln("err from sql:", err)
	}
	fmt.Println("connected")
	defer db.Close()
}
