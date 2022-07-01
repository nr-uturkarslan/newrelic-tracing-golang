package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/nr-turkarslan/newrelic-tracing-golang/apps/second/services"
)

func CreateHandlers(
	router *gin.Engine,
) {

	proxy := router.Group("/second")
	{
		// Health check
		proxy.GET("/health", func(ginctx *gin.Context) {
			ginctx.JSON(http.StatusOK, gin.H{
				"message": "OK!",
			})
		})

		// Second method
		proxy.POST("/second", services.SecondMethod)
	}
}
