package tests

import (
	"testing"
	"wallet_api/internal/config"
	"wallet_api/internal/entity"
	"wallet_api/internal/errors"
	"wallet_api/internal/service"
)

var cfg = config.Config{MaxRowsLimit: 1000}

// Мок реализации репозитория транзакций
type mockTransactionRepo struct {
	sendErr          error
	lastTransactions []entity.Transaction
}

func (m *mockTransactionRepo) Send(from, to string, amount float64) error {
	return m.sendErr // nil, либо имитация возврата ошибки
}
func (m *mockTransactionRepo) GetLastTransactions(count int) ([]entity.Transaction, error) {
	return m.lastTransactions, nil
}

type mockWalletService struct {
	validateErr error
}

func (m *mockWalletService) ValidateAddresses(from, to string) error {
	return m.validateErr
}

func (m *mockWalletService) GetBalance(address string) (float64, error) {
	return 0, nil
}

func TestSend_Success(t *testing.T) {
	mockTransactionRepo := &mockTransactionRepo{}
	mockWalletService := &mockWalletService{}
	service := service.NewTransactionService(mockTransactionRepo, mockWalletService, &cfg)

	err := service.Send("A", "B", 10)
	if err != nil {
		t.Errorf("Ошибка: %v", err)
	}
}

func TestSend_ValidationError(t *testing.T) {
	mockTransactionRepo := &mockTransactionRepo{}
	mockWalletService := &mockWalletService{validateErr: errors.ErrSelfTransfer}
	service := service.NewTransactionService(mockTransactionRepo, mockWalletService, &cfg)

	err := service.Send("A", "A", 10)
	if err == nil {
		t.Error("Ожидалась ошибка валидации")
	}
}

func TestSend_RepError(t *testing.T) {
	mockTransactionRepo := &mockTransactionRepo{sendErr: errors.ErrDatabase}
	mockWalletService := &mockWalletService{}
	service := service.NewTransactionService(mockTransactionRepo, mockWalletService, &cfg)

	err := service.Send("A", "B", 10)
	if err == nil {
		t.Error("Ожидалась ошибка репозитория")
	}
}

func TestGetLastTransactions(t *testing.T) {
	mockTransactionRepo := &mockTransactionRepo{
		lastTransactions: []entity.Transaction{
			{From: "A", To: "B", Amount: 10},
		},
	}
	mockWalletService := &mockWalletService{}
	service := service.NewTransactionService(mockTransactionRepo, mockWalletService, &cfg)

	transactions, err := service.GetLastTransactions(1)
	if err != nil || len(transactions) != 1 {
		t.Errorf("Ожидалась одна транзакция. %v", err)
	}
}
