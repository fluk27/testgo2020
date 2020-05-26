package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//UserRoute is all end about user
func UserRoute(route *echo.Echo) {
	userGroup := route.Group("/user")
	userGroup.Use(middleware.Logger())
	userGroup.POST("/login", login)
	// userGroup.POST("/register")
	// userGroup.PUT("/editDataPerson")
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username == "peewlaom" && password == "Ws0844038001" {
		log.Println("logined by user")
		url := "https://notify-api.line.me/api/notify"
		method := "POST"

		payload := strings.NewReader("message=I'm miss P' Bew")

		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
		}
		// req, err := http.PostForm(url, url.Values{"key": {"Value"}, "id": {"123"}})
		req.Header.Add("Authorization", "Bearer zZhHLegDUVgZg06h9oSM33h7uZj7deicsCVRvbxOnWl")
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			// return nil, err
		}
		defer res.Body.Close()
		// body, err := ioutil.ReadAll(res.Body)

		// fmt.Println(string(body))
		return c.String(http.StatusOK, "logined")
	}
	return c.JSON(http.StatusOK, map[string]string{"masseage": "login failed"})
}
