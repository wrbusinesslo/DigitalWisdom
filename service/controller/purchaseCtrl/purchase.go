package purchaseCtrl

import (
	purchaseModel "DigitalWisdom/service/dao/daoModels/purchase"
	"DigitalWisdom/service/dao/postgres/purchase"
	"DigitalWisdom/service/internal/database"
	bo "DigitalWisdom/service/internal/model/bo/purchase"
	"context"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

// PurchaseCtrl defines the interface for purchase controller
type PurchaseCtrl interface {
	Create(ctx context.Context, args *bo.CreateArgs) error
	Update(ctx context.Context, args *bo.UpdateArgs) error
	Get(ctx context.Context, args *bo.GetArgs) (*bo.GetReply, error)
}

// NewPurchaseCtrl creates a new PurchaseCtrl
func NewPurchaseCtrl(pack purchaseCtrlPack) PurchaseCtrl {
	return &purchaseCtrl{
		pack: pack,
	}
}

type purchaseCtrlPack struct {
	dig.In
	PostgresDigitalWisdom *gorm.DB `name:"postgres_digitalWisdom"`
}

type purchaseCtrl struct {
	pack purchaseCtrlPack
}

func (c *purchaseCtrl) Create(ctx context.Context, args *bo.CreateArgs) error {

	dao := purchase.NewPurchaseDAO(c.pack.PostgresDigitalWisdom)

	// Convert bo.Purchase to dao.Purchase
	daoPurchase := &purchaseModel.Purchase{
		UserID:       args.Purchase.UserID,
		PurchaseDate: args.Purchase.PurchaseDate,
		LocationName: args.Purchase.LocationName,
	}

	var totalAmount float64
	for _, item := range args.Purchase.Items {
		daoPurchase.Items = append(daoPurchase.Items, purchaseModel.PurchaseItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Unit:      item.Unit,
			UnitPrice: item.UnitPrice,
			SubTotal:  item.Quantity * item.UnitPrice,
		})
		totalAmount += item.Quantity * item.UnitPrice
	}
	daoPurchase.TotalAmount = totalAmount

	err := dao.CreatePurchase(daoPurchase)
	if err != nil {
		return err
	}

	return nil
}

func (c *purchaseCtrl) Update(ctx context.Context, args *bo.UpdateArgs) error {
	// You need to define how to get the purchase to update.
	// This is a placeholder.
	// For example, you might pass an ID and the data to update.
	// purchase := &purchaseModel.Purchase{ ... }
	// return c.purchaseDAO.UpdatePurchase(purchase)
	return nil
}

func (c *purchaseCtrl) Get(ctx context.Context, args *bo.GetArgs) (*bo.GetReply, error) {
	dao := purchase.NewPurchaseDAO(c.pack.PostgresDigitalWisdom)

	query := purchaseModel.Query{
		UserID: &args.UserID,
		TimeInterval: database.TimeInterval{
			StartTime: args.StartTime,
			EndTime:   args.EndTime,
		},
	}

	purchases, err := dao.GetPurchases(&query)
	if err != nil {
		return nil, err
	}

	if len(purchases) == 0 {
		// Return an empty response instead of an error if no records are found
		return &bo.GetReply{}, nil
	}

	result := &bo.GetReply{
		Result: purchases,
	}

	return result, nil
}
