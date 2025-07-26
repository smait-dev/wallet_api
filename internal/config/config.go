// Package config реализация работы с настройками проекта
package config

import "github.com/spf13/viper"

type Config struct {
	MaxRowsLimit int
}

func NewConfig() *Config {
	return &Config{
		MaxRowsLimit: viper.GetInt("db.max_rows_query_limit"),
	}
}
