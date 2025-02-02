package config

import "gorm.io/gorm"

type Config struct {
	DB *gorm.DB
}

func NewConfig() *Config {
	config := Config{}

	config.DB = connectPostgres()

	return &config
}
