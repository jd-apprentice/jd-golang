package main

import (
	"log"

	"jd-golang/src/config"

	"github.com/getsentry/sentry-go"
)

func main() {

	config := config.GetConfig()

	if config.Sentry.Environment == "develop" {
		config.Sentry.Release = "unreleased"
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Sentry.Dsn,
		TracesSampleRate: config.Sentry.TracesSampleRate,
		Release:          config.Sentry.Release,
		Environment:      config.Sentry.Environment,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	log.Println("Hello, world!")
	sentry.CaptureMessage("Integration is working!!")
}
