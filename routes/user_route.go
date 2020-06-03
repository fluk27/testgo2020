package routes

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/peewlaom/testgo/controllers"
)

//UserRoute is all end about user
func UserRoute(route *echo.Echo) {
	uc := &controllers.User{}
	userGroup := route.Group("/user")
	// userGroup.Use(middleware.Logger())
	userGroup.POST("/login", uc.Login)
	// userGroup.POST("/register")
	// userGroup.PUT("/editDataPerson")
}
