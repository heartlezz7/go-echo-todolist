package main

import (
	"net/http"

	"github.com/heartlezz7/go-echo-project/database"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	database.ConnectPostgres()

	e.Logger.Fatal(e.Start(":8080"))
}
