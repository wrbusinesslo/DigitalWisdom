package purchase

import (
	"DigitalWisdom/service/internal/database"
	"time"
)

// Purchase is the GORM model for a purchase
type Purchase struct {
	ID           uint           `gorm:"primarykey"`
	UserID       string         `gorm:"column:user_id"`
	PurchaseDate string         `gorm:"column:purchase_date"`
	LocationName string         `gorm:"column:location_name"`
	Items        []PurchaseItem `gorm:"foreignKey:PurchaseID"`
	TotalAmount  float64        `gorm:"column:total_amount"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	DeletedAt    time.Time      `gorm:"column:deleted_at"`
}

// PurchaseItem is the GORM model for a purchased item
type PurchaseItem struct {
	ID         uint    `gorm:"primarykey"`
	PurchaseID uint    `gorm:"column:purchase_id"`
	ProductID  string  `gorm:"column:product_id"`
	Quantity   float64 `gorm:"column:quantity"`
	Unit       int     `gorm:"column:unit"`
	UnitPrice  float64 `gorm:"column:unit_price"`
	SubTotal   float64 `gorm:"column:subtotal"`
}

// Query struct for querying purchases
type Query struct {
	UserID *string `json:"user_id"`
	database.TimeInterval
}

func (Purchase) TableName() string {
	return "user_purchase_log.purchase_records"
}

func (PurchaseItem) TableName() string {
	return "user_purchase_log.purchase_items"
}
