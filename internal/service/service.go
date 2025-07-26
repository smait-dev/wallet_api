// Package service содержит бизнес-логику приложения.
package service


// Service предоставляет доступ к слоям бизнес-логики.
type Service struct {
	Wallet      IWalletService
	Transaction ITransactionService
}

// NewService создает экземпляр Service.
func NewService(rep *repository.Repository) *Service {
	walletService := NewWalletService(rep.Wallet)
	return &Service{
		Wallet:      walletService,
		Transaction: NewTransactionService(rep.Transaction, walletService),
	}
}
