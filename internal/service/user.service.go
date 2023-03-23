package service

import (
	dto "github.com/nuttchai/go-rest/internal/dto/user"
	"github.com/nuttchai/go-rest/internal/model"
	irepository "github.com/nuttchai/go-rest/internal/repository/interface"
	iservice "github.com/nuttchai/go-rest/internal/service/interface"
)

type TUserService struct {
	Repository irepository.IUserRepository
}

var (
	UserService iservice.IUserService
)

func initUserService(userService *TUserService) {
	UserService = userService
}

func (s *TUserService) GetUser(id string) (*model.User, error) {
	return s.Repository.FindOne(id)
}

func (s *TUserService) GetUsers() ([]*model.User, error) {
	return s.Repository.FindAll()
}

func (s *TUserService) CreateUser(userDto *dto.CreateUserDTO) (*model.User, error) {
	user := userDto.ToModel()
	return s.Repository.CreateOne(user)
}

func (s *TUserService) UpdateUser(userDto *dto.UpdateUserDTO) (*model.User, error) {
	user := userDto.ToModel()
	return s.Repository.UpdateOne(user)
}

func (s *TUserService) DeleteUser(id string) error {
	return s.Repository.DeleteOne(id)
}
