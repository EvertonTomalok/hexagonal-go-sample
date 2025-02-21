package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	port_actor "github.com/EvertonTomalok/ports-challenge/internal/handlers/port_actors"
	"github.com/EvertonTomalok/ports-challenge/internal/ports"
	"github.com/EvertonTomalok/ports-challenge/internal/repositories"
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
	portsRepository := repositories.NewMemDB()
	portsService := ports.NewService(portsRepository)
	portsHandler := port_actor.NewJsonActor(portsService)

	go func() {
		// Handle input json file
		err := portsHandler.HandleUpsertStream(filePath)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Number of registries added to memDB: %d \n", portsRepository.Size())
	}()

	<-shutdown
	fmt.Println("gracefully shutdown...")
}
