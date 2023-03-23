package router

import (
	"github.com/labstack/echo"
	handler "github.com/nuttchai/go-rest/internal/handler"
	"github.com/nuttchai/go-rest/internal/util/api"
)

func initUserRouter(e *echo.Echo) {
	e.GET(api.CreatePath("user/:id"), handler.UserHandler.GetUser)
	e.GET(api.CreatePath("user"), handler.UserHandler.GetUsers)
	e.POST(api.CreatePath("user"), handler.UserHandler.CreateUser)
	e.PUT(api.CreatePath("user"), handler.UserHandler.UpdateUser)
	e.DELETE(api.CreatePath("user/:id"), handler.UserHandler.DeleteUser)
}
