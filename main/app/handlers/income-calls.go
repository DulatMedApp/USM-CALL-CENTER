package handlers

import (
	"github.com/gin-gonic/gin"
)

func getIncomeCallsHandler(c *gin.Context) {
	// Отправьте HTML-страницу custom-tables.html
	c.File("custom-tables.html")
}
