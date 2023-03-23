package iservice

import (
	dto "github.com/nuttchai/go-rest/internal/dto/user"
	"github.com/nuttchai/go-rest/internal/model"
)

type IUserService interface {
	GetUser(id string) (*model.User, error)
	GetUsers() ([]*model.User, error)
	CreateUser(user *dto.CreateUserDTO) (*model.User, error)
	UpdateUser(user *dto.UpdateUserDTO) (*model.User, error)
	DeleteUser(id string) error
}
