package service

import (
	"projecttemplatebyfrans/modules/role"
	"projecttemplatebyfrans/modules/role/model"
	"projecttemplatebyfrans/schemas"
)

type RolesService struct {
	RoleRepository role.Repository
}

func InitRolesService(RoleRepository role.Repository) role.Service {
	return &RolesService{
		RoleRepository: RoleRepository,
	}
}

func (u *RolesService) CreateRoleService(user schemas.CreateRoleRequest) (error, int64) {
	roleModel := &model.Roles{
		Name: &user.Name,
	}
	err, id := u.RoleRepository.CreateRoleRepository(roleModel)
	if err != nil {
		return err, 0
	}

	return nil, id
}

func (u *RolesService) GetRolesService() (*[]model.Roles, error) {
	roles, err := u.RoleRepository.GetRolesRepository()
	if err != nil {
		return nil, err
	}

	return roles, nil
}
