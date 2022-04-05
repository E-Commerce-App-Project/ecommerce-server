package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	lock          = &sync.Mutex{}
	_, b, _, _    = runtime.Caller(0)
	basepath      = filepath.Dir(b)
	defaultConfig *viper.Viper
)

// Provider the config provider
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
	GetDatabaseConfig() DatabaseConfig
	GetAppConfig() AppConfig
}

type DatabaseConfig struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Name     string `mapstructure:"DB_NAME"`
	Address  string `mapstructure:"DB_ADDRESS"`
	Port     int    `mapstructure:"DB_PORT"`
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type AppConfig struct {
	Host      string `mapstructure:"APP_HOST"`
	Port      int    `mapstructure:"APP_PORT"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

type RedisConfig struct {
	Address  string `mapstructure:"REDIS_ADDRESS"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

type AppConfigs struct {
	*viper.Viper
}

func (c *AppConfigs) GetDatabaseConfig() DatabaseConfig {
	var config DatabaseConfig
	lock.Lock()
	defer lock.Unlock()
	c.Unmarshal(&config)
	return config
}

func (c *AppConfigs) GetAppConfig() AppConfig {
	var config AppConfig
	lock.Lock()
	defer lock.Unlock()
	c.Unmarshal(&config)
	return config
}

func init() {
	defaultConfig = readViperConfig()
}

func readViperConfig() *viper.Viper {
	v := viper.New()
	parentDir := filepath.Dir(basepath)
	v.AddConfigPath(parentDir)
	v.SetConfigName(".env")
	v.SetConfigType("env")
	v.AutomaticEnv()
	v.SetDefault("APP_HOST", os.Getenv("APP_HOST"))
	v.SetDefault("APP_PORT", os.Getenv("APP_PORT"))
	v.SetDefault("JWT_SECRET", os.Getenv("JWT_SECRET"))
	v.SetDefault("API_ENDPOINT", os.Getenv("API_ENDPOINT"))
	v.SetDefault("DB_DRIVER", os.Getenv("DB_DRIVER"))
	v.SetDefault("DB_NAME", os.Getenv("DB_NAME"))
	v.SetDefault("DB_ADDRESS", os.Getenv("DB_ADDRESS"))
	v.SetDefault("DB_PORT", os.Getenv("DB_PORT"))
	v.SetDefault("DB_USERNAME", os.Getenv("DB_USERNAME"))
	v.SetDefault("DB_PASSWORD", os.Getenv("DB_PASSWORD"))

	err := v.ReadInConfig()
	if err == nil {
		fmt.Printf("Using config file: %s \n\n", v.ConfigFileUsed())
	} else {
		fmt.Printf("Config error: %s", err)
	}

	return v
}

// Config return provider so that you can read config anywhere
func Config() Provider {
	return &AppConfigs{defaultConfig}
}
