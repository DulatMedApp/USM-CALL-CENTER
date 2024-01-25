// D:/GoProjects/UsmCallCenter/main/app/routers/router.go
package routers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"usmcallcenter/main/app/handlers"

	"github.com/gin-gonic/gin"
)

var templates *template.Template

func init() {
	loadTemplates()
}

func loadTemplates() {
	// Путь должен быть относительным к корню проекта
	templatesPath := "./main/app/static/*.html"
	absPath, err := filepath.Abs(templatesPath)
	if err != nil {
		log.Fatal("Error getting absolute path:", err)
	}

	// Подгружаем шаблоны
	templates = template.Must(template.ParseGlob(absPath))
}

// SetupRouter настраивает маршруты для вашего приложения
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middleware для обработки статических файлов
	router.Static("/static", "./main/app/static")

	// Маршрут для отображения HTML-страницы по /income-calls
	router.GET("/income-calls", handlers.IncomeCallsHandler)

	router.GET("/header", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: header.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "header.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/outgoing-calls", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: outgoingCalls.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "outgoingCalls.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/lost-calls", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: lostCalls.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "lostCalls.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/by-employee-calls", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: byEmployeeCalls.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "byEmployeeCalls.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/by-number-calls", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: byNumberCalls.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "byNumberCalls.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/trends", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: trends.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "trends.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/avarage-load", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: avarageLoad.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "avarageLoad.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/flot", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: flot.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "flot.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/chat", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: chat.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "chat.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/tasks", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: tasks.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "tasks.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})

	router.GET("/login", func(c *gin.Context) {
		// Попробуйте выводить имя файла, чтобы убедиться, что оно правильное
		log.Println("Rendering template: login.html")
		// Рендерим шаблон
		err := templates.ExecuteTemplate(c.Writer, "login.html", nil)
		if err != nil {
			log.Println("Error executing template:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}
	})
	// Добавьте другие маршруты, как необходимо

	return router
}
