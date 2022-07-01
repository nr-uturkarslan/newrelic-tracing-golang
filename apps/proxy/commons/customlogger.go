package commons

import (
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrzerolog"
	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type CustomLogger struct {
	Logger zerolog.Logger
}

var logger *CustomLogger

func CreateCustomLogger(
	nrapp *newrelic.Application,
) {

	log.Info().Msg("Initializing custom logger...")

	zerolog := zerolog.New(os.Stdout)

	nrLogger := zerolog.Hook(nrzerolog.NewRelicHook{
		App: nrapp,
	})

	logger = &CustomLogger{}

	logger.Logger = nrLogger
	logger.Logger.Info().Msg("Custom logger is initialized successfully.")
}

func Log(
	logLevel zerolog.Level,
	message string,
) {

	if logger == nil {
		panic("Custom logger is not initiated.")
	} else {

		switch logLevel {
		case zerolog.ErrorLevel:
			logger.Logger.Error().Msg(message)
		case zerolog.PanicLevel:
			logger.Logger.Panic().Msg(message)
		default:
			logger.Logger.Info().Msg(message)
		}
	}
}
