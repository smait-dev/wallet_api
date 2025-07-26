// Package handler содержит маршруты и обработчики для HTTP запросов.
package handler

import (
	"wallet_api/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler представляет обработчик HTTP запросов с доступом к слою бизнес-логики.
type Handler struct {
	Service *service.Service
}

// NewHandler возвращает экземпляр обработчика.
func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

// InitRoutes инициализирует маршруты API и возвращает обьект роутера.
func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/send", handler.Send)                         // транзакция между кошельками
		api.GET("/transactions", handler.GetLastTransactions)   // получение последних транзакций
		api.GET("/wallet/:address/balance", handler.GetBalance) // получение баланса кошелька
	}

	return router
}
