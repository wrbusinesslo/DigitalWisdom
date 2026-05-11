package DigitalWisdomApi

import (
	"DigitalWisdom/service/controller/purchaseCtrl"
	boPurchase "DigitalWisdom/service/internal/model/bo/purchase"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"net/http"
	"time"
)

func NewPurchase(pack purchaseApiPack) {
	c := &purchaseApi{pack: pack}
	group := pack.Root.Group("purchase")
	{
		group.POST("create", c.createPurchase)
		group.POST("update", c.updatePurchase)
		group.POST("get", c.getPurchase)
	}
}

type purchaseApiPack struct {
	dig.In
	PurchaseCtrl purchaseCtrl.PurchaseCtrl
	Root         *gin.RouterGroup
}

type purchaseApi struct {
	pack purchaseApiPack
}

func (api *purchaseApi) createPurchase(ctx *gin.Context) {
	var body boPurchase.PurchaseBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	args := &boPurchase.CreateArgs{
		Purchase: body,
	}

	err := api.pack.PurchaseCtrl.Create(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (api *purchaseApi) updatePurchase(ctx *gin.Context) {
	var body boPurchase.PurchaseBody

	if err := ctx.BindQuery(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid query params"})
		return
	}

	if _, err := govalidator.ValidateStruct(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// You need to define what the body of the update request looks like.
	// var models []*purchaseModel.Purchase
	// if err := ctx.BindJSON(&models); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
	// 	return
	// }

	query := &boPurchase.UpdateArgs{}

	err := api.pack.PurchaseCtrl.Update(ctx, query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (api *purchaseApi) getPurchase(ctx *gin.Context) {
	var form struct {
		UserID    string `json:"user_id" valid:"required"`
		StartDate string `json:"start_date" valid:"required"`
		EndDate   string `json:"end_date" valid:"required"`
	}
	if err := ctx.BindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})
		return
	}

	layout := "2006-01-02 15:03:04"

	startDate, _ := time.Parse(layout, form.StartDate)
	endDate, _ := time.Parse(layout, form.EndDate)

	args := &boPurchase.GetArgs{
		UserID:    form.UserID,
		StartTime: startDate,
		EndTime:   endDate,
	}

	result, err := api.pack.PurchaseCtrl.Get(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
