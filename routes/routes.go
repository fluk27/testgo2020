package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
)

func init() {
	//production
	// port := os.Getenv("PORT")
	// dev mode
	port := "8000"
	e := echo.New()
	UserRoute(e)
	e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":" + port))
}
