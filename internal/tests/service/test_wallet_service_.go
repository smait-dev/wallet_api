// Package tests содержит тесты приложения
package tests

import (
	"slices"
	"testing"
	"wallet_api/internal/errors"
	"wallet_api/internal/service"
)

type mockWalletRepository struct {
	addressExists []string
	balances      map[string]float64
}

func (m *mockWalletRepository) CheckWalletAddress(address string) (bool, error) {
	exists := slices.Contains(m.addressExists, address)
	return exists, nil
}

func (m *mockWalletRepository) GetBalance(address string) (float64, error) {
	balance, ok := m.balances[address]
	if !ok {
		return 0, errors.ErrWalletNotFound
	}
	return balance, nil
}

func TestValidateAddresses(t *testing.T) {
	rep := &mockWalletRepository{
		addressExists: []string{"A", "B"},
	}
	service := service.NewWalletService(rep)

	// Одинаковые адреса
	err := service.ValidateAddresses("A", "A")
	if err == nil {
		t.Error("Ожидалась ошибка при переводе самому себе")
	}

	// Неверный отправитель
	err = service.ValidateAddresses("X", "B")
	if err == nil {
		t.Error("Ожидалась ошибка при неверном отправителе")
	}

	// Неверный получатель
	err = service.ValidateAddresses("A", "Y")
	if err == nil {
		t.Error("Ожидалась ошибка при неверном получателе")
	}

	// Позитивный тест
	err = service.ValidateAddresses("A", "B")
	if err != nil {
		t.Errorf("Неожиданная ошибка: %v", err)
	}
}

func TestGetBalance(t *testing.T) {
	rep := &mockWalletRepository{
		addressExists: []string{"A"},
		balances:      map[string]float64{"A": 100.5},
	}
	service := service.NewWalletService(rep)

	// Существующий адрес
	balance, err := service.GetBalance("A")
	if err != nil || balance != 100.5 {
		t.Errorf("Ожидался баланс 100.5, получено %v, ошибка %v", balance, err)
	}

	// Несуществующий адрес
	_, err = service.GetBalance("B")
	if err == nil {
		t.Error("Ожидалась ошибка для несуществующего адреса")
	}
}
