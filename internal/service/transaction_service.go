package service

import (
	"errors"
	"wallet_api/internal/entity"
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
}

// NewTransactionService возвращает экземпляр TransactionService.
func NewTransactionService(repository repository.ITransactionRepository) *TransactionService {
	return &TransactionService{
		transactionRepository: repository,
	}
}

func (ts *TransactionService) Send(from string, to string, amount float64) error {
	if amount <= 0 {
		return errors.New("tmp error")
	}
	return ts.transactionRepository.Send(from, to, amount)
}

func (ts *TransactionService) GetLastTransactions(count int) ([]entity.Transaction, error) {
	// if count > MaxRowsLimit {
	// 	count = MaxRowsLimit
	// }
	transactions, err := ts.transactionRepository.GetLastTransactions(count)
	if err != nil {
		return nil, errors.New("tmp error")
	}
	return transactions, err
}
