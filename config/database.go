package config

import (
	"github.com/MarcosSmeets/Goportunities/schemas"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitializeDb() (*gorm.DB, error) {
	logger := GetLogger("mssql")

	// connect to database
	db, err := gorm.Open(sqlserver.Open("sqlserver://nycbank_admin:nyc607060@@nycbank.database.windows.net?database=wellify_dashboard"), &gorm.Config{})
	if err != nil {
		logger.ErrorF("mssql fail to initialize db %v", err.Error())
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("mssql fail to migrate %v", err.Error())
		return nil, err
	}

	// Return the DB
	return db, nil
}
