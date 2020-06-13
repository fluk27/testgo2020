package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// ManagerSQL is all fuction manager MySQL
type ManagerSQL struct {
}

//ConnectSQL is function connect to MySQL
func (ManagerSQL) ConnectSQL() {

	psqlInfo := "postgres://xniowcizeiwuce:e6bc64383dc9e17612c9c7378cd0fa1e45338a84eb59347d31c8403fde3c75ef@ec2-52-87-135-240.compute-1.amazonaws.com:5432/dc1km11drum90m"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	const qry = `CREATE TABLE label
	(
	   id int PRIMARY KEY NOT NULL,
	   name varchar(45)
	)`
	const deleteDB=`DROP TABLE label`
	_, err = db.Exec(deleteDB)

	if err != nil {
		log.Fatalln("err from crate table", err)
	}
	defer db.Close()

}
