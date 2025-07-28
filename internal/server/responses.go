package server

import (
	"github.com/gin-gonic/gin"
	"wallet_api/internal/errors"
)

// ErrorResponse формирование ответа сервера при возникновении ошибки.
func ErrorResponse(c *gin.Context, err error) {
	c.JSON(errors.ErrorStatusMap[err], map[string]interface{}{
		"error": err.Error(),
	})
}

// Response формирование общего ответа сервера.
func Response(c *gin.Context, code int, data any) {
	c.JSON(code, data)
}
