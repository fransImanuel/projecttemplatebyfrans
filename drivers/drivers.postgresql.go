package drivers

import (
	"fmt"
	"os"
	"projecttemplatebyfrans/schemas"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	roleModel "projecttemplatebyfrans/modules/role/model"
	userModel "projecttemplatebyfrans/modules/users/model"
)

/*
? I use the GORM framework for database operations.
? If you compare it to a bare Go SQL driver, the bare driver is indeed faster.
? However, using GORM allows me to accomplish tasks that would take much longer with
? just the bare driver, such as automating migrations and writing simpler query code.
*/

func SetupDBSQL(config schemas.SchemaEnvironment) (*gorm.DB, error) {
	logrus.Debug("🔌 Connecting into Database Postgres")
	dbHost := config.DB_HOST
	dbUsername := config.DB_USER
	dbPassword := config.DB_PASS
	dbName := config.DB_NAME
	dbPort := config.DB_PORT
	dbSSLMode := config.DB_SSLMODE
	timezone := config.TIMEZONE

	path := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		dbHost, dbUsername, dbPassword, dbName, dbPort, dbSSLMode, timezone)

	db, err := gorm.Open(postgres.Open(path), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		defer logrus.Errorln("❌ Error Connect into Database Postgres", err.Error())

		return nil, err
	}

	postgreSQL, err := db.DB()
	// Set connection pool parameters
	postgreSQL.SetMaxOpenConns(10)   // Maximum number of open connections
	postgreSQL.SetMaxIdleConns(5)    // Maximum number of idle connections
	postgreSQL.SetConnMaxLifetime(0) // Connection lifetime (0 means connections are reused indefinitely)

	if os.Getenv("GO_ENV") == "development" {
		db.Debug()
	}

	if err != nil {
		logrus.Errorln("❌ Error Migrate ", err.Error())
		return nil, err
	}

	fmt.Println("💚 Connect into Database Postgres Success")

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	// err = db.AutoMigrate(
	// 	//mTasklistDetail.TasklistDetail{},
	// 	//model3.TRVisit{},
	// 	//&entities.EntityVenue{},
	// 	sopModel.SOP{},
	// 	instructionModel.Instruction{},
	// 	confModel.Config{},
	// )

	if err := db.Transaction(func(tx *gorm.DB) error {

		if err := roleModel.Roles.Migrate(roleModel.Roles{}, tx); err != nil {
			return nil
		}
		if err := userModel.Users.Migrate(userModel.Users{}, tx); err != nil {
			return nil
		}

		return nil
	}); err != nil {
		panic("Fail to Migrate")
	}

}
