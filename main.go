package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)

	// TODO: Initialize database
	// TODO: Initialize router
	// TODO: Setup routes
	// TODO: Start server

	log.Println("Server is ready!")
	log.Println("Press Ctrl+C to stop")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
}
