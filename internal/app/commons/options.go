package commons

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/logger"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

const (
	CTX_USER_KEY = "user"
)

// Options common option for all object that needed
type Options struct {
	Config    config.Provider
	DbMysql   *gorm.DB
	CachePool *redis.Pool
	Logger    logger.Logger
}
