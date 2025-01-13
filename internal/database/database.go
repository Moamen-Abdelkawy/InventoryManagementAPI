package database

import (
	"log"
	"os"

	"github.com/Moamen-Abdelkawy/InventoryManagementAPI/internal/models"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection handle.
var DB *gorm.DB

// InitDatabase connects to PostgreSQL, runs migrations, and seeds initial data.
func InitDatabase() {
	var err error

	// Fetch DSN from environment or use a fallback
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable"
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the PostgreSQL database:", err)
	}

	// Auto-migrate the Item model (creates or updates the schema)
	if err := DB.AutoMigrate(&models.Item{}); err != nil {
		log.Fatal("Auto-migration failed:", err)
	}

	seedDatabase()
}

// seedDatabase inserts sample data if the database is empty.
func seedDatabase() {
	var count int64
	DB.Model(&models.Item{}).Count(&count)

	if count == 0 {
		items := []models.Item{
			{ID: uuid.New().String(), Name: "Laptop", Stock: 10, Price: 999.99},
			{ID: uuid.New().String(), Name: "Smartphone", Stock: 20, Price: 699.99},
			{ID: uuid.New().String(), Name: "Headphones", Stock: 15, Price: 199.99},
			{ID: uuid.New().String(), Name: "Keyboard", Stock: 25, Price: 89.99},
			{ID: uuid.New().String(), Name: "Mouse", Stock: 30, Price: 49.99},
			{ID: uuid.New().String(), Name: "Monitor", Stock: 12, Price: 299.99},
			{ID: uuid.New().String(), Name: "Webcam", Stock: 18, Price: 79.99},
			{ID: uuid.New().String(), Name: "Printer", Stock: 7, Price: 149.99},
			{ID: uuid.New().String(), Name: "Tablet", Stock: 5, Price: 399.99},
			{ID: uuid.New().String(), Name: "Smartwatch", Stock: 14, Price: 249.99},
			{ID: uuid.New().String(), Name: "External Hard Drive", Stock: 8, Price: 119.99},
			{ID: uuid.New().String(), Name: "USB Flash Drive", Stock: 50, Price: 19.99},
			{ID: uuid.New().String(), Name: "Router", Stock: 6, Price: 89.99},
			{ID: uuid.New().String(), Name: "Projector", Stock: 3, Price: 499.99},
			{ID: uuid.New().String(), Name: "Bluetooth Speaker", Stock: 22, Price: 129.99},
			{ID: uuid.New().String(), Name: "Gaming Console", Stock: 11, Price: 499.99},
			{ID: uuid.New().String(), Name: "Camera", Stock: 4, Price: 599.99},
			{ID: uuid.New().String(), Name: "Fitness Tracker", Stock: 16, Price: 99.99},
			{ID: uuid.New().String(), Name: "Drone", Stock: 2, Price: 899.99},
			{ID: uuid.New().String(), Name: "VR Headset", Stock: 9, Price: 399.99},
		}
		DB.Create(&items)
		log.Println("Database seeded with 20 sample items.")
	} else {
		log.Println("Database already contains items; skipping seeding.")
	}
}
