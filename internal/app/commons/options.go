package commons

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/logger"
	"gorm.io/gorm"
)

// Options common option for all object that needed
type Options struct {
	Config  config.Provider
	DbMysql *gorm.DB
	Logger  logger.Logger
}
