package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/<YourUsername>/InventoryManagementAPI/internal/database"
	"github.com/<YourUsername>/InventoryManagementAPI/internal/models"
)

// GetAllItems handles GET /inventory with optional pagination, filtering, and sorting.
func GetAllItems(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	nameFilter := c.Query("name")
	minStock, _ := strconv.Atoi(c.DefaultQuery("min_stock", "0"))

	sortField := c.DefaultQuery("sort", "id")
	sortOrder := c.DefaultQuery("order", "asc")

	query := database.DB.Model(&models.Item{})

	// Filtering by name
	if nameFilter != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(nameFilter)+"%")
	}

	// Filtering by minimum stock
	if minStock > 0 {
		query = query.Where("stock >= ?", minStock)
	}

	// Sorting
	if sortField != "" {
		if sortOrder == "desc" {
			sortField += " desc"
		}
		query = query.Order(sortField)
	}

	// Count total items
	var total int64
	query.Count(&total)

	// Retrieve paginated items
	var items []models.Item
	if err := query.Offset(offset).Limit(limit).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  items,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}

// GetSingleItem handles GET /inventory/:id to retrieve one item by UUID.
func GetSingleItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	c.JSON(http.StatusOK, item)
}

// CreateItem handles POST /inventory to add a new item.
func CreateItem(c *gin.Context) {
	var input struct {
		Name  string  `json:"name"`
		Stock int     `json:"stock"`
		Price float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	newItem := models.Item{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Stock: input.Stock,
		Price: input.Price,
	}

	if err := database.DB.Create(&newItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

// UpdateItem handles PUT /inventory/:id to modify an existing item.
func UpdateItem(c *gin.Context) {
	id := c.Param("id")
	var existing models.Item

	if err := database.DB.First(&existing, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var input struct {
		Name  *string  `json:"name"`
		Stock *int     `json:"stock"`
		Price *float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if input.Name != nil {
		existing.Name = *input.Name
	}
	if input.Stock != nil {
		existing.Stock = *input.Stock
	}
	if input.Price != nil {
		existing.Price = *input.Price
	}

	if err := database.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}
	c.JSON(http.StatusOK, existing)
}

// DeleteItem handles DELETE /inventory/:id to remove an item by UUID.
func DeleteItem(c *gin.Context) {
	id := c.Param("id")
	var item models.Item

	if err := database.DB.First(&item, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := database.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item successfully deleted"})
}
