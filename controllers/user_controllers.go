package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/peewlaom/testgo/services"
)

// User is all fuction manager user
type User struct {
}

var header map[string]string
var url string = "https://notify-api.line.me/api/notify"
var method string = "POST"

func (u User) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "peewlaom" && password == "Ws0844038001" {
		log.Println("logined by user")
		// u.sendMessageToLineNoutify("I cannot forget you remember me")
		service := services.RSAKey{}
		service.EncyptDataWithPKC()
		// service.ReadPemFile()

		return c.String(http.StatusOK, "logined")
	}
	return c.JSON(http.StatusOK, map[string]string{"masseage": "login failed"})
}

// sendMessageToLineNoutify
func (u User) sendMessageToLineNoutify(message string) {
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
	// body, err := ioutil.ReadAll(res.Body)

	// fmt.Println(string(body))
}
