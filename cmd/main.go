package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/macreai/todo-app-clean-architecture/internal/http"
	"github.com/macreai/todo-app-clean-architecture/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	router := http.NewRouter(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Listen(":" + port))
}
