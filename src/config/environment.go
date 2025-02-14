package config

import (
	"os"
	"strconv"
	"strings"
)

type SentryConfig struct {
	Dsn              string
	TracesSampleRate float64
	Release          string
	Environment      string
}

type Config struct {
	Sentry SentryConfig
}

// GetConfig retrieves the configuration for the application from environment variables.
// It returns a Config struct containing Sentry configuration details such as DSN, release,
// environment, and traces sample rate. If the environment variable for the traces sample
// rate is not set or invalid, it defaults to zero.
func GetConfig() Config {

	parsedFloat, _ := strconv.ParseFloat(strings.TrimSpace(os.Getenv("SENTRY_TRACES_SAMPLE_RATE")), 64)

	return Config{
		Sentry: SentryConfig{
			TracesSampleRate: parsedFloat,
			Dsn:              os.Getenv("SENTRY_DSN"),
			Release:          os.Getenv("SENTRY_RELEASE"),
			Environment:      os.Getenv("SENTRY_ENVIRONMENT"),
		},
	}
}
