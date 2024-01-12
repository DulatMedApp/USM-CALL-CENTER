package routers

import (
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/income-calls", handlers.getIncomeCallsHandler)

	return router
}
