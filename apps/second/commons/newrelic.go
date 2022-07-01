package commons

import (
	"os"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/rs/zerolog/log"
)

func CreateNewRelicAgent() *newrelic.Application {

	log.Info().Msg("Starting New Relic agent...")

	os.Setenv("NEW_RELIC_ENABLED", "true")
	os.Setenv("NEW_RELIC_LICENSE_KEY", os.Getenv("NEWRELIC_LICENSE_KEY"))
	os.Setenv("NEW_RELIC_APP_NAME", "second-go")
	os.Setenv("NEW_RELIC_DISTRIBUTED_TRACING_ENABLED", "true")
	os.Setenv("NEW_RELIC_LOG", "stdout")
	// os.Setenv("NEW_RELIC_LOG_LEVEL", "debug")

	nrapp, err := newrelic.NewApplication(
		newrelic.ConfigFromEnvironment(),
	)

	if err != nil {
		message := "New Relic agent could not be started."
		log.Panic().Msg(message)
		panic(message)
	}

	log.Info().Msg("New Relic agent is started successfully.")

	CreateCustomLogger(nrapp)
	return nrapp
}
