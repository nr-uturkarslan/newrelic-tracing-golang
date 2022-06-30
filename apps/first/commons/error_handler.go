package commons

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/first/dto"
)

func CreateSuccessfulHttpResponse(
	c *gin.Context,
	httpStatusCode int,
	responseDto *dto.ResponseDto,
) {
	response, _ := json.Marshal(responseDto)
	c.JSON(httpStatusCode, string(response))
}

func CreateFailedHttpResponse(
	c *gin.Context,
	httpStatusCode int,
	message string,
) {
	log.Error(message)

	responseDto := dto.ResponseDto{
		Message: message,
	}

	response, _ := json.Marshal(responseDto)
	c.JSON(httpStatusCode, response)
}
