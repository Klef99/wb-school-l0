package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/klef99/wb-school-l0/cmd/di"
)

func main() {
	_, shutdown, err := di.InitializeDependencies()
	if err != nil {
		log.Fatalf("Error initializing dependencies: %w", err)
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)
	defer signal.Stop(quit)
	<-quit
	shutdown()

	return
}
