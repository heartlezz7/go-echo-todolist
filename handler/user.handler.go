package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/heartlezz7/go-echo-todolist/model"
	"github.com/heartlezz7/go-echo-todolist/service"
	"github.com/labstack/echo"
)

func CreateUserHandler(c echo.Context) error {

	var user model.CreateUser

	var err error

	err = c.Bind(&user)

	if err != nil {
		return c.String(http.StatusInternalServerError, "validate fail")
	}

	err = service.CreateUserService(user)

	if err != nil {
		return c.String(http.StatusBadRequest, "create fail")
	}

	return c.String(http.StatusCreated, "create success")

}

func GetUserDetail(c echo.Context) error {

	var err error
	var id int
	params := c.Param("id")

	id, err = strconv.Atoi(params)
	if err != nil {
		return err
	}
	var user model.UserData

	err = service.GetUserDetailService(id, &user)
	fmt.Println(user)
	if err != nil {
		return c.String(http.StatusNotFound, "fail")
	}

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {

	var err error
	var userLogin model.UserLogin

	err = c.Bind(&userLogin)
	if err != nil {
		return c.String(http.StatusBadRequest, "Validate fail")
	}

	err = service.LoginService(userLogin, userLogin)

	return nil
}
