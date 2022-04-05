package database

import "gorm.io/gorm"

func Migrate(dbInstance *gorm.DB) {
	// Regiser migration
	dbInstance.AutoMigrate()
}
