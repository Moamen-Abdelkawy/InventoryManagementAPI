package router

import (
	"github.com/Moamen-Abdelkawy/InventoryManagementAPI/internal/handlers"
	"github.com/Moamen-Abdelkawy/InventoryManagementAPI/internal/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRoutes creates and configures the Gin engine with all routes and middleware.
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Apply rate limiting to every request
	r.Use(middleware.RateLimitMiddleware)

	// Inventory routes
	r.GET("/inventory", handlers.GetAllItems)
	r.GET("/inventory/:id", handlers.GetSingleItem)
	r.POST("/inventory", handlers.CreateItem)
	r.PUT("/inventory/:id", handlers.UpdateItem)
	r.DELETE("/inventory/:id", handlers.DeleteItem)

	return r
}
