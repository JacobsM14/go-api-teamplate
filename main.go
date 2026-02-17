package main

import (
	"flag"
	"log"
	"os"

	"go-api-template/api"
	db "go-api-template/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	listenAddress := flag.String("port", ":3000", "Port to run the server on")
	flag.Parse()

	store, err := db.NewPostgresStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer store.Close()

	// Initialize the database
	if err := store.Init(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	server := api.NewServer(*listenAddress, store)
	log.Fatal(server.Start())
}
