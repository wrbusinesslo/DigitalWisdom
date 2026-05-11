package purchase

import (
	purchaseModel "DigitalWisdom/service/dao/daoModels/purchase"
	"DigitalWisdom/service/internal/utils"
	"gorm.io/gorm"
	"time"
)

type PurchaseDAO struct {
	db *gorm.DB
}

// PurchaseCtrl defines the interface for purchase controller
type PurchaseDao interface {
	CreatePurchase(args *purchaseModel.Purchase) error
	UpdatePurchase(purchase *purchaseModel.Purchase) error
	GetPurchases(query *purchaseModel.Query) ([]*purchaseModel.Purchase, error)
}

func NewPurchaseDAO(db *gorm.DB) *PurchaseDAO {
	return &PurchaseDAO{db: db}
}

// CreatePurchase creates a new purchase record in the database using a transaction
func (dao *PurchaseDAO) CreatePurchase(args *purchaseModel.Purchase) error {
	tx := dao.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 建立主檔 (purchase_records)
	queryPurchase := &purchaseModel.Purchase{
		UserID:       args.UserID,
		PurchaseDate: args.PurchaseDate,
		LocationName: args.LocationName,
		TotalAmount:  args.TotalAmount,
		CreatedAt:    time.Now().In(utils.Loc),
	}

	if err := tx.Create(queryPurchase).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 準備批次建立的 purchase_items
	var queryItems []purchaseModel.PurchaseItem
	for _, item := range args.Items {
		item.PurchaseID = queryPurchase.ID
		queryItems = append(queryItems, item)
	}

	// 一次性建立所有明細
	if err := tx.Create(&queryItems).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交交易
	return tx.Commit().Error
}

// GetPurchases retrieves purchases from the database based on a query
func (dao *PurchaseDAO) GetPurchases(query *purchaseModel.Query) ([]*purchaseModel.Purchase, error) {
	var purchases []*purchaseModel.Purchase
	db := dao.db.Model(&purchaseModel.Purchase{})

	if query.UserID != nil {
		db = db.Where("user_id = ?", *query.UserID)
	}
	if !query.StartTime.IsZero() {
		db = db.Where("purchase_date >= ?", query.StartTime)
	}
	if !query.EndTime.IsZero() {
		db = db.Where("purchase_date <= ?", query.EndTime)
	}

	err := db.Preload("Items").Find(&purchases).Error
	return purchases, err
}

// UpdatePurchase updates a purchase in the database
func (dao *PurchaseDAO) UpdatePurchase(purchase *purchaseModel.Purchase) error {
	// This will update the purchase and its associated items
	return dao.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(purchase).Error
}
