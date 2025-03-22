package config

import (
	"oopLab1/pkg/logger"
	"sync"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string
	Port int
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	TimeZone string
}

type Config struct {
	Server   *ServerConfig
	Database *DBConfig
}

var once sync.Once
var config *Config

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./")
	})

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal(err.Error())
	}

	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatal(err.Error())
	}

	return config
}
