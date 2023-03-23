package irepository

import "github.com/nuttchai/go-rest/internal/model"

type IUserRepository interface {
	FindOne(id string) (*model.User, error)
	FindAll() ([]*model.User, error)
	CreateOne(user *model.User) (*model.User, error)
	UpdateOne(user *model.User) (*model.User, error)
	DeleteOne(id string) error
}
