package main

import (
	"log"

	"github.com/dimakirio/calculator-sprint-2/internal/application"
)

func main() {
	app := application.NewOrchestrator()
	log.Println("Starting Orchestrator on port", app.Config.Addr)
	if err := app.RunServer(); err != nil {
		log.Fatal(err)
	}
}
