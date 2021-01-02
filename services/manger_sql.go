package services

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

// ManagerSQL is all fuction manager MySQL
type ManagerSQL struct {
}
type ManagerMSSQL struct {
}

//ConnectSQL is function connect to MySQL
func connectSQL() (*sql.DB, error) {

	psqlInfo := "postgres://xniowcizeiwuce:e6bc64383dc9e17612c9c7378cd0fa1e45338a84eb59347d31c8403fde3c75ef@ec2-52-87-135-240.compute-1.amazonaws.com:5432/dc1km11drum90m"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return db, nil

}

//CreateTable is function crate table in database. Before create table, it read sql from sql file
func (ManagerSQL) CreateTable(filename string, path string) error {
	db, err := connectSQL()
	if err != nil {
		return errors.New("errr form connectSQL:" + err.Error())
	}
	createTable, err := ioutil.ReadFile(filename + path)
	if err != nil {
		log.Fatalln(err)
		return errors.New("err read sql file:" + err.Error())
	}

	_, err = db.Exec(string(createTable))

	if err != nil {
		// log.Fatalln("err from create table:", err)
		return errors.New("err from create table:" + err.Error())
	}
	return nil
}

func (ManagerSQL) DropTable() error {
	db, err := connectSQL()
	if err != nil {
		return errors.New("errr form connectSQL:" + err.Error())
	}
	_, err = db.Exec(`DROP TABLE Custromer;`)
	if err != nil {
		return err
	}
	return nil
}

func (ManagerSQL) InsertDataToTable() error {
	db, err := connectSQL()
	if err != nil {
		return errors.New("errr form connectSQL:" + err.Error())
	}
	_, err = db.Exec("INSERT INTO Custromer(Custfirstname,Custlastname,Custusername,Custpassword,statusPDPA) VALUES($1,$2,$3,$4,$5)", "wongsathon", "sengcharoen", "wongsathon2539", "Ws084038001", "true")
	if err != nil {
		log.Fatalln("error from insert function:", err)
	}
	return nil
}

func (ManagerSQL) InsertDataToTableCar() error {
	db, err := connectSQL()
	if err != nil {
		return errors.New("errr form connectSQL:" + err.Error())
	}
	_, err = db.Exec("INSERT INTO Car (carName) VALUES($1)", "toyota dxi")
	if err != nil {
		log.Fatalln("error from insert function:", err)
	}
	return nil
}

func (ManagerSQL) ReadDataFromTable() ([]map[string]interface{}, error) {
	db, err := connectSQL()
	if err != nil {
		return nil, errors.New("errr form connectSQL:" + err.Error())
	}
	row, err := db.Query("SELECT CustID,Custfirstname,Custlastname,Custusername,Custpassword FROM Custromer")
	if err != nil {
		log.Fatalln("err from read data from table:", err)
	}
	// UM := models.User{}
	colums, err := row.Columns()
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)

	value := make([]sql.RawBytes, len(colums))
	pointerValue := make([]interface{}, len(value))
	for i := range value {
		pointerValue[i] = &value[i]
	}
	for row.Next() {
		item := make(map[string]interface{})
		err := row.Scan(pointerValue...)
		if err != nil {
			return nil, err
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
		fmt.Println("result databse = ", result)
	}
	if row.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (ManagerSQL) ReadDataFromTableCar() ([]map[string]interface{}, error) {
	db, err := connectSQL()
	if err != nil {
		return nil, errors.New("errr form connectSQL:" + err.Error())
	}
	row, err := db.Query("SELECT * FROM Car INNER JOIN Custromer.CustID = Car.CustID ")
	if err != nil {
		log.Fatalln("err from read data from table:", err)
	}
	// UM := models.User{}
	colums, err := row.Columns()
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)

	value := make([]sql.RawBytes, len(colums))
	pointerValue := make([]interface{}, len(value))
	for i := range value {
		pointerValue[i] = &value[i]
	}
	for row.Next() {
		item := make(map[string]interface{})
		err := row.Scan(pointerValue...)
		if err != nil {
			return nil, err
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
		fmt.Println("result databse = ", result)
	}
	if row.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (ManagerSQL) DeleteTable() {
	db, err := connectSQL()
	if err != nil {
		// return nil, errors.New("errr form connectSQL:" + err.Error())
	}
	_, err = db.Query("DELETE FROM Custromer")
	if err != nil {
		log.Fatalln("err Prepare :", err)
	}
}

func (ManagerMSSQL) ConnectMSSQL() (*sql.DB,error) {

	// var server= "172.16.61.22"
	// var database= "CCMS_PCS_TEST_2"
	// var user="appuser"
	// var password= "spt1234"
	var server= "172.16.15.19" //JPAC
	var database= "CCMS_PCS_TEST"
	var user= "sa"
	var password= "Bj4free"
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, password, database)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	log.Printf("Connected!\n")
	// QueryCount:="select count (*) FROM vwReceiveList"
	return db,nil
	
}

func (m *ManagerMSSQL)ReadReceive()(interface{}, error)  {

db, err := m.ConnectMSSQL()
if err != nil {
	return nil, errors.New("errr form connectSQL:" + err.Error())
}

	quertStr := "SELECT \n" +
		"FORMAT (vw.CreatedDate, 'dd/MM/yyyy HH:mm') AS CreatedDate1,\n" +
		"FORMAT (vw.InformDate, 'dd/MM/yyyy HH:mm') AS InformDate1,\n" +
		"FORMAT (vw.EstActionDate, 'dd/MM/yyyy HH:mm') AS EstActionDate1,\n" +
		"FORMAT (vw.TempReceiveDate, 'dd/MM/yyyy HH:mm') AS TempReceiveDate1,\n" +
		"FORMAT (vw.UpdatedDate, 'dd/MM/yyyy HH:mm') AS UpdatedDate1,\n" +
		"vw.*,\n" +
		"rw.Worker AS WorkerName\n" +
		"FROM vwReceiveList vw\n" +
		"LEFT JOIN vwReceiveWorker rw ON rw.OrderId = vw.OrderId\n" +
		"WHERE vw.State IN ('2', '3', '4', '5')"
	row, err := db.Query(quertStr)
	if err != nil {
		log.Fatalln("err from read data from table:", err)
	}
	// UM := models.User{}
	
	colums, err := row.Columns()
	if err != nil {
		return nil, err
	}
	result := make([]map[string]interface{}, 0)

	value := make([]sql.RawBytes, len(colums))
	pointerValue := make([]interface{}, len(value))
	for i := range value {
		pointerValue[i] = &value[i]
	}
	for row.Next() {
		item := make(map[string]interface{})
		err := row.Scan(pointerValue...)
		if err != nil {
			return nil, err
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
		// fmt.Println("result databse = ", result)
	}
	if row.Err(); err != nil {
		return nil, err
	}
	defer db.Close()
	return result, nil
}
