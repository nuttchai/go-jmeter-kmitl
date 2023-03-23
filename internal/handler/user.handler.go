package handler

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/nuttchai/go-rest/internal/constant"
	shareddto "github.com/nuttchai/go-rest/internal/dto/shared"
	dto "github.com/nuttchai/go-rest/internal/dto/user"
	ihandler "github.com/nuttchai/go-rest/internal/handler/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
	"github.com/nuttchai/go-rest/internal/util/api"
)

type TUserHandler struct {
	UserService iservice.IUserService
}

var (
	UserHandler ihandler.IUserHandler
)

func initUserHandler(userHandler *TUserHandler) {
	UserHandler = userHandler
}

func (h *TUserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		err := errors.New(constant.NoGivenId)
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	user, err := h.UserService.GetUser(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(user, constant.GetUserSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TUserHandler) GetUsers(c echo.Context) error {
	user, err := h.UserService.GetUsers()
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(user, constant.GetUsersSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TUserHandler) CreateUser(c echo.Context) error {
	var userDto dto.CreateUserDTO
	if err := api.DecodeDTO(c, &userDto); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	users, err := h.UserService.CreateUser(&userDto)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(users, constant.CreateUserSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TUserHandler) UpdateUser(c echo.Context) error {
	var userDto dto.UpdateUserDTO
	if err := api.DecodeDTO(c, &userDto); err != nil {
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	users, err := h.UserService.UpdateUser(&userDto)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(users, constant.UpdateUserSuccessMsg)
	return c.JSON(res.Status, res)
}

func (h *TUserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		err := errors.New(constant.NoGivenId)
		jsonErr := api.BadRequestError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	err := h.UserService.DeleteUser(id)
	if err != nil {
		jsonErr := api.InternalServerError(err)
		return c.JSON(jsonErr.Status, jsonErr)
	}

	res := api.SuccessResponse(
		&shareddto.ValidatorResultDTO{
			IsSuccess: true,
			Action:    "delete_user",
		},
		constant.DeleteUserSuccessMsg,
	)

	return c.JSON(res.Status, res)
}
