package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/services"
)

type User struct {
	FirstName string
}

func CreateHandlers(
	router *gin.Engine,
) {

	proxy := router.Group("/proxy")
	{
		// Health check
		proxy.GET("/health", func(ginctx *gin.Context) {
			ginctx.JSON(http.StatusOK, gin.H{
				"message": "OK!",
			})
		})

		// First method
		proxy.POST("/method1", services.FirstMethod)
	}
}
