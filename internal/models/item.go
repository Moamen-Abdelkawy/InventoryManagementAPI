package models

// Item represents a record in the inventory system.
//
// Fields:
//   - ID:    Unique identifier for each item (UUID string).
//   - Name:  Name of the product.
//   - Stock: Quantity of the product in stock.
//   - Price: Unit price of the product.
type Item struct {
	ID    string  `gorm:"primaryKey" json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}
