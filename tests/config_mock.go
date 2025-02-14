package tests

import "jd-golang/src/config"

var config_mock = []struct {
	name           string
	envVars        map[string]string
	expectedConfig config.Config
}{
	{
		name: "default values",
		envVars: map[string]string{
			"SENTRY_TRACES_SAMPLE_RATE": "0.5",
			"SENTRY_DSN":                "https://examplePublicKey@o0.ingest.sentry.io/0",
			"SENTRY_RELEASE":            "1.0.0",
			"SENTRY_ENVIRONMENT":        "production",
		},
		expectedConfig: config.Config{
			Sentry: config.SentryConfig{
				TracesSampleRate: 0.5,
				Dsn:              "https://examplePublicKey@o0.ingest.sentry.io/0",
				Release:          "1.0.0",
				Environment:      "production",
			},
		},
	},
	{
		name:    "empty env vars",
		envVars: map[string]string{},
		expectedConfig: config.Config{
			Sentry: config.SentryConfig{
				TracesSampleRate: 0,
				Dsn:              "",
				Release:          "",
				Environment:      "",
			},
		},
	},
	{
		name: "invalid TracesSampleRate",
		envVars: map[string]string{
			"SENTRY_TRACES_SAMPLE_RATE": "abc",
		},
		expectedConfig: config.Config{
			Sentry: config.SentryConfig{
				TracesSampleRate: 0,
				Dsn:              "",
				Release:          "",
				Environment:      "",
			},
		},
	},
}
