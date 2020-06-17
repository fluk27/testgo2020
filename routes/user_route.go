package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/peewlaom/testgo/controllers"
)

//UserRoute is all end about user
func UserRoute(route *echo.Echo) {
	uc := &controllers.UserController{}
	userGroup := route.Group("/user")
	// userGroup.Use(middleware.Logger())
	userGroup.POST("/register", uc.Register)
	// userGroup.POST("/login")
	// userGroup.PUT("/editDataPerson")
}
