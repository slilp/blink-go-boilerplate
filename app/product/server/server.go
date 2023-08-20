package server

import (
	product "blink-go-gin-boilerplate/app/product/api"
	"blink-go-gin-boilerplate/middleware"
	"blink-go-gin-boilerplate/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func HttpMount(router gin.IRouter, service product.Service) {
	handler := NewHttpHandler(service) 
	productGroup := router.Group("product")
	{
		productGroup.POST("",middleware.AuthorizedUser([]utils.RoleType{"ADMIN"}), handler.create)
		productGroup.PUT("/:id", handler.update)
		productGroup.DELETE("/:id",middleware.AuthorizedUser([]utils.RoleType{"ADMIN"}), handler.delete)
		productGroup.GET("/:id" , handler.findOne)
	}

}

func NewHttpHandler(service product.Service) *handler {
	return &handler{service: service}
}

type handler struct {
	service product.Service
}

// Product by ID godoc
// @summary Product by ID
// @tags product
// @id Product by ID
// @param id path int true "id of product"
// @router /product/{id} [get]
func (h *handler) findOne(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	id := ctx.Param("id")

	product , err := h.service.GetById(id)
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_PRODUCT)
		return
	}
	
	ctx.OK(product)
}

// Create product godoc
// @summary Create product
// @tags product
// @id Create product
// @accept json
// @produce json
// @param Product body product.CreateProductRequest true "''"
// @security BearerAuth
// @router /product [post]
func (h *handler) create(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form product.CreateProductRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}

	var product product.ProductEntity
	copier.Copy(&product, &form)

	if _ ,err := h.service.Create(&product); err != nil {
		ctx.InternalServerError(utils.ERROR_CREATE_PRODUCT)
		return
	}

	ctx.OK(product)

}

// Update product godoc
// @summary Update product
// @tags product
// @id Update product
// @accept json
// @produce json
// @param id path int true "id of product"
// @param Product body product.UpdateProductRequest true "''"
// @security BearerAuth
// @router /product/{id} [put]
func (h *handler) update(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form product.UpdateProductRequest
	if err := c.ShouldBind(&form); err != nil {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}
	id := ctx.Param("id")

	product , err := h.service.GetById(id)
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_PRODUCT)
		return
	}

	copier.Copy(&product, &form)

	if err := h.service.Update(product); err != nil {
		ctx.InternalServerError(utils.ERROR_UPDATE_PRODUCT)
		return
	}

	ctx.OK(product)
}

// Delete product godoc
// @summary Delete product
// @tags product
// @id Delete product
// @produce plain
// @param id path int true "id of product"
// @security BearerAuth
// @router /product/{id} [delete]
func (h *handler) delete(c *gin.Context) {

	ctx := utils.HttpStatusContext{Context: c}

	id := ctx.Param("id")
	product , err := h.service.GetById(id)
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_PRODUCT)
		return
	}

	if err := h.service.Delete(product)  ;  err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ERROR_NOT_FOUND_PRODUCT})
		return
	}

	ctx.NoContent()

}
