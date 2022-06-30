package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/nr-turkarslan/newrelic-tracing-golang/apps/first/services"
)

type User struct {
	FirstName string
}

func CreateHandlers(
	r *gin.Engine,
) {
	createHealthHandler(r)
	createFirstMethodHandler(r)
}

// Health check
func createHealthHandler(
	r *gin.Engine,
) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	})
}

// First method
func createFirstMethodHandler(
	r *gin.Engine,
) {
	r.POST("/method1", services.FirstMethod)
}
