package repository

import "github.com/nuttchai/go-rest/internal/shared/config"

func Init() {
	initUserRepository(&TUserRepository{
		DB: config.GetAppDB(),
	})
}
