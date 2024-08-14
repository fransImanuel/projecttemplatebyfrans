package service

import (
	"projecttemplatebyfrans/modules/users"
	"projecttemplatebyfrans/modules/users/model"
	"projecttemplatebyfrans/schemas"
)

type UsersService struct {
	UsersRepository users.Repository
}

func InitUsersRepository(UsersRepository users.Repository) users.Service {
	return &UsersService{
		UsersRepository: UsersRepository,
	}
}

func (u *UsersService) CreateUserService(user schemas.CreateUserRequest) (error, int64) {
	userModel := &model.Users{
		Name:   &user.Name,
		Age:    &user.Age,
		Email:  &user.Email,
		Phone:  &user.Phone,
		RoleID: &user.RoleID,
	}
	err, id := u.UsersRepository.CreateUserRepository(userModel)
	if err != nil {
		return err, 0
	}

	return nil, id
}

func (u *UsersService) GetUsersService() (*[]model.Users, error) {
	users, err := u.UsersRepository.GetUsersRepository()
	if err != nil {
		return nil, err
	}

	return users, nil
}
