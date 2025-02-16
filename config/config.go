package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize DB
	db, err = InitializeDb()
	if err != nil {
		return fmt.Errorf("error initialize mssql: %v", err)
	}

	return nil
}

func GetDb() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
