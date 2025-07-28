package server

import (
	"wallet_api/internal/errors"
	"github.com/gin-gonic/gin"
)

// ErrorResponse формирование ответа сервера при возникновении ошибки.
func ErrorResponse(c *gin.Context, err error) {
	c.JSON(errors.ErrorStatusMap[err], map[string]interface{}{
		"error": err.Error(),
	})
}

// Error формирование общего ответа сервера.
func Response(c *gin.Context, code int, data any) {
	c.JSON(code, data)
}
