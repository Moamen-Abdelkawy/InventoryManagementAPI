package main

import (
	"log"

	"github.com/Moamen-Abdelkawy/InventoryManagementAPI/internal/database"
	"github.com/Moamen-Abdelkawy/InventoryManagementAPI/internal/router"
)

// main is the entry point of our application. It initializes the database
// and then starts the Gin server with defined routes.
func main() {
	// 1. Initialize the database and run migrations/seed data
	database.InitDatabase()

	// 2. Setup the Gin engine with routes and middlewares
	r := router.SetupRoutes()

	// 3. Start the server on port 8080
	log.Println("Starting server on http://localhost:8080 ...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
