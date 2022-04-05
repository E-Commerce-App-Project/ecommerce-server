package appcontext

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/config"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/driver"
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
