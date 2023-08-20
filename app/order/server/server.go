package server

import (
	order "blink-go-gin-boilerplate/app/order/api"
	"blink-go-gin-boilerplate/middleware"
	"blink-go-gin-boilerplate/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HttpMount(router gin.IRouter, service order.Service) {
	handler := NewHttpHandler(service) 
	orderGroup := router.Group("order")
	{
		orderGroup.POST("",middleware.AuthorizedUser([]utils.RoleType{"USER"}), handler.create)
		orderGroup.PATCH("/order-status/:id",middleware.AuthorizedUser([]utils.RoleType{"ADMIN"}), handler.updateStatus)
		orderGroup.PATCH("/:id",middleware.AuthorizedUser([]utils.RoleType{"USER"}), handler.update)
		orderGroup.DELETE("/:id",middleware.AuthorizedUser([]utils.RoleType{"USER"}), handler.delete)
		orderGroup.GET("/:id",middleware.AuthorizedUser([]utils.RoleType{"USER"}), handler.findOne)
		
	}
}

func NewHttpHandler(service order.Service) *handler {
	return &handler{service: service}
}

type handler struct {
	service order.Service
}

// Order by ID godoc
// @summary Order by ID
// @tags order
// @id Order by ID
// @param id path int true "id of order"
// @security BearerAuth
// @router /order/{id} [get]
func (h *handler) findOne(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	id := ctx.Param("id")

    i, _ := strconv.Atoi(id)

	order , err := h.service.GetByIdWithProduct(uint(i))
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_ORDER)
		return
	}

	ctx.OK(order)
}

// Create order godoc
// @summary Create order
// @tags order
// @id Create order
// @accept json
// @produce json
// @param Order body order.CreateOrderRequest true "''"
// @security BearerAuth
// @router /order [post]
func (h *handler) create(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form order.CreateOrderRequest
	if err := ctx.ShouldBind(&form); err != nil || len(form.Products) == 0 {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}

	user := middleware.ExtractUserToken(c)
	
	var order order.OrderEntity
	order.UserID = user.ID
	order.Products = form.Products

	if _ , err := h.service.Create(&order); err != nil {
		ctx.InternalServerError(utils.ERROR_CREATE_ORDER)
		return
	}

	resultOrder , _ := h.service.GetByIdWithProduct(order.ID)

	ctx.OK(resultOrder)

}

// Update order status godoc
// @summary Update order status
// @tags order
// @id Update order status
// @accept json
// @produce json
// @param id path int true "id of order"
// @param Order body order.UpdateOrderStatusRequest true "''"
// @security BearerAuth
// @router /order/order-status/{id} [patch]
func (h *handler) updateStatus(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form order.UpdateOrderStatusRequest
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}

	id := ctx.Param("id")

	order , err := h.service.GetById(id)
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_ORDER)
		return
	}

	order.Status = form.Status

	if err := h.service.Update(order); err != nil {
		ctx.InternalServerError(utils.ERROR_UPDATE_ORDER)
		return
	}

	ctx.OK(order)
}


// Update order godoc
// @summary Update order
// @tags order
// @id Update order
// @accept json
// @produce json
// @param id path int true "id of order"
// @param Order body order.CreateOrderRequest true "''"
// @security BearerAuth
// @router /order/{id} [patch]
func (h *handler) update(c *gin.Context) {
	ctx := utils.HttpStatusContext{Context: c}

	var form order.CreateOrderRequest
	if err := ctx.ShouldBind(&form); err != nil || len(form.Products) == 0 {
		ctx.BadRequest(utils.ERROR_INVALID_REQUEST)
		return
	}

	id := ctx.Param("id")

	order , err := h.service.GetById(id)
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_ORDER)
		return
	}

	user := middleware.ExtractUserToken(c)
	if order.UserID != user.ID {
		ctx.NotFound(utils.ERROR_NOT_FOUND_ORDER)
		return
	}

	order.Products = form.Products

	if err :=  h.service.UpdateProduct(order); err != nil {
		ctx.InternalServerError(utils.ERROR_UPDATE_ORDER)
		return
	}

	resultOrder , _ := h.service.GetByIdWithProduct(order.ID)

	ctx.OK(resultOrder)
}

// Delete order godoc
// @summary Delete order
// @tags order
// @id Delete order
// @produce plain
// @param id path int true "id of order"
// @security BearerAuth
// @router /order/{id} [delete]
func (h *handler) delete(c *gin.Context) {

	ctx := utils.HttpStatusContext{Context: c}
	id := ctx.Param("id")

	order , err := h.service.GetById(id)
	if err != nil {
		ctx.NotFound(utils.ERROR_NOT_FOUND_ORDER)
		return
	}

	user := middleware.ExtractUserToken(c)
	if order.UserID != user.ID {
		ctx.NotFound(utils.ERROR_NOT_FOUND_ORDER)
		return
	}

	h.service.Delete(order) 

	ctx.NoContent()
}
