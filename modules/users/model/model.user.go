package model

import (
	"projecttemplatebyfrans/constant"
	role "projecttemplatebyfrans/modules/role/model"
	"projecttemplatebyfrans/schemas"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	schemas.FullAudit

	Name     *string     `json:"name,omitempty"  gorm:"column:Name"`
	Age      *string     `json:"age,omitempty"  gorm:"column:Age"`
	Email    *string     `json:"email,omitempty"  gorm:"column:Email"`
	Phone    *string     `json:"phone,omitempty"  gorm:"column:Phone"`
	IsActive *bool       `json:"is_active,omitempty"  gorm:"column:IsActive;default:true"`
	RoleID   *int64      `json:"role_id" gorm:"column:RoleID"`
	Role     *role.Roles `json:"role" gorm:"foreignKey:RoleID"`
}

// ? this is just gorm way of custom table name
func (t *Users) TableName() string {
	return constant.TABLE_USERS_NAME
}

func (Users) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Users{})
	if err != nil {
		return err
	}

	return nil
}

func (t *Users) InitAudit(operation string /*, user string, user_id int64*/) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		// t.CreatedByUserName = user
		t.CreatedTime = timeNow
		// t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_UPDATE:
		// t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_DELETE:
		// t.DeletedByUserId = &user_id
		t.DeletedTime = gorm.DeletedAt{Time: timeNow, Valid: true}
	}
}
