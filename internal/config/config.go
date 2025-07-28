// Package config реализация работы с настройками проекта
package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	MaxRowsLimit int
}

func NewConfig() *Config {
	return &Config{
		MaxRowsLimit: viper.GetInt("db.max_rows_query_limit"),
	}
}

func InitConfig() *Config {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return NewConfig()
}
