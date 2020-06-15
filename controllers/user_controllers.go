package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/peewlaom/testgo/models"
	"github.com/peewlaom/testgo/services"
)

// User is all fuction manager user
type User struct {
}

var header map[string]string
var url string = "https://notify-api.line.me/api/notify"
var method string = "POST"

func (u *User) Login(c echo.Context) error {
	SQLS := &services.ManagerSQL{}
	SQLS.ConnectSQL()
	UM := &models.User{}
	c.Bind(UM)
	us := &services.UserServices{}
	ciherText, err := us.Register(UM.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.String(http.StatusOK, ciherText)
}

// sendMessageToLineNoutify
func (u User) sendMessageToLineNotify(message string) {
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
