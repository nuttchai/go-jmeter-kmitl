package handler

import (
	"github.com/nuttchai/go-rest/internal/service"
)

func Init() {
	initUserHandler(&TUserHandler{
		UserService: service.UserService,
	})
}
