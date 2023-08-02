package controllers

import (
	"blink-go-gin-boilerplate/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type createProductRequest struct {
	Sku   		string      `form:"sku" json:"sku" binding:"required"`
	Name    	string      `form:"name" json:"name" binding:"required"`
	Description	string 		`form:"description" json:"description"`
	Pic			string 		`form:"pic" json:"pic"`
	Price		float64 	`form:"price" json:"price" binding:"numeric,gte=0"`
}

type updateProductRequest struct {
	Sku   		string      `form:"sku" json:"sku"`
	Name    	string      `form:"name" json:"name"`
	Description	string 		`form:"description" json:"description"`
	Pic			string 		`form:"pic" json:"pic"`
	Price		float64 	`form:"price" json:"price" binding:"numeric,gte=0"`
}


type Product struct {
	DB *gorm.DB
}

// Product by ID godoc
// @summary Product by ID
// @tags product
// @id Product by ID
// @param id path int true "id of product"
// @router /product/{id} [get]
func (p *Product) FindOne(ctx *gin.Context) {
	product, err := p.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// Create product godoc
// @summary Create product
// @tags product
// @id Create product
// @accept json
// @produce json
// @param Product body createProductRequest true "''"
// @security BearerAuth
// @router /product [post]
func (p *Product) Create(ctx *gin.Context) {
	var form createProductRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var product models.ProductEntity
	copier.Copy(&product, &form)

	if err := p.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

// Update product godoc
// @summary Update product
// @tags product
// @id Update product
// @accept json
// @produce json
// @param id path int true "id of product"
// @param Product body updateProductRequest true "''"
// @security BearerAuth
// @router /product/{id} [put]
func (p *Product) Update(ctx *gin.Context) {
	var form updateProductRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	product , err := p.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	copier.Copy(&product, &form)

	if err := p.DB.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

// Delete product godoc
// @summary Delete product
// @tags product
// @id Delete product
// @produce plain
// @param id path int true "id of product"
// @security BearerAuth
// @router /product/{id} [delete]
func (p *Product) Delete(ctx *gin.Context) {
	product , err := p.FindById(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not found"})
		return
	}

	p.DB.Unscoped().Delete(&product)
	ctx.Status(http.StatusNoContent)
}


func (p *Product) FindById(ctx *gin.Context) (models.ProductEntity,error) {
	id := ctx.Param("id")
	var tempProduct models.ProductEntity
	if error := p.DB.First(&tempProduct,id).Error; error != nil {
		return models.ProductEntity{},error
	}
	return tempProduct , nil
}

