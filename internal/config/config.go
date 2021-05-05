package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config Config

// Config contains all configurations used by application
type Config struct {
	Database DatabaseConfig
	API      APIConfig
	Log      LogConfig
}

// DatabaseConfig contains all configurations used by datbase
type DatabaseConfig struct {
	URL string
}

// APIConfig contains all configuations used by httpserver
type APIConfig struct {
	EnableCORS bool
	Port       int
}

// LogConfig contains LogLevel, which can be debug,info,warning,fatal,panic
type LogConfig struct {
	Level string
}

// Get returns configurations initialied by config.file. Has zero-value if key is not set
func Get() *Config {
	return &config
}

func init() {
	// set config options
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		// panic behaves like throw: call defer, go downward until recover
		panic(fmt.Errorf("fatal errors reading config file %s", err))
	}
	// returns config. filled with zero-values if config is empty
	config = Config{
		Database: DatabaseConfig{
			URL: viper.GetString("database.url"),
		},
		API: APIConfig{
			EnableCORS: viper.GetBool("api.enable_cors"),
			Port:       viper.GetInt("api.port"),
		},
		Log: LogConfig{
			Level: viper.GetString("log.level"),
		},
	}
}
