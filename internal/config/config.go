package config

import (
	"github.com/spf13/viper"
)

var config Config

// Config contains all configurations used by application
type Config struct {
	Database   DatabaseConfig
	API        APIConfig
	Log        LogConfig
	Eureka     EurekaConfig
	StaticRoot string
}

// EurekaConfig contains all configurations used by eureka
type EurekaConfig struct {
	GatewayURL string
	AppID      string
	InstanceID string
	HostName   string
	IPAdress   string
	Port       int
	TTL        uint
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
	// set config options(all variables required)
	viper.SetEnvPrefix("COOKER")
	viper.AutomaticEnv()
	config = Config{
		Database: DatabaseConfig{
			URL: viper.GetString("database_url"),
		},
		API: APIConfig{
			EnableCORS: true,
			Port:       viper.GetInt("api_port"),
		},
		Log: LogConfig{
			Level: viper.GetString("log_level"),
		},
		Eureka: EurekaConfig{
			GatewayURL: viper.GetString("eureka_gateway_url"),
			InstanceID: viper.GetString("eureka_instance_id"),
			AppID:      viper.GetString("eureka_instance_app_id"),
			HostName:   viper.GetString("eureka_instance_hostname"),
			IPAdress:   viper.GetString("eureka_instance_ipaddress"),
			Port:       viper.GetInt("eureka_instance_port"),
			TTL:        viper.GetUint("eureka_instance_ttl"),
		},
		StaticRoot: viper.GetString("static_path"),
	}
}
