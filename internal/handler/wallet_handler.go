package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet_api/internal/server"
)

// GetBalance возвращает баланс кошелька по указанному адресу.
func (h *Handler) GetBalance(c *gin.Context) {
	address := c.Param("address")
	balance, err := h.Service.Wallet.GetBalance(address)
	if err != nil {
		server.ErrorResponse(c, err)
		return
	}

	server.Response(c, http.StatusOK, map[string]interface{}{"balance": balance})
}
