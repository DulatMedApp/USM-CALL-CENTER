// В routers/router.go

package routers

import (
	"html/template"
	"log"
	"path/filepath"
	"usmcallcenter/main/app/handlers"

	"github.com/gin-gonic/gin"
)

var templates *template.Template

func init() {
	loadTemplates()
}

func loadTemplates() {
	templatesPath := "./main/app/static/*.html"
	absPath, err := filepath.Abs(templatesPath)
	if err != nil {
		log.Fatal("Error getting absolute path:", err)
	}

	templates = template.Must(template.ParseGlob(absPath))
}

// SetupRouter настраивает маршруты для вашего приложения
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware для обработки статических файлов
	router.Static("/static", "./main/app/static")

	// Маршрут для отображения HTML-страницы по /income-calls
	router.GET("/income-calls", handlers.GETIncomeCallsHandler)

	// Добавьте другие маршруты, как необходимо

	return router
}
