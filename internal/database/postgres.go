// Package database предоставляет функционал подключения к базе данных и выполнения миграций.
package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
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

// RunMigrations выполняет миграции.
func RunMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// DBSeedWallets выполняет стартовое наполнение БД кошельками
func DBSeedWallets(db *sqlx.DB) error {
	var count int
	err := db.Get(&count, "SELECT COUNT(*) FROM wallets")
	if err != nil {
		return err
	}
	if count != 0 {
		return nil
	}

	sql, err := os.ReadFile("db_seeds/wallets_seeder.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sql))
	return err
}
