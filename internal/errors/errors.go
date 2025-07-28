// Package errors содержит вспомогательные данные и функционал.
package errors

import (
	"errors"
	"net/http"
)

// Шаблоны сообщений об ошибках.
var (
	ErrServerRun          = errors.New("ошибка запуска сервера")
	ErrInvalidRequestData = errors.New("неверные данные запроса")
	ErrInvalidCountParam  = errors.New("неверный параметр count")
	ErrInvalidAmount      = errors.New("указана некорректная сумма")

	ErrSelfTransfer    = errors.New("отправка себе запрещена")
	ErrInvalidSender   = errors.New("неверный адрес отправителя")
	ErrInvalidReceiver = errors.New("неверный адрес получателя")
	ErrWalletNotFound  = errors.New("кошелек не найден")

	ErrInsufficientFunds = errors.New("недостаточно средств")
	ErrDebitFailed       = errors.New("ошибка списания средств")
	ErrCreditFailed      = errors.New("ошибка зачисления средств")
	ErrTransactionFailed = errors.New("ошибка выполнения транзакции")

	ErrDatabase       = errors.New("ошибка базы данных")
	ErrDatabaseDriver = errors.New("ошибка драйвера базы данных")
	ErrMigrations     = errors.New("ошибка выполнения миграций")
	ErrSeed           = errors.New("не удалось добавить записи кошельков")
)

var ErrorStatusMap = map[error]int{
	ErrSelfTransfer:       http.StatusBadRequest,
	ErrInvalidSender:      http.StatusBadRequest,
	ErrInvalidReceiver:    http.StatusBadRequest,
	ErrInvalidAmount:      http.StatusBadRequest,
	ErrInsufficientFunds:  http.StatusBadRequest,
	ErrInvalidRequestData: http.StatusBadRequest,
	ErrWalletNotFound:     http.StatusNotFound,
	ErrDatabase:           http.StatusInternalServerError,
	ErrDebitFailed:        http.StatusInternalServerError,
	ErrCreditFailed:       http.StatusInternalServerError,
	ErrTransactionFailed:  http.StatusInternalServerError,
}
