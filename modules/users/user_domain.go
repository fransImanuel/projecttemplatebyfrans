package users

import (
	"projecttemplatebyfrans/modules/users/model"
	"projecttemplatebyfrans/schemas"
)

type Repository interface {
	CreateUserRepository(user *model.Users) (error, int64)
	GetUsersRepository() (*[]model.Users, error)
}
type Service interface {
	CreateUserService(user schemas.CreateUserRequest) (error, int64)
	GetUsersService() (*[]model.Users, error)
}
