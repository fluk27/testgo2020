package routes

import (
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	//production
	port := os.Getenv("PORT")
	// dev mode
	port = "8000"
	e := echo.New()
	UserRoute(e)
	// e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":" + port))
	// e.Logger.Fatal(e.StartTLS(":8080", "./cert.pem", "./key.pem"))
}
