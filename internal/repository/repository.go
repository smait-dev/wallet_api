// Package repository предоставляет слой доступа к базе данных.
package repository

import (
	"wallet_api/internal/config"
	"github.com/jmoiron/sqlx"
)

// Repository предоставляет доступ к репозиториям транзакций и кошельков.
type Repository struct {
	Transaction ITransactionRepository
	Wallet      IWalletRepository
	Config      *config.Config
}

// NewRepository создает экземпляр репозитория.
func NewRepository(db *sqlx.DB, cfg *config.Config) *Repository {
	return &Repository{
		Transaction: NewTransactionRepository(db),
		Wallet:      NewWalletRepository(db),
		Config:      cfg,
	}
}
