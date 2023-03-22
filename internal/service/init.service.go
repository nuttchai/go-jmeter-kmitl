package service

import "github.com/nuttchai/go-rest/internal/repository"

func Init() {
	initUserService(&TUserService{
		Repository: repository.UserRepository,
	})
}
