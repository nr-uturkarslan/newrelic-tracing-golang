package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/commons"
	dto "github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/dtos"
)

type SecondMethodService struct{}

func (s SecondMethodService) SecondMethod(
	ginctx *gin.Context,
) {

	commons.Log(zerolog.InfoLevel, "Second method is triggered...")

	requestBody, err := s.parseRequestBody(ginctx)

	if err != nil {
		return
	}

	responseDtoFromSecondService, err := s.makeRequestToSecondService(ginctx,
		requestBody)

	if err != nil {
		return
	}

	commons.Log(zerolog.InfoLevel, "Second method is executed.")

	commons.CreateSuccessfulHttpResponse(ginctx, http.StatusOK,
		s.createResponseDto(responseDtoFromSecondService))
}

func (SecondMethodService) parseRequestBody(
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

	commons.Log(zerolog.InfoLevel, "Value provided: "+requestDto.Value)
	commons.Log(zerolog.InfoLevel, "Tag provided: "+requestDto.Tag)

	return &requestDto, nil
}

func (SecondMethodService) makeRequestToSecondService(
	ginctx *gin.Context,
	requestDto *dto.RequestDto,
) (
	*dto.ResponseDto,
	error,
) {

	url := "http://second.second.svc.cluster.local:8080/second"

	requestDtoInBytes, _ := json.Marshal(requestDto)

	httpResponse, err := http.Post(url, "application/json",
		bytes.NewBuffer(requestDtoInBytes))

	if err != nil {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Call to SecondService has failed.")

		return nil, err
	}

	if httpResponse.StatusCode != http.StatusOK {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Call to SecondService has failed.")

		return nil, errors.New("call to second service has failed")
	}

	defer httpResponse.Body.Close()

	responseDtoInBytes, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Response from second service could not be parsed.")

		return nil, err
	}

	var responseDto dto.ResponseDto
	json.Unmarshal(responseDtoInBytes, &responseDto)

	commons.Log(zerolog.InfoLevel, "Value retrieved: "+requestDto.Value)
	commons.Log(zerolog.InfoLevel, "Tag retrieved: "+requestDto.Tag)

	return &responseDto, nil
}

func (SecondMethodService) createResponseDto(
	data *dto.ResponseDto,
) *dto.ResponseDto {
	return &dto.ResponseDto{
		Message: "Succeeded.",
		Value:   data.Value,
		Tag:     data.Tag,
	}
}
