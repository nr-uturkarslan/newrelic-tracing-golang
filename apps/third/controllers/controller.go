package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/services"
)

func CreateHandlers(
	router *gin.Engine,
) {

	proxy := router.Group("/third")
	{
		// Health check
		proxy.GET("/health", func(ginctx *gin.Context) {
			ginctx.JSON(http.StatusOK, gin.H{
				"message": "OK!",
			})
		})

		// Third method
		proxy.POST("/third", services.ThirdMethod)
	}
}
