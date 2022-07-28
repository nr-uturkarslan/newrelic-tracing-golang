package services

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog"
	"github.com/segmentio/kafka-go"

	"github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/commons"
	dto "github.com/nr-turkarslan/newrelic-tracing-golang/apps/third/dtos"
)

func ThirdMethod() {

	// Start New Relic
	nrapp := commons.CreateNewRelicAgent()

	kafkaReader := createKafkaReader()

	for {

		// Read message
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			commons.Log(zerolog.ErrorLevel, "Kafka message could not be received.")
			continue
		}

		// Start transaction
		txn := nrapp.StartTransaction("test")

		// Get distributed tracing headers
		dtHeader := http.Header{}
		for _, header := range msg.Headers {
			headerKey := header.Key
			headerValue := string(header.Value)

			commons.LogWithContext(txn, zerolog.InfoLevel, "Header key: "+headerKey)
			commons.LogWithContext(txn, zerolog.InfoLevel, "Header value: "+headerValue)

			if header.Key == "traceparent" {
				dtHeader.Add("traceparent", string(headerValue))
			} else if header.Key == "tracestate" {
				dtHeader.Add("tracestate", string(headerValue))
			}
		}

		// Set distributed tracing headers
		txn.AcceptDistributedTraceHeaders(newrelic.TransportKafka, dtHeader)

		// Parse message
		body, err := parseMessage(msg.Value)
		if err != nil {
			commons.Log(zerolog.ErrorLevel, "Message could not be parsed.")
			continue
		}

		commons.LogWithContext(txn, zerolog.InfoLevel, "Value: "+body.Value)
		commons.LogWithContext(txn, zerolog.InfoLevel, "Tag: "+body.Tag)

		txn.End()
	}
}

func createKafkaReader() *kafka.Reader {
	commons.Log(zerolog.InfoLevel, "Starting Kafka...")

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka.kafka.svc.cluster.local:9092"},
		Topic:   "tracing",
		GroupID: "tracingconsumer",
	})

	commons.Log(zerolog.InfoLevel, "Kafka is started.")

	return kafkaReader
}

func parseMessage(
	message []byte,
) (
	*dto.RequestDto,
	error,
) {
	var requestDto dto.RequestDto

	err := json.Unmarshal(message, &requestDto)

	if err != nil {
		return nil, err
	}

	return &requestDto, nil
}
