package main

import (
	"github.com/gin-gonic/gin"

	controller "github.com/nr-turkarslan/newrelic-tracing-golang/apps/second/controllers"
)

const PORT string = ":8080"

func main() {
	r := gin.Default()
	controller.CreateHandlers(r)
	r.Run(PORT)
}
