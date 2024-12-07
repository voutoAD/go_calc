package application

import (
	"bufio"
	calc "github.com/voutoad/go_calc/pkg/go_calc"
	"log"
	"os"
	"strings"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Run() error {
	for {
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read from console")
		}
		text = strings.TrimSpace(text)
		if text == "exit" {
			log.Println("application stopped")
			return nil
		}
		result, err := calc.Calc(text)
		if err != nil {
			log.Printf("%s calculation failed with error: %s\n", text, err)
		} else {
			log.Printf("%s=%f\n", text, result)
		}
	}
}
