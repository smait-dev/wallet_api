// Package service содержит бизнес-логику приложения.
package service

import (
	"wallet_api/internal/config"
	"wallet_api/internal/repository"
)

// Service предоставляет доступ к слоям бизнес-логики.
type Service struct {
	Wallet      IWalletService
	Transaction ITransactionService
}

// NewService создает экземпляр Service.
func NewService(rep *repository.Repository, cfg *config.Config) *Service {
	walletService := NewWalletService(rep.Wallet)
	return &Service{
		Wallet:      walletService,
		Transaction: NewTransactionService(rep.Transaction, walletService, cfg),
	}
}
