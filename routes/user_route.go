package routes

import (
	"github.com/fluk27/testgo/controllers"
	"github.com/labstack/echo/v4"
)

//UserRoute is all end about user
func UserRoute(route *echo.Echo) {
	uc := &controllers.UserController{}
	userGroup := route.Group("/user")
	// userGroup.Use(middleware.Logger())
	userGroup.GET("/getInfoElasticsearch", uc.GetInfoElasticsreach)
	userGroup.POST("/register", uc.Register)
	// userGroup.POST("/login")
	route.GET("/testMSSQL", uc.TestMSSQL)
	// userGroup.PUT("/editDataPerson")
}
