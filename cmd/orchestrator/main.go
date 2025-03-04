package main

import (
	"fmt"

	"github.com/dimakirio/calculator-sprint-2/internal"
)

func main() {
	app := application.New()
	fmt.Println("RunServer")
	// app.Run()
	app.RunServer()
}
