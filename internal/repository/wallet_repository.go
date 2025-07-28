package repository

import (
	"wallet_api/internal/errors"

	"github.com/jmoiron/sqlx"
)

// IWalletRepository определяет интерфейс для работы с кошельками.
type IWalletRepository interface {
	// CheckWalletAddress выполняет проверку существования кошелька по адресу.
	CheckWalletAddress(address string) (bool, error)
	// GetBalance возвращает баланс кошелька.
	GetBalance(address string) (float64, error)
}

// WalletRepository реализует IWalletRepository.
type WalletRepository struct {
	db *sqlx.DB
}

// NewWalletRepository возвращает экземпляр WalletRepository.
func NewWalletRepository(db *sqlx.DB) *WalletRepository {
	return &WalletRepository{db}
}

func (rep *WalletRepository) CheckWalletAddress(address string) (bool, error) {
	var exists bool
	err := rep.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM wallets WHERE address = $1)", address)
	if err != nil {
		return false, errors.ErrDatabase
	}
	return exists, nil
}

func (rep *WalletRepository) GetBalance(address string) (float64, error) {
	var balance float64
	err := rep.db.Get(&balance, "SELECT balance FROM wallets WHERE address = $1", address)
	if err != nil {
		return 0, errors.ErrDatabase
	}
	return balance, nil
}
