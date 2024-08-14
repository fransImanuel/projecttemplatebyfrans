package repository

import (
	"projecttemplatebyfrans/constant"
	"projecttemplatebyfrans/modules/role"
	"projecttemplatebyfrans/modules/role/model"

	"gorm.io/gorm"
)

type RolesRepository struct {
	DBPostgres *gorm.DB
	//DBMongoDB
	//DBMinio, etc
}

func InitRolesRepository(db *gorm.DB) role.Repository {
	return &RolesRepository{
		DBPostgres: db,
	}
}

func (u *RolesRepository) CreateRoleRepository(role *model.Roles) (error, int64) {
	db := u.DBPostgres

	role.InitAudit(constant.OPERATION_SQL_INSERT)

	results := db.Create(&role)
	if results.Error != nil {
		return results.Error, 0
	}

	return nil, role.ID

}

func (u *RolesRepository) GetRolesRepository() (*[]model.Roles, error) {
	var roles *[]model.Roles
	db := u.DBPostgres

	// Get all records
	results := db.Find(&roles)
	// SELECT * FROM users;
	if results.Error != nil {
		return nil, results.Error
	}

	return roles, nil
}
