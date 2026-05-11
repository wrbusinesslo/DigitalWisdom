package purchase

import (
	purchaseModel "DigitalWisdom/service/dao/daoModels/purchase"
	"time"
)

// PurchaseBody represents the purchase information from the API request
type PurchaseBody struct {
	UserID       string `json:"user_id" valid:"required"`
	PurchaseDate string `json:"purchase_date" valid:"required"`
	LocationName string `json:"location_name"`
	Items        []Item `json:"items" valid:"required"`
}

// Item represents a single item in a purchase
type Item struct {
	ProductID string  `json:"product_id" valid:"required"`
	Quantity  float64 `json:"quantity" valid:"required"`
	Unit      int     `json:"unit" valid:"required"`
	UnitPrice float64 `json:"unit_price" valid:"required"`
}

// CreateArgs arguments for creating a purchase
type CreateArgs struct {
	Purchase PurchaseBody
}

// UpdateArgs arguments for updating a purchase
type UpdateArgs struct {
	IsUpsert bool `valid:"-" form:"is_upsert"`
	Purchase PurchaseBody
}

// GetArgs arguments for getting purchases
type GetArgs struct {
	UserID    string
	StartTime time.Time
	EndTime   time.Time
}

// GetReply defines the reply for getting purchase records.
type GetReply struct {
	Result []*purchaseModel.Purchase
}
