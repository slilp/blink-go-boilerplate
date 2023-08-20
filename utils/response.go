package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpStatusContext struct{
	*gin.Context
}

func (c *HttpStatusContext) OK(data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (c *HttpStatusContext) Created(data interface{}) {
	c.JSON(http.StatusCreated, data)
}

func (c *HttpStatusContext) BadRequest(err string) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err})
}

func (c *HttpStatusContext) NotFound(err string) {
	c.JSON(http.StatusNotFound, gin.H{"error": err})
}

func (c *HttpStatusContext) InternalServerError(err string) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

func (c *HttpStatusContext) NoContent() {
	c.Status(http.StatusNoContent)
}