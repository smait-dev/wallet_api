package handler

import (
	"strconv"
	"wallet_api/internal/entity"
	"github.com/gin-gonic/gin"
)

// Send обрабатывает отправку средств между кошельками.
func (h *Handler) Send(c *gin.Context) {
	var request entity.Transaction
	if err := c.BindJSON(&request); err != nil {
		// ответ
		return
	}

	err := h.Service.Transaction.Send(request.From, request.To, request.Amount)
	if err != nil {
		// ответ
		return
	}

	// ответ
}

// GetLastTransactions возвращает последние транзакции используя параметр запроса count.
func (h *Handler) GetLastTransactions(c *gin.Context) {
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		// ответ
		return
	} else if count <= 0 {
		// ответ
		return
	}

	transactions, err := h.Service.Transaction.GetLastTransactions(count)
	if err != nil {
		// ответ
		return
	} else if transactions == nil {
		// ответ
		return
	}

	// ответ
}
