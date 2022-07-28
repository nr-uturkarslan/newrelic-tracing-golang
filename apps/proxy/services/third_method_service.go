package services

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/commons"
	dto "github.com/nr-turkarslan/newrelic-tracing-golang/apps/proxy/dtos"
)

type ThirdMethodService struct {
	Nrapp     *newrelic.Application
	KafkaConn *kafka.Conn
}

func (s ThirdMethodService) ThirdMethod(
	ginctx *gin.Context,
) {

	commons.LogWithContext(ginctx, zerolog.InfoLevel, "Third method is triggered...")

	requestBody, err := s.parseRequestBody(ginctx)

	if err != nil {
		return
	}

	responseDtoFromThirdService, err := s.publishToKafka(ginctx,
		requestBody)

	if err != nil {
		return
	}

	commons.LogWithContext(ginctx, zerolog.InfoLevel, "Third method is executed.")

	commons.CreateSuccessfulHttpResponse(ginctx, http.StatusOK,
		s.createResponseDto(responseDtoFromThirdService))
}

func (*ThirdMethodService) parseRequestBody(
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

	commons.LogWithContext(ginctx, zerolog.InfoLevel, "Value provided: "+requestDto.Value)
	commons.LogWithContext(ginctx, zerolog.InfoLevel, "Tag provided: "+requestDto.Tag)

	return &requestDto, nil
}

func (s *ThirdMethodService) publishToKafka(
	ginctx *gin.Context,
	requestDto *dto.RequestDto,
) (
	*dto.ResponseDto,
	error,
) {

	commons.LogWithContext(ginctx, zerolog.InfoLevel, "Publishing Kafka message...")
	requestDtoInBytes, _ := json.Marshal(requestDto)

	_, err := s.KafkaConn.WriteMessages(kafka.Message{
		Value: requestDtoInBytes,
	})

	if err != nil {
		commons.CreateFailedHttpResponse(ginctx, http.StatusBadRequest,
			"Message could not be published.")

		return nil, err
	}

	commons.LogWithContext(ginctx, zerolog.InfoLevel, "Kafka message is published.")

	responseDto := dto.ResponseDto{
		Message: "Message is published.",
	}

	return &responseDto, nil
}

func (*ThirdMethodService) createResponseDto(
	data *dto.ResponseDto,
) *dto.ResponseDto {
	return &dto.ResponseDto{
		Message: "Succeeded.",
		Value:   data.Value,
		Tag:     data.Tag,
	}
}
