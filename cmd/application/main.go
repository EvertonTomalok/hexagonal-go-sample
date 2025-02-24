package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/EvertonTomalok/ports-challenge/internal/adapters/infra"
	"github.com/EvertonTomalok/ports-challenge/internal/adapters/services"
	"github.com/EvertonTomalok/ports-challenge/internal/core/application"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}
	filePath := os.Args[1]

	// Setting up signal handling
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	// Setting up the injection dependencies
	portsRepository := infra.NewMemDB()
	portsService := services.NewService(portsRepository)
	portsHandler := application.NewJsonParser(portsService)

	go func() {
		// Handle input json file
		err := portsHandler.ParseAndUpsertFile(filePath)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Number of registries added to memDB: %d \n", portsRepository.Size())
	}()

	<-shutdown
	fmt.Println("gracefully shutdown...")
}
