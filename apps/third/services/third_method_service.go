package services

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/commons"
	dto "github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/dtos"
)

func ThirdMethod(
	ctx context.Context,
) {

	log.Info("Third method is triggered...")

	commons.LogWithContext(zerolog.InfoLevel, "Starting Kafka...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka.kafka.svc.cluster.local:9092"},
		Topic:   "tracing",
		GroupID: "tracingconsumer",
	})

	commons.LogWithContext(zerolog.InfoLevel, "Kafka is started.")

	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			commons.Log(zerolog.ErrorLevel, "Kafka message could not be received.")
			return
		}

		commons.Log(zerolog.InfoLevel, "Kafka message is received: "+string(msg.Value))
	}
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
