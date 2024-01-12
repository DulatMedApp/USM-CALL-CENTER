package routers

import (
	handlers "command-line-argumentsD:\\GoProjects\\Call Center\\main\\app\\handlers\\income-calls.go"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/income-calls", handlers.getIncomeCallsHandler)

	return router
}
