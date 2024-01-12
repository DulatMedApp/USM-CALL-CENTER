// main.go
package main

import "usmcallcenter/main/app/routers"

func main() {
	r := routers.SetupRouter()

	// Запустите сервер
	r.Run(":3000")
}
