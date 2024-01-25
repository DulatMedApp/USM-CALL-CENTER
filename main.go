// main.go
package main

import (
	"log"
	"usmcallcenter/main/app/db"
	"usmcallcenter/main/app/routers"
)

func main() {
	r := routers.SetupRouter()

	db, err := db.InitDB()
	if err != nil {
		log.Fatal("Error in DB", err)
	}

	defer db.Close()

	// Запустите сервер
	r.Run(":3000")

}
