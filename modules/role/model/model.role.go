package model

import (
	"projecttemplatebyfrans/constant"
	"projecttemplatebyfrans/schemas"
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	schemas.FullAudit

	Name     *string `json:"name,omitempty"  gorm:"column:Name"`
	IsActive *bool   `json:"is_active,omitempty"  gorm:"column:IsActive;default:true"`
}

// ? this is just gorm way of custom table name
func (t *Roles) TableName() string {
	return constant.TABLE_ROLES_NAME
}

func (Roles) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Roles{})
	if err != nil {
		return err
	}

	return nil
}

func (t *Roles) InitAudit(operation string /*, user string, user_id int64*/) {
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
