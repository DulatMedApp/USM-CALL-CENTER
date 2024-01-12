// В handlers/income_calls.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GETIncomeCallsHandler обработчик для страницы income-calls
func GETIncomeCallsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "income-calls.html", nil)
}
