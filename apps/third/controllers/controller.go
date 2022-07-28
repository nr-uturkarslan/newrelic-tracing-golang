package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/services"
)

func CreateHandlers(
	router *gin.Engine,
	nrapp *newrelic.Application,
) {

	proxy := router.Group("/third")
	{
		// Health check
		proxy.GET("/health", func(ginctx *gin.Context) {
			ginctx.JSON(http.StatusOK, gin.H{
				"message": "OK!",
			})
		})
	}

	go services.ThirdMethod(context.Background())
}
