package userdto

import (
	"github.com/nuttchai/go-rest/internal/model"
	"github.com/nuttchai/go-rest/internal/util/validators"
)

type UpdateUserDTO struct {
	Id        string `json:"id" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required"`
}

func (dto *UpdateUserDTO) ToModel() *model.User {
	return &model.User{
		Id:        dto.Id,
		UserName:  dto.UserName,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
	}
}

func (dto *UpdateUserDTO) Validate() error {
	return validators.ValidateStruct(dto)
}
