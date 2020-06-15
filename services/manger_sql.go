package services

import (
	"database/sql"
	"fmt"
	"io/ioutil"
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
	// createTable(db)
	// InsertDataToTable(db)
	ReadDataFromTable(db)
	// db.Exec(`DROP TABLE Car`)
	// db.Exec(`DROP TABLE Custromer`)
	defer db.Close()

}

func createTable(db *sql.DB) error {
	createTable, err := ioutil.ReadFile("./models/sql/create_table.sql")
	if err != nil {
		log.Fatalln(err)
	}

	_, err = db.Exec(string(createTable))

	if err != nil {
		log.Fatalln("err from create table:", err)
	}
	return nil
}

func InsertDataToTable(db *sql.DB) error {

	_, err := db.Exec("INSERT INTO Custromer(Custfirstname,Custlastname,Custusername,Custpassword,statusPDPA) VALUES($1,$2,$3,$4)", "wongsathon", "sengcharoen", "wongsathon27", "Ws084038001", "Yes")
	if err != nil {
		log.Fatalln("error from insert function:", err)
	}
	return nil
}


func ReadDataFromTable(db *sql.DB) {
	row, err := db.Query("SELECT CustID,Custfirstname,Custlastname,Custusername,Custpassword FROM Custromer")
	if err != nil {
		log.Fatalln("err from read data from table:", err)
	}
	// UM := models.User{}
	colums, err := row.Columns()
	if err != nil {
		// return nil, err
	}
	result := make([]map[string]interface{}, 0)

	value := make([]sql.RawBytes, len(colums))

	pointerValue := make([]interface{}, len(value))
	for i := range value {
		pointerValue[i] = &value[i]
	}
	fmt.Print("pointerValue:", pointerValue)
	for row.Next() {
		item := make(map[string]interface{})
		err := row.Scan(pointerValue...)
		if err != nil {
			// return nil, err
		}
		var pasueValue string
		for i, col := range value {
			if col == nil {
				pasueValue = "NULL"
			} else {
				pasueValue = string(col)
			}
			item[colums[i]] = pasueValue
		}
		result = append(result, item)
		fmt.Println("result databse = ",result)
	}
	if row.Err(); err != nil {
		// return nil, err
	}
	// return result, nil
}
