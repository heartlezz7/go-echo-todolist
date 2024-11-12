package route

import (
	"github.com/heartlezz7/go-echo-todolist/handler"
	"github.com/labstack/echo"
)

type router func(e *echo.Echo)

var Router = map[string]router{"auth": authRoute}

func authRoute(e *echo.Echo) {

	g := e.Group("/auth")

	// improve sign up with hash
	g.POST("/sign-up", handler.CreateUserHandler)

	g.GET("/get-user/:id", handler.GetUserDetail)
	g.POST("/login", handler.Login)

}
