package role

import (
	"projecttemplatebyfrans/modules/role/model"
	"projecttemplatebyfrans/schemas"
)

type Service interface {
	CreateRoleService(user schemas.CreateRoleRequest) (error, int64)
	GetRolesService() (*[]model.Roles, error)
}

type Repository interface {
	CreateRoleRepository(role *model.Roles) (error, int64)
	GetRolesRepository() (*[]model.Roles, error)
}
