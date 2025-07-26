// Package repository предоставляет слой доступа к базе данных.
package repository

import (
	"github.com/jmoiron/sqlx"
)

// Repository предоставляет доступ к репозиториям транзакций и кошельков.
type Repository struct {
	Transaction ITransactionRepository
	Wallet      IWalletRepository
}

// NewRepository создает экземпляр репозитория.
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Transaction: NewTransactionRepository(db),
		Wallet:      NewWalletRepository(db),
	}
}
