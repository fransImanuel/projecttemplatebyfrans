package repository

import (
	"projecttemplatebyfrans/constant"
	"projecttemplatebyfrans/modules/users"
	"projecttemplatebyfrans/modules/users/model"

	"gorm.io/gorm"
)

type UsersRepository struct {
	DBPostgres *gorm.DB
	//DBMongoDB
	//DBMinio, etc
}

func InitUsersRepository(db *gorm.DB) users.Repository {
	return &UsersRepository{
		DBPostgres: db,
	}
}

func (u *UsersRepository) CreateUserRepository(user *model.Users) (error, int64) {
	db := u.DBPostgres

	user.InitAudit(constant.OPERATION_SQL_INSERT)

	results := db.Create(&user)
	if results.Error != nil {
		return results.Error, 0
	}

	return nil, user.ID

}

func (u *UsersRepository) GetUsersRepository() (*[]model.Users, error) {
	var users *[]model.Users
	db := u.DBPostgres

	// Get all records
	results := db.Find(&users)
	// SELECT * FROM users;
	if results.Error != nil {
		return nil, results.Error
	}

	return users, nil
}
