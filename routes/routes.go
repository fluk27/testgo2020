package routes

import (
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	port:=os.Getenv("PORT")
	e := echo.New()
	UserRoute(e)

	e.Logger.Fatal(e.Start(":"+port))
}
