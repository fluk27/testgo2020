package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	e := echo.New()
	UserRoute(e)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8000"))
}
