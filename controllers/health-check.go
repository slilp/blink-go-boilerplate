package controllers

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @summary Health Check
// @description Health checking for the service
// @tags health-check
// @id HealthCheck
// @produce plain
// @router /health-check [get]
func  HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "health check",
	})
}