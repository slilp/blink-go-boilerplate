package controllers

import (
	"blink-go-gin-boilerplate/middleware"
	"blink-go-gin-boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type Order struct {
	DB *gorm.DB
}

type createOrderRequest struct{
	Products  []models.ProductEntity `form:"products" json:"products" binding:"required"`
}


type updateOrderStatusRequest struct{
	Status  models.OrderStatus `form:"status" json:"status" binding:"required"`
}

// Order by ID godoc
// @summary Order by ID
// @tags order
// @id Order by ID
// @param id path int true "id of order"
// @security BearerAuth
// @router /order/{id} [get]
func (o *Order) FindOne(ctx *gin.Context) {
	order, err := o.FindWithRelationById(ctx)
	user := middleware.ExtractUserToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if order.UserID != user.ID {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found data"})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

// Create order godoc
// @summary Create order
// @tags order
// @id Create order
// @accept json
// @produce json
// @param Order body createOrderRequest true "''"
// @security BearerAuth
// @router /order [post]
func (o *Order) Create(ctx *gin.Context) {
	var form createOrderRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if len(form.Products) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	user := middleware.ExtractUserToken(ctx)
	
	var order models.OrderEntity
	order.UserID = user.ID
	order.Products = form.Products

	if err := o.DB.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var resultOrder models.OrderEntity

	o.DB.Preload("Products").First(&resultOrder,order.ID)

	ctx.JSON(http.StatusCreated, resultOrder)
}

// Update order status godoc
// @summary Update order status
// @tags order
// @id Update order status
// @accept json
// @produce json
// @param id path int true "id of order"
// @param Order body updateOrderStatusRequest true "''"
// @security BearerAuth
// @router /order/order-status/{id} [patch]
func (o *Order) UpdateStatus(ctx *gin.Context) {
	var form updateOrderStatusRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	order , err := o.FindWithRelationById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	order.Status = form.Status

	if err := o.DB.Save(&order).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

// Update order godoc
// @summary Update order
// @tags produorderct
// @id Update order
// @accept json
// @produce json
// @param id path int true "id of order"
// @param Product body createOrderRequest true "''"
// @security BearerAuth
// @router /order/{id} [patch]
func (o *Order) Update(ctx *gin.Context) {
	var form createOrderRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	order , err := o.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	user := middleware.ExtractUserToken(ctx)
	if order.UserID != user.ID {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found data"})
		return
	}

	order.Products = form.Products

	if err := o.DB.Model(&order).Association("Products").Replace(order.Products); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error})
		return
	}

	var resultOrder models.OrderEntity
	o.DB.Preload("Products").First(&resultOrder,order.ID)

	ctx.JSON(http.StatusCreated, resultOrder)
}

// Delete order godoc
// @summary Delete order
// @tags order
// @id Delete order
// @produce plain
// @param id path int true "id of order"
// @security BearerAuth
// @router /order/{id} [delete]
func (o *Order) Delete(ctx *gin.Context) {
	order , err := o.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	user := middleware.ExtractUserToken(ctx)
	if order.UserID != user.ID {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found data"})
		return
	}

	o.DB.Unscoped().Delete(&order)
	ctx.Status(http.StatusNoContent)
}


func (o *Order) FindById(ctx *gin.Context,) (models.OrderEntity,error) {
	id := ctx.Param("id")
	var tempOrder models.OrderEntity
	if error := o.DB.First(&tempOrder,id).Error; error != nil {
		return models.OrderEntity{},error
	}
	return tempOrder , nil
}

func (o *Order) FindWithRelationById(ctx *gin.Context) (models.OrderEntity,error) {
	id := ctx.Param("id")
	var tempOrder models.OrderEntity
	if error := o.DB.Preload("Products").First(&tempOrder,id).Error; error != nil {
		return models.OrderEntity{},error
	}
	return tempOrder , nil
}


