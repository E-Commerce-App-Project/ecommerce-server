package driver

import (
	"fmt"

	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewMysqlDatabase return gorp dbmap object with MySQL options param
func NewMysqlDatabase(option config.DatabaseConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		option.Username,
		option.Password,
		option.Address,
		option.Port,
		option.Name,
	)
	gorm, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		// silently logger
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return gorm, err
}
