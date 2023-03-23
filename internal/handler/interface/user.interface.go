package ihandler

import "github.com/labstack/echo"

type IUserHandler interface {
	GetUser(c echo.Context) error
	GetUsers(c echo.Context) error
	CreateUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}
