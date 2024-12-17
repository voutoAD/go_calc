package main

import (
	"github.com/voutoad/go_calc/internal/application"
)

func main() {
	app := application.NewApplication()
	app.RunServer()
}
