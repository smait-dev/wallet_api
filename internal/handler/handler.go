// Package handler содержит маршруты и обработчики для HTTP запросов.
package handler

import (
	"github.com/gin-gonic/gin"
	"wallet_api/internal/config"
	"wallet_api/internal/service"
)

// Handler представляет обработчик HTTP запросов с доступом к слою бизнес-логики.
type Handler struct {
	Service *service.Service
	Config  *config.Config
}

// NewHandler возвращает экземпляр обработчика.
func NewHandler(service *service.Service, cfg *config.Config) *Handler {
	return &Handler{Service: service, Config: cfg}
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

func (handler *Handler) GetConfig() *config.Config {
	return handler.Config
}
