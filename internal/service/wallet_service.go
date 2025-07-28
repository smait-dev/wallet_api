package service

import (
	"strings"
	"wallet_api/internal/errors"
	"wallet_api/internal/repository"
)

// IWalletService определяет интерфейс бизнес-логики для работы с кошельками.
type IWalletService interface {
	// ValidateAddresses валидирует адреса для перевода.
	ValidateAddresses(from string, to string) error
	// GetBalance возвращает баланс кошелька.
	GetBalance(address string) (float64, error)
	// ValidateBalance(address string) error
}

// WalletService реализует IWalletService.
type WalletService struct {
	walletRepository repository.IWalletRepository
}

// NewWalletService возвращает новый экземпляр WalletService.
func NewWalletService(repository repository.IWalletRepository) *WalletService {
	return &WalletService{repository}
}

func (ws *WalletService) ValidateAddresses(from string, to string) error {
	from = strings.TrimSpace(from)
	to = strings.TrimSpace(to)
	switch {
	case from == "":
		return errors.ErrInvalidSender
	case to == "":
		return errors.ErrInvalidReceiver
	case from == to:
		return errors.ErrSelfTransfer
	}

	isValid, err := ws.walletRepository.CheckWalletAddress(from)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.ErrInvalidSender
	}

	isValid, err = ws.walletRepository.CheckWalletAddress(to)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.ErrInvalidReceiver
	}

	return nil
}

func (ws *WalletService) GetBalance(address string) (float64, error) {
	isValid, err := ws.walletRepository.CheckWalletAddress(address)
	if !isValid || err != nil {
		return 0, errors.ErrWalletNotFound
	}

	balance, err := ws.walletRepository.GetBalance(address)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
