package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Добавьте маршруты
	r.GET("/custom-tables", func(c *gin.Context) {
		c.File("/public/custom-tables.html")
	})

	// Запустите сервер
	r.Run(":8080")
}
