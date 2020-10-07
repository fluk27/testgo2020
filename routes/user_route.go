package routes

import (
	"log"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/peewlaom/testgo/controllers"
)

//UserRoute is all end about user
func UserRoute(route *echo.Echo) {
	uc := &controllers.UserController{}
	userGroup := route.Group("/user")
	// userGroup.Use(middleware.Logger())
	userGroup.GET("/testCert", func(c echo.Context) error {
		elastic, _ := elasticsearch.NewDefaultClient()
		log.Println(elasticsearch.Version)
		log.Println(elastic.Info())
		return c.String(http.StatusOK, "test https certs")
	})
	userGroup.POST("/register", uc.Register)
	// userGroup.POST("/login")
	// userGroup.PUT("/editDataPerson")
}
