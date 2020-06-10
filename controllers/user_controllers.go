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

func (u *User) Login(c echo.Context) error {
	username := c.FormValue("username")
	cipherText := c.FormValue("password")
	RSAService := &services.RSAKey{}
	// RSAService.PathPrivateKey= "./"
	RSAService.FileNamePrivateKey = "privateKey.pem"
	RSAService.GenerateRSAKey(256)
	// password, err := RSAService.DncyptDataWithPKC(cipherText)
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, err.Error())
	// }
	if username == "peewlaom" && cipherText== "Ws0844038001" {
		log.Println("logined by user")
		// u.sendMessageToLineNotify("I cannot forget you remember me")
		return c.String(http.StatusOK, "logined")
	}
	return c.JSON(http.StatusOK, map[string]string{"masseage": "login failed"})
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
