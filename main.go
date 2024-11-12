package main

import (
	"net/http"

	"github.com/heartlezz7/go-echo-todolist/database"
	"github.com/heartlezz7/go-echo-todolist/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog/log"
)

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// difine skipper func
		// Skipper:      DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	for key, v := range route.Router {
		log.Info().Msg(key + " route")
		v(e)
	}

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	database.ConnectPostgres()

	e.Logger.Fatal(e.Start(":8080"))
}
