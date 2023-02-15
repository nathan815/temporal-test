package config

import (
	"fmt"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getRequiredEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(fmt.Sprintf("Required env var %s not set", key))
}

const MainTaskQueue = "main-task-queue"

var TemporalNamespace = getRequiredEnv("TEMPORAL_NAMESPACE")

var TemporalServerHostPort = getEnv("TEMPORAL_SERVER_HOST", "")
