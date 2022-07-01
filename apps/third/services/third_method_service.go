package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/commons"
	dto "github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/dtos"
)

func ThirdMethod(
	ginctx *gin.Context,
) {

	log.Info("Third method is triggered...")

	requestBody, err := parseRequestBody(ginctx)

	if err != nil {
		return
	}

	log.Info("Third method is executed.")

	commons.CreateSuccessfulHttpResponse(ginctx, http.StatusOK,
		createResponseDto(requestBody))
}

func parseRequestBody(
	ginctx *gin.Context,
) (
	*dto.RequestDto,
	error,
) {
	var requestDto dto.RequestDto

	err := ginctx.BindJSON(&requestDto)

	if err != nil {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Request body could not be parsed.")

		return nil, err
	}

	log.Info("Value provided: " + requestDto.Value)
	log.Info("Tag provided: " + requestDto.Tag)

	return &requestDto, nil
}

func createResponseDto(
	data *dto.RequestDto,
) *dto.ResponseDto {
	return &dto.ResponseDto{
		Message: "Succeeded.",
		Value:   data.Value,
		Tag:     data.Tag,
	}
}
