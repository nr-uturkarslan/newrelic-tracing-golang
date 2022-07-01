package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/services"
)

func CreateHandlers(
	router *gin.Engine,
	nrapp *newrelic.Application,
) {

	router.Use(nrgin.Middleware(nrapp))

	firstMethodService := services.FirstMethodService{}
	secondMethodService := services.SecondMethodService{}

	proxy := router.Group("/proxy")
	{
		// Health check
		proxy.GET("/health", func(ginctx *gin.Context) {
			ginctx.JSON(http.StatusOK, gin.H{
				"message": "OK!",
			})
		})

		// First method
		proxy.POST("/method1", firstMethodService.FirstMethod)

		// Second method
		proxy.POST("/method2", secondMethodService.SecondMethod)
	}
}
