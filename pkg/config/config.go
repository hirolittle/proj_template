package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config holds the configuration for the application
type Config struct {
	LogConfig LogConfig `yaml:"logConfig"`
}

// LogConfig holds the logging configuration
type LogConfig struct {
	Output     string `yaml:"output"`
	FileName   string `yaml:"fileName"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	Compress   bool   `yaml:"compress"`
}

// Cfg holds the global configuration instance
var Cfg Config

// InitConfig initializes the configuration by reading from a YAML file
func InitConfig() {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath("./conf")
	config.SetConfigType("yaml")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := config.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}
}
