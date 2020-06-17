package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/peewlaom/testgo/models"
	"github.com/peewlaom/testgo/services"
)

// User is all fuction manager user
type UserController struct {
}

var header map[string]string
var url string = "https://notify-api.line.me/api/notify"
var method string = "POST"

func (u *UserController) Register(c echo.Context) error {
	SQLS := &services.ManagerSQL{}
	// err:=SQLS.CreateTable("./models/sql/","create_table.sql")
	var UM []models.User
	// err:=SQLS.DropTable()
	err := SQLS.InsertDataToTableCar()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	result, err := SQLS.ReadDataFromTableCar()
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	mapstructure.Decode(result, &UM)
	//data, _ := json.Marshal(result)
	return c.JSON(http.StatusCreated, UM)
}

// sendMessageToLineNoutify
func (u UserController) sendMessageToLineNotify(message string) {
	payload := strings.NewReader("message=" + message)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer zZhHLegDUVgZg06h9oSM33h7uZj7deicsCVRvbxOnWl")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		// return nil, err
	}
	defer res.Body.Close()

}
