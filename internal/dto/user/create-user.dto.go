package userdto

import (
	"github.com/nuttchai/go-rest/internal/model"
	"github.com/nuttchai/go-rest/internal/util/validators"
)

type CreateUserDTO struct {
	UserName  string `json:"username" default:""`
	FirstName string `json:"firstname" default:""`
	LastName  string `json:"lastname" default:""`
	Email     string `json:"email" default:""`
	Phone     string `json:"phone" default:""`
}

func (dto *CreateUserDTO) ToModel() *model.User {
	return &model.User{
		UserName:  dto.UserName,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
	}
}

func (dto *CreateUserDTO) Validate() error {
	return validators.ValidateStruct(dto)
}
