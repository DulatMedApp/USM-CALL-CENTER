// В handlers/income_calls.go

package handlers

// GETIncomeCallsHandler обработчик для страницы income-calls

import (
	"html/template"
	"log"
	"net/http"
	"usmcallcenter/main/app/db"
	"usmcallcenter/main/app/models"

	"github.com/gin-gonic/gin"
)

func IncomeCallsHandler(c *gin.Context) {
	// Получение экземпляра базы данных
	dbInstance, err := db.InitDB()
	if err != nil {
		log.Println("Error connecting to the database:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer dbInstance.Close()

	rows, err := dbInstance.Query("SELECT DATE_FORMAT(`calldate`, '%H:%i %d.%m.%Y') AS `formatted_calldate`, `src`, `dst`, TIME_FORMAT(SEC_TO_TIME(`billsec`), '%i:%s') AS `call_duration_formatted`, `disposition`, TIME_FORMAT(SEC_TO_TIME(IFNULL(TIMESTAMPDIFF(SECOND, LAG(`calldate`) OVER (PARTITION BY `src` ORDER BY `calldate`), `calldate`), 0)), '%i:%s') AS `wait_time` FROM `asteriskcdrdb`.`cdr` WHERE LENGTH(`src`) > 3 AND LENGTH(`dst`) >= 3 AND `disposition` <> 'FAILED' ORDER BY `calldate` DESC;")
	if err != nil {
		log.Println("Error executing database query:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer rows.Close()

	// Обработка результатов запроса
	var callData []models.CallInfo
	for rows.Next() {
		var calldate, src, dst, billsec, disposition, wait_time string
		err := rows.Scan(&calldate, &src, &dst, &billsec, &disposition, &wait_time)
		if err != nil {
			log.Println("Error scanning database rows:", err)
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		callData = append(callData, models.CallInfo{
			CallerNumber: src,
			Duration:     billsec,
			CallTime:     calldate,
			CallStatus:   disposition,
			WaitTime:     wait_time,
			// Добавьте другие поля по необходимости
		})
	}

	// Рендеринг HTML-шаблона с данными из базы данных
	tmpl, err := template.ParseFiles("main/app/static/incomeCalls.html", "main/app/static/slidebar.html", "main/app/static/header.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmpl.Execute(c.Writer, gin.H{
		"CallData": callData,
	})
	if err != nil {
		log.Println("Error executing template:", err)
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}
