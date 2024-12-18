package main

import (
	"github.com/voutoad/go_calc/internal/application"
)

func main() {
	app := application.NewApplication()
	//TIP Запуск сервиса в консольном режиме
	//app.Run()

	//TIP хапуск сервиса в серверном режиме
	app.RunServer()
}
