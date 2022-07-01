package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/commons"
	dto "github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/dtos"
)

func FirstMethod(
	ginctx *gin.Context,
) {

	log.Info("First method is triggered...")

	requestBody, err := parseRequestBody(ginctx)

	if err != nil {
		return
	}

	responseDtoFromFirstService, err := makeRequestToFirstService(ginctx,
		requestBody)

	if err != nil {
		return
	}

	log.Info("First method is executed.")

	commons.CreateSuccessfulHttpResponse(ginctx, http.StatusOK,
		createResponseDto(responseDtoFromFirstService))
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

func makeRequestToFirstService(
	ginctx *gin.Context,
	requestDto *dto.RequestDto,
) (
	*dto.ResponseDto,
	error,
) {

	url := "http://first.first.svc.cluster.local:8080/first"

	requestDtoInBytes, _ := json.Marshal(requestDto)

	httpResponse, err := http.Post(url, "application/json",
		bytes.NewBuffer(requestDtoInBytes))

	if err != nil {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Call to FirstService has failed.")

		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Call to FirstService has failed.")

		return nil, errors.New("call to first service has failed")
	}

	defer httpResponse.Body.Close()

	responseDtoInBytes, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Response from first service could not be parsed.")

		return nil, err
	}

	var responseDto dto.ResponseDto
	json.Unmarshal(responseDtoInBytes, &responseDto)

	log.Info("Value retrieved: " + responseDto.Value)
	log.Info("Tag retrieved: " + responseDto.Tag)

	return &responseDto, nil
}

func createResponseDto(
	data *dto.ResponseDto,
) *dto.ResponseDto {
	return &dto.ResponseDto{
		Message: "Succeeded.",
		Value:   data.Value,
		Tag:     data.Tag,
	}
}
