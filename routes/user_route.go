package routes

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// all end about user
func UserRoute(route *echo.Echo) {
	userGroup := route.Group("/user")
	userGroup.POST("/login", login)
	// userGroup.POST("/register")
	// userGroup.PUT("/editDataPerson")
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "peewlaom" && password == "Ws0844038001" {
		log.Println("logined by user")
		res, err := http.Post("https://notify-api.line.me/api/notify",
			"application/x-www-form-urlencoded",
			bytes.NewBuffer([]byte(`your query`)))
		if err != nil {

		}
		client := &http.Client{}
		resp, err := client.Do(res)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body:", string(body))
		res.Header.Set("Authorization", "Bearer zZhHLegDUVgZg06h9oSM33h7uZj7deicsCVRvbxOnWl")
		return c.String(http.StatusOK, "logined")
	} else {
		return c.JSON(http.StatusOK, map[string]string{"masseage": "login failed"})
	}
}
