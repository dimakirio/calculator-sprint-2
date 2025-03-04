package main

import (
	"log"

	"github.com/dimakirio/calculator-sprint-2/internal/application"
)

func main() {
	agent := application.NewAgent()
	log.Println("Starting Agent...")
	agent.Run()
}
