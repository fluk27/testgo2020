package routes

import (
	"github.com/labstack/echo/v4"
)

func init() {
	e := echo.New()
	AuthRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
