// В handlers/header

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GETIncomeCallsHandler обработчик для страницы header
func GETHeaderHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "header.html", nil)
}
