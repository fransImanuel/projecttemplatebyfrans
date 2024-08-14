package model

import (
	"fmt"
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

	var count int64
	if err := tx.Raw(`SELECT count(id)  FROM "Roles"`).Scan(&count).Error; err != nil {
		fmt.Println(err)
	}
	if count == 0 {
		// if err := tx.Exec(`INSERT INTO "Roles" (id, created_at, created_user, updated_at, updated_user, deleted_at, deleted_user, code, "name", value) VALUES (1, ?, 'system', null, null, null, null,'MAX_RANGE_SURVEY', null, '150')`, time.Now()).Error; err != nil {
		// 	logrus.Errorln("❌ Error Insert RANGE_TIME_SURVEY in Config : ", err.Error())
		// }
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
