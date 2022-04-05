package database

import "gorm.io/gorm"

func Migrate(dbInstance *gorm.DB) {
	// Regiser migration
	dbInstance.AutoMigrate(&UserEntity{})
	dbInstance.AutoMigrate(&ProductEntity{})
	dbInstance.AutoMigrate(&CartItemEntity{})
	dbInstance.AutoMigrate(&TransactionEntity{})
	dbInstance.AutoMigrate(&TransactionDetailEntity{})
}
