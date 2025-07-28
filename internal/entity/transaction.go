// Package entity определяет основные доменные сущности приложения.
package entity

// Transaction представляет транзакцию между кошельками.
type Transaction struct {
	From   string  `json:"from"`   // адрес отправителя
	To     string  `json:"to"`     // адрес получателя
	Amount float64 `json:"amount"` // сумма транзакции
}
