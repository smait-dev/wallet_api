// Package database предоставляет функционал подключения к базе данных и выполнения миграций.
package database

import (
	"fmt"
	"os"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// ConnectPostgres устанавливает соединение с базой данных PostgreSQL.
// Возвращает подключение к БД и ошибку.
func ConnectPostgres() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", GetPostgresDSN())
}

// GetPostgresDSN формирует строковую ссылку для подключения.
func GetPostgresDSN() string {
	host := viper.GetString("db.host")
	port := viper.GetInt("db.port")
	username := viper.GetString("db.username")
	pass := os.Getenv("DB_PASSWORD")
	dbname := viper.GetString("db.dbname")
	sslmode := viper.GetString("db.sslmode")

	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		username, pass, host, port, dbname, sslmode,
	)
}

