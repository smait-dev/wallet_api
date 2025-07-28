package service

import (
	"wallet_api/internal/config"
	"wallet_api/internal/entity"
	"wallet_api/internal/errors"
	"wallet_api/internal/repository"
)

// ITransactionService определяет интерфейс бизнес-логики для работы с транзакциями.
type ITransactionService interface {
	// Send выполненяет отправку средств между кошельками.
	Send(from string, to string, amount float64) error
	// GetLastTransactions возвращает записи транзакции в количестве count.
	GetLastTransactions(count int) ([]entity.Transaction, error)
}

// TransactionService реализует ITransactionService.
type TransactionService struct {
	transactionRepository repository.ITransactionRepository
	walletService         IWalletService
	config                *config.Config
}

// NewTransactionService возвращает экземпляр TransactionService.
func NewTransactionService(repository repository.ITransactionRepository, walletService IWalletService, cfg *config.Config) *TransactionService {
	return &TransactionService{
		transactionRepository: repository,
		walletService:         walletService,
		config:                cfg,
	}
}

func (ts *TransactionService) Send(from string, to string, amount float64) error {
	if amount <= 0 {
		return errors.ErrInvalidAmount
	}

	if err := ts.walletService.ValidateAddresses(from, to); err != nil {
		return err
	}
	return ts.transactionRepository.Send(from, to, amount)
}

func (ts *TransactionService) GetLastTransactions(count int) ([]entity.Transaction, error) {
	if count > ts.config.MaxRowsLimit {
		count = ts.config.MaxRowsLimit
	}
	transactions, err := ts.transactionRepository.GetLastTransactions(count)
	if err != nil {
		return nil, errors.ErrDatabase
	}
	return transactions, err
}
