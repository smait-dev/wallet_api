package entity

// Wallet представляет кошелек пользователя.
type Wallet struct {
	Address string  `json:"address"` // адрес кошелька
	Balance float64 `json:"balance"` // текущий баланс
}
