package appcontext

import (
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/driver"
	"github.com/gomodule/redigo/redis"
	"gorm.io/gorm"
)

// AppContext the app context struct
type AppContext struct {
	config config.Provider
}

// NewAppContext initiate appcontext object
func NewAppContext(config config.Provider) *AppContext {
	return &AppContext{
		config: config,
	}
}

// GetDBInstance getting gorp instance, param: dbType
func (a *AppContext) GetDBInstance() (*gorm.DB, error) {
	db, err := driver.NewMysqlDatabase(a.config.GetDatabaseConfig())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a *AppContext) GetCachedPool() (pool *redis.Pool, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = commons.ErrCacheConn
		}
	}()

	pool = driver.NewCache(a.getCacheOption())

	return
}

func (a *AppContext) getCacheOption() driver.CacheOption {
	cachedConfig := a.config.GetRedisConfig()
	return driver.CacheOption{
		Host:               cachedConfig.Address,
		Port:               cachedConfig.Port,
		Password:           cachedConfig.Password,
		DialConnectTimeout: time.Second * 5,
		Namespace:          "0",
		ReadTimeout:        time.Second * 5,
		WriteTimeout:       time.Second * 5,
		MaxIdle:            5,
		MaxActive:          1000,
		IdleTimeout:        time.Second * 10,
		Wait:               true,
		MaxConnLifetime:    0,
	}
}
