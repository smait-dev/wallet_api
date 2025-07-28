package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wallet_api/internal/entity"
	"wallet_api/internal/errors"
	"wallet_api/internal/server"
)

// Send обрабатывает отправку средств между кошельками.
func (h *Handler) Send(c *gin.Context) {
	var request entity.Transaction
	if err := c.BindJSON(&request); err != nil {
		server.ErrorResponse(c, errors.ErrInvalidRequestData)
		return
	}

	err := h.Service.Transaction.Send(request.From, request.To, request.Amount)
	if err != nil {
		server.ErrorResponse(c, err)
		return
	}

	server.Response(c, http.StatusOK, map[string]interface{}{"status": "success"})
}

// GetLastTransactions возвращает последние транзакции используя параметр запроса count.
// При слишком большом значении count вернет только максимально допустимое количество записей.
func (h *Handler) GetLastTransactions(c *gin.Context) {
	count, err := strconv.Atoi(c.Query("count"))
	if err != nil {
		server.ErrorResponse(c, errors.ErrInvalidRequestData)
		return
	} else if count <= 0 {
		server.Response(c, http.StatusOK, gin.H{})
		return
	}

	transactions, err := h.Service.Transaction.GetLastTransactions(count)
	if err != nil {
		server.ErrorResponse(c, err)
		return
	} else if transactions == nil {
		server.Response(c, http.StatusOK, gin.H{})
		return
	}

	server.Response(c, http.StatusOK, transactions)
}
