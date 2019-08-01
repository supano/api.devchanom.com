package main

import (
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://devchanom.test", "https://devchanom.com"},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello From Dev Chanom")
	})

	e.Logger.Fatal(e.Start(":8100"))
}
